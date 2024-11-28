package cmd

import (
	"github.com/spf13/cobra"
)

// categoryCmd represents the category command
var categoryCmd = &cobra.Command{
	Use:   "category",
	Short: "Handle category commands",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	rootCmd.AddCommand(categoryCmd)
}
