package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var cfgFile string

// VERSION number: change manually
const VERSION = "0.0.2"

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

		searchString := func() string {
			if empty, _ := (cmd.Flags().GetBool("empty")); empty {
				return ""
			}
			return getSearchString(args)
		}()
		parsedQuery := parseInput(searchString, languageList, desc, user, topicList)
		repos := getRepos(parsedQuery)
		PromptList := getSelectionPrompt(repos, colour)

		idx, _, err := PromptList.Run()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(repos[idx].URL)
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
	rootCmd.Flags().BoolP("empty", "E", false, "allow for empty name search")
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
  -E, --empty   allow to pass an empty string as name, that is search
  				github repositories based on topic and language only
  -l, --lang    search repositories with specific language
  				multiple languages can be specified:
				-l go -l rust -l lua
  -d, --desc    match repository description
  -u, --user    narrow the search down to a specific user's repositories
  -t, --topic   search for topics in repositories
  				multiple topics can be specified:
				-t go -t gh-extension
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

	# all neovim plugins in lua of nvim-*
	gh s nvim -t plugin -l lua

	# the most famous go or rust frameworks
	gh s -E -l go -l rust

Help commands:
  help        show this help page
  version     print current version
`
}
