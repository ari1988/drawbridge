package config_test

import (
	"testing"
	"path"
	"github.com/stretchr/testify/require"
	"drawbridge/pkg/config"
)

func TestConfiguration_GetQuestion(t *testing.T) {

	//setup
	testConfig, _ := config.Create()

	//test
	err := testConfig.ReadConfig(path.Join("testdata", "valid_questions.yaml"))
	question, err := testConfig.GetQuestion("environment")

	//assert
	require.NoError(t, err, "should not have an error when requesting question by key")
	require.EqualValues(t, question.Description, "what is the stack environment","should retrieve correct question description")
}


func TestConfiguration_GetQuestion_Invalid(t *testing.T) {

	//setup
	testConfig, _ := config.Create()

	//test
	err := testConfig.ReadConfig(path.Join("testdata", "valid_questions.yaml"))
	_, err = testConfig.GetQuestion("invalidkey")

	//assert
	require.Error(t, err, "should have an error when requesting invalid question by key")
}

func TestQuestion_Validate(t *testing.T) {

	//setup
	testConfig, _ := config.Create()

	//test
	err := testConfig.ReadConfig(path.Join("testdata", "valid_questions.yaml"))
	question, err := testConfig.GetQuestion("environment")
	err = question.Validate("environment", "testing")

	//assert
	require.NoError(t, err, "should not throw an error")
}

func TestQuestion_Validate_LengthRuleBroken(t *testing.T) {

	//setup
	testConfig, _ := config.Create()

	//test
	err := testConfig.ReadConfig(path.Join("testdata", "valid_questions.yaml"))
	question, err := testConfig.GetQuestion("environment")
	err = question.Validate("environment", "testing.this.value.is.too.long")

	//assert
	require.Error(t, err, "should fail when length rule is broken")
}
