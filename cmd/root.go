package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "gh-s",
	Short: "short description",
	Long:  "Long description",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		lang, _ := cmd.Flags().GetString("lang")
		desc, _ := cmd.Flags().GetString("desc")
		colour, _ := cmd.Flags().GetString("colour")
		parsedQuery := getInputPrompt(args, lang, desc)
		repos := getRepos(parsedQuery)
		PromptList := getSelectionPrompt(repos, colour)

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
	rootCmd.Flags().StringP("colour", "c", "cyan", "colour of selection prompt")
	rootCmd.SetHelpTemplate(getRootHelp())
}

func getRootHelp() string {
	return `
	test help for gh-s
`
}
