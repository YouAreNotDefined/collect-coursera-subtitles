/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

// collectCmd represents the collect command
var collectCmd = &cobra.Command{
	Use:   "collect",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: collect,
}

func collect(cmd *cobra.Command, args []string) {
	if err := filepath.Walk(currentDir(), traverse); err != nil {
		handleErr(err)
	}
}

func traverse(path string, info os.FileInfo, err error) error {
	if info.IsDir() {
		fmt.Printf(" dir : %s\n", path)
		return nil
	}

	// Read file order by asc and write that file content to another file.
	fmt.Printf("file : %s\n", path)
	return nil
}

func currentDir() string {
	dir, _ := os.Getwd()
	return dir
}

func handleErr(err error) {
	if err != nil {
		log.Printf("Error: %v\n", err)
	}
}

func init() {
	rootCmd.AddCommand(collectCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// collectCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// collectCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
