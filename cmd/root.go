package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "gh-s",
	Short: "short description",
	Long: "Long description",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("hello from gh-s")
	},
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	rootCmd.SetHelpTemplate(getRootHelp())
}

func getRootHelp() string {
	return `
	test help for gh-s
`
}
