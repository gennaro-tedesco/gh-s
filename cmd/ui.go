package cmd

import (
	"fmt"

	"github.com/manifoldco/promptui"
)

func getTemplate() *promptui.SelectTemplates {
	funcMap := promptui.FuncMap
	funcMap["parseStars"] = func(starCount float64) string {
		if starCount >= 1000 {
			return fmt.Sprintf("%.1f k", starCount/1000)
		}
		return fmt.Sprint(starCount)
	}

	funcMap["truncate"] = func(input string) string {
		length := 80
		if len(input) <= length {
			return input
		}
		return input[:length-3] + "..."
	}

	return &promptui.SelectTemplates{
		Active:   "\U0001F449 {{ .Name | cyan | bold }}",
		Inactive: "   {{ .Name | cyan }}",
		Selected: `{{ "✔" | green | bold }} {{ .Name | cyan | bold }}`,
		Details: `
	{{ "Name:" | faint }} 	 {{ .Name }}
	{{ "Description:" | faint }} 	 {{ .Description | truncate }}
	{{ "Url address:" | faint }} 	 {{ .URL }}
	{{ "⭐" | faint }}	{{ .Stars | parseStars }}`,
	}

}
