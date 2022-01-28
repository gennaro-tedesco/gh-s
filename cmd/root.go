package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "gh-s",
	Short: "short description",
	Long:  "Long description",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		lang, _ := cmd.Flags().GetString("lang")
		desc, _ := cmd.Flags().GetString("desc")
		var parsedQuery = parseInput(args[0], lang, desc)
		repos := getRepos(parsedQuery)

		var PromptList = promptui.Select{
			Label:     "repository list",
			Items:     repos,
			Templates: Template,
			Size:      20,
			Searcher: func(input string, idx int) bool {
				repo := repos[idx]
				title := strings.ToLower(repo.Name)

				return strings.Contains(title, input)
			},
		}

		idx, _, err := PromptList.Run()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(repos[idx])
	},
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	rootCmd.Flags().StringP("lang", "l", "", "specify repository language")
	rootCmd.Flags().StringP("desc", "d", "", "search in repository description")
	rootCmd.SetHelpTemplate(getRootHelp())
}

func getRootHelp() string {
	return `
	test help for gh-s
`
}
