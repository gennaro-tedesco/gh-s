package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseInput(t *testing.T) {
	inputValues := map[string]interface{}{
		"search":       "awesome",
		"languageList": []string{"go", "lua"},
		"desc":         "framework",
		"user":         "@me",
		"topicList":    []string{"cli", "gh-extension"},
	}

	trueString := "awesome in:name language:go language:lua framework in:description user:@me topic:cli topic:gh-extension"
	parsedString := parseInput(inputValues["search"].(string), inputValues["languageList"].([]string), inputValues["desc"].(string), inputValues["user"].(string), inputValues["topicList"].([]string))

	assert.Equal(t, trueString, parsedString["q"][0])
}
