package cmd

import (
	"github.com/manifoldco/promptui"
)

var Template = &promptui.SelectTemplates{
	Active:   "\U0001F449 {{ .Name | cyan | bold }}",
	Inactive: "   {{ .Name | cyan }}",
	Selected: `{{ "✔" | green | bold }} {{ .Name | cyan | bold }}`,
	Details: `
{{ "Name:" | faint }} 	 {{ .Name }}
{{ "Description:" | faint }} 	 {{ .Description }}
{{ "Url address:" | faint }} 	 {{ .URL }}
{{ "⭐" | faint }}	{{ .Stars }}`,
}
