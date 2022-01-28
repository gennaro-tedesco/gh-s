package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/cli/go-gh"
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
		user, _ := cmd.Flags().GetString("user")
		colour, _ := cmd.Flags().GetString("colour")
		parsedQuery := getInputPrompt(args, lang, desc, user)
		repos := getRepos(parsedQuery)
		PromptList := getSelectionPrompt(repos, colour)

		idx, _, err := PromptList.Run()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		exArgs := []string{"repo", "view", "-w", repos[idx].URL}
		_, _, err = gh.Exec(exArgs...)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	rootCmd.Flags().StringP("lang", "l", "", "specify repository language")
	rootCmd.Flags().StringP("desc", "d", "", "search in repository description")
	rootCmd.Flags().StringP("user", "u", "", "search repository by user")
	rootCmd.Flags().StringP("colour", "c", "cyan", "colour of selection prompt")
	rootCmd.SetHelpTemplate(getRootHelp())
}

func getRootHelp() string {
	return `
	test help for gh-s
`
}
