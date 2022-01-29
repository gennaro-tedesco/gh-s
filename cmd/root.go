package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/cli/go-gh"
	"github.com/spf13/cobra"
)

var cfgFile string

// VERSION number: change manually
const VERSION = "0.0.1"

var rootCmd = &cobra.Command{
	Use:   "gh-s",
	Short: "gh-s: search repositories interactively",
	Long:  "gh-s: interactive prompt to search and browse github repositories",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		version, _ := cmd.Flags().GetBool("version")
		if version {
			fmt.Println(VERSION)
			os.Exit(1)
		}
		languageList, _ := cmd.Flags().GetStringSlice("lang")
		desc, _ := cmd.Flags().GetString("desc")
		user, _ := cmd.Flags().GetString("user")
		topicList, _ := cmd.Flags().GetStringSlice("topic")
		colour, _ := cmd.Flags().GetString("colour")

		searchString := getSearchString(args)
		parsedQuery := parseInput(searchString, languageList, desc, user, topicList)
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
	var topics []string
	var languages []string
	rootCmd.Flags().StringSliceVarP(&languages, "lang", "l", []string{}, "specify repository language")
	rootCmd.Flags().StringP("desc", "d", "", "search in repository description")
	rootCmd.Flags().StringP("user", "u", "", "search repository by user")
	rootCmd.Flags().StringSliceVarP(&topics, "topic", "t", []string{}, "search repository by topic")
	rootCmd.Flags().StringP("colour", "c", "cyan", "colour of selection prompt")
	rootCmd.Flags().BoolP("version", "V", false, "print current version")
	rootCmd.SetHelpTemplate(getRootHelp())
}

func getRootHelp() string {
	return `
gh-s: search repositories interactively. The search returns results
matching the indicated query ordered by number of repository stars.

Synopsis:
	gh s [search] [flags]

Usage:
	gh s

	if no arguments or flags are given, show an interactive prompt
	to search, browse and filter repositories. Selecting an entry
	from the list opens the repository in the web browser.

	Flags can be passed so that the search is narrowed down (see available
	flags below). For example:

	gh s -l lua -d quickfix

	If you provide an argument before the flags the prompt is skipped and such
	argument is used in the name field to search for repositories:

	gh s ripgrep -l rust

Prompt commands:

	arrow keys  : move up and down the list
	/           : toggle fuzzy search
	enter (<CR>): open selected repository in the web browser

Flags:
  -l, --lang    search repositories with specific language
  -d, --desc    match repository description
  -u, --user    narrow the search down to a specific user's repositories
  -c, --colour  change prompt colour
  -V, --version print current version
  -h, --help    show this help page

Examples:

	# search for name=ripgrep and language=rust
	gh s ripgrep -l rust

	# what is the most starred neovim plugin?
	gh s neovim -d plugin

	# restrict to one user only
	gh s lsp -u neovim

Help commands:
  help        show this help page
  version     print current version
`
}
