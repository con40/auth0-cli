/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
    "github.com/withfig/autocomplete-tools/integrations/cobra"
)

// generateFigSpecCmd represents the generateFigSpec command
var generateFigSpecCmd = &cobra.Command{
	Use:   "generateFigSpec",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("generateFigSpec called")
	},
}

func init() {
    rootCmd.AddCommand(cobracompletefig.CreateCompletionSpecCommand())
}
