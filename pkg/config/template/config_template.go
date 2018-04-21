package template

import (
	"path"
	"fmt"
	"drawbridge/pkg/utils"
)

// for configs `filepath`, must be relative to config_dir
//for configs `pem_filepath` must be relative to pem_dir
type ConfigTemplate struct {
	FileTemplate `mapstructure:",squash"`
	PemFilePath string `mapstructure:"pem_filepath"`
}

func (t *ConfigTemplate) WriteTemplate(answerData map[string]interface{}) error {
	//TODO: validate that we have all the required templates variables populated.

	// modify/tweak the config template because its a known type.
	//expand PemFilePath
	t.PemFilePath = path.Join(answerData["pem_dir"].(string), t.PemFilePath)
	templatedPemFilePath, err := utils.PopulateTemplate(t.PemFilePath, answerData)
	if err != nil {
		return err
	}
	templatedPemFilePath, err = utils.ExpandPath(templatedPemFilePath)
	if err != nil {
		return err
	}
	answerData["pem_filepath"] = templatedPemFilePath

	t.FilePath = path.Join(answerData["config_dir"].(string), t.FilePath)
	t.Content = configTemplatePrefix(answerData) + t.Content

	return t.FileTemplate.WriteTemplate(answerData)
}

func configTemplatePrefix(answerData map[string]interface{}) string {
	prefix := utils.StripIndent(
		`
		# This file was automatically generated by Drawbridge
		# Do not modify.
		#
		# Answers:`)
	for key, value := range answerData {
		if key == "config_dir" || key == "pem_dir" || key == "active_extra_templates" || key == "ui_group_priority" {
			continue
		}

		prefix += fmt.Sprintf("\n# %v = %v", key, value)
	}
	prefix += "\n"
	return prefix
}