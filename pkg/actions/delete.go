package actions

import (
	"drawbridge/pkg/config"
	"drawbridge/pkg/utils"
	"fmt"
	"github.com/fatih/color"
	"path"
	"strconv"
	"strings"
)

type DeleteAction struct {
	Config config.Interface
}

func (e *DeleteAction) All(answerDataList []map[string]interface{}, force bool) error {

	for _, v := range answerDataList {
		err := e.One(v, force)
		if err != nil {
			color.Red("ERROR IGNORED: %v", err)
		}
	}
	return nil
}
func (e *DeleteAction) One(answerData map[string]interface{}, force bool) error {

	//delete the config file by answerData
	renderedConfigFilePath := answerData["filepath"].(string)

	if !force {

		questionStr := []string{"Are you sure you would like to delete this config and associated templates? (PEM files will not be deleted)\n"}

		for k, v := range answerData {
			if utils.StringInSlice(e.Config.InternalQuestionKeys(), k) {
				continue
			}
			questionStr = append(questionStr, fmt.Sprintf("%v: %v", k, v))
		}
		questionStr = append(questionStr, "\nPlease confirm [true/false]:")

		resp := utils.StdinQuery(strings.Join(questionStr, "\n"))
		val, err := strconv.ParseBool(resp)
		if err != nil {
			return err
		}
		if !val {
			color.Red("Cancelled delete operation.")
			return nil
		}
	}

	fmt.Printf("Deleting config file: %v\n", renderedConfigFilePath)
	if utils.FileExists(renderedConfigFilePath) {
		utils.FileDelete(renderedConfigFilePath)
	} else {
		color.Yellow(" - Skipping. Could not find config file at: %v", renderedConfigFilePath)
	}

	//delete any custom templates.
	if val, ok := answerData["active_custom_templates"]; ok {
		renderedCustomTemplateNames := val.([]interface{})

		// load up all customTemplates
		customTemplates, err := e.Config.GetCustomTemplates()
		if err != nil {
			return err
		}

		fmt.Println("Deleting custom template files")
		for _, renderedCustomTemplateName := range renderedCustomTemplateNames {
			if renderedCustomTemplate, ok := customTemplates[renderedCustomTemplateName.(string)]; ok {
				err = renderedCustomTemplate.DeleteTemplate(answerData)
				if err != nil {
					color.Yellow(" - Skipping. An error occurred while deleting %v: %v", renderedCustomTemplateName, err)
				}
			}
		}

	}

	//delete the .answers.yaml
	fmt.Println("Deleting answers file")
	answersFilePath := path.Join(answerData["config_dir"].(string), fmt.Sprintf(".%v.answers.yaml", path.Base(renderedConfigFilePath)))
	if utils.FileExists(answersFilePath) {
		utils.FileDelete(answersFilePath)
	} else {
		color.Yellow(" - Skipping. Could not find answers file at: %v", answersFilePath)
	}

	return nil
}
