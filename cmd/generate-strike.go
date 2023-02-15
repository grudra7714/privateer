package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var genStrikeName = "generate-strike"

// versionCmd represents the version command
var genStrikeCmd = &cobra.Command{
	Use:   genStrikeName,
	Short: "Generate a new strike",
	Long:  `TODO - Long description`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("%s called", genStrikeName)
	},
}

func init() {
	rootCmd.AddCommand(genStrikeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// versionCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// versionCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
