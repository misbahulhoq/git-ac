/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

// commitCmd represents the commit command
var commitCmd = &cobra.Command{
	Use:   "commit",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if !IsGitRepo() {
			fmt.Println("❌ Not a git repository. Did you run \"git init\" ?")
			return
		}
		output, err := GetStagedChanges()

		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(output)
	},
}

func init() {
	rootCmd.AddCommand(commitCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// commitCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	commitCmd.Flags().BoolP("staged", "s", false, "Commit the staged changes")
	commitCmd.Flags().BoolP("all", "a", true, "Commit both staged and unstaged changes.")
}

func GetStagedChanges() (string, error) {
	cmd := exec.Command("git", "diff", "--staged")
	output, err := cmd.Output()

	if err != nil {
		return "", err
	}

	diff := strings.TrimSpace(string(output))

	if diff == "" {
		return "", fmt.Errorf("No staged changes found. Did you run \"git add\" ?")
	}

	return diff, nil
}

func Commit() {
	cmd := exec.Command("git", "diff")
	output, err := cmd.Output()

	if err != nil {
		fmt.Println(output)
	}
}
