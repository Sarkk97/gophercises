package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "task",
	Short: "task is a CLI for managing your to-do",
}

//Initialize initializes sub commands
func Initialize() {
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(doCmd)
	rootCmd.AddCommand(completeCmd)
	rootCmd.AddCommand(rmCmd)

}

//Execute runs the cmd
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
