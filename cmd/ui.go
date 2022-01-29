package cmd

import (
	"errors"
	"fmt"
	"log"
	"net/url"
	"strings"

	"github.com/manifoldco/promptui"
)

func getSearchString(args []string) string {
	if len(args) == 0 {
		prompt := promptui.Prompt{
			Label: "Repository name",
			Validate: func(input string) error {
				if len(input) == 0 {
					return errors.New("no input provided")
				}
				return nil
			},
		}

		result, err := prompt.Run()
		if err != nil {
			log.Fatal(err)
		}
		return result
	}
	return args[0]
}

func parseInput(search string, lang string, desc string, user string, topicList []string) url.Values {
	queryString := fmt.Sprintf("%s in:name", search)
	if lang != "" {
		queryString = queryString + fmt.Sprintf(" language:%s", lang)
	}
	if desc != "" {
		queryString = queryString + fmt.Sprintf(" %s in:description", desc)
	}
	if user != "" {
		queryString = queryString + fmt.Sprintf(" user:%s", user)
	}
	for _, topic := range topicList {
		queryString = queryString + fmt.Sprintf(" topic:%s", topic)
	}
	query := url.Values{}
	query.Add("q", queryString)
	query.Add("sort", "stars")
	query.Add("per_page", "100")
	return query
}

func getTemplate(colour string) *promptui.SelectTemplates {
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
		Active:   fmt.Sprintf("\U0001F449 {{ .Name | %s | bold }}", colour),
		Inactive: fmt.Sprintf("   {{ .Name | %s }}", colour),
		Selected: fmt.Sprintf(`{{ "✔" | green | bold }} {{ .Name | %s | bold }}`, colour),
		Details: `
	{{ "Name:" | faint }} 	 {{ .Name }}
	{{ "Description:" | faint }} 	 {{ .Description | truncate }}
	{{ "Url address:" | faint }} 	 {{ .URL }}
	{{ "⭐" | faint }}	{{ .Stars | parseStars }}`,
	}

}

func getSelectionPrompt(repos []repoInfo, colour string) *promptui.Select {
	return &promptui.Select{
		Label:     "repository list",
		Items:     repos,
		Templates: getTemplate(colour),
		Size:      20,
		Searcher: func(input string, idx int) bool {
			repo := repos[idx]
			title := strings.ToLower(repo.Name)

			return strings.Contains(title, input)
		},
	}
}
