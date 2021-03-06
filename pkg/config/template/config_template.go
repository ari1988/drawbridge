package template

import (
	"drawbridge/pkg/utils"
	"fmt"
	"github.com/fatih/color"
	"path"
)

// for configs `filepath`, must be relative to config_dir
//for configs `pem_filepath` must be relative to pem_dir
type ConfigTemplate struct {
	FileTemplate `mapstructure:",squash"`
	PemFilePath  string `mapstructure:"pem_filepath"`
}

func (t *ConfigTemplate) DeleteTemplate(answerData map[string]interface{}) error {
	t.FilePath = path.Join(answerData["config_dir"].(string), t.FilePath)
	return t.FileTemplate.DeleteTemplate(answerData)
}

func (t *ConfigTemplate) WriteTemplate(answerData map[string]interface{}, ignoreKeys []string, dryRun bool) (map[string]interface{}, error) {
	//intialize template data.
	if t.data == nil {
		t.data = map[string]interface{}{}
	}

	answerData, err := utils.MapDeepCopy(answerData)
	if err != nil {
		return nil, err
	}

	// modify/tweak the config template because its a known type.
	//expand PemFilePath
	t.PemFilePath = path.Join(answerData["pem_dir"].(string), t.PemFilePath)
	templatedPemFilePath, err := utils.PopulateTemplate(t.PemFilePath, answerData)
	if err != nil {
		return nil, err
	}
	templatedPemFilePath, err = utils.ExpandPath(templatedPemFilePath)
	if err != nil {
		return nil, err
	}

	t.data["pem_filepath"] = templatedPemFilePath
	answerData["template"] = t.data

	if !utils.FileExists(templatedPemFilePath) {
		color.Yellow("WARNING: PEM file missing. Place it at the following location before attempting to connect. %v", templatedPemFilePath)
	}

	t.FilePath = path.Join(answerData["config_dir"].(string), t.FilePath)
	t.Content = configTemplatePrefix(answerData, ignoreKeys) + t.Content

	_, err = t.FileTemplate.WriteTemplate(answerData, dryRun)
	if err != nil {
		return nil, err
	}

	return t.data, nil

}

func configTemplatePrefix(answerData map[string]interface{}, ignoreKeys []string) string {
	prefix := utils.StripIndent(
		`
		# This file was automatically generated by Drawbridge
		# Do not modify.
		#
		# Answers:`)
	for key, value := range answerData {
		if utils.SliceIncludes(ignoreKeys, key) {
			continue
		}

		prefix += fmt.Sprintf("\n# %v = %v", key, value)
	}
	prefix += "\n"
	return prefix
}
