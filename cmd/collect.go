/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"

	"github.com/spf13/cobra"
)

const OutFileName = "subtitle.txt"

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
	if outFileNotExist() {
		file, err := os.Create(OutFileName)
		handleErr(err)
		defer file.Close()
	}
	if err := filepath.Walk(currentDir(), traverse); err != nil {
		handleErr(err)
	}
}

func traverse(path string, info os.FileInfo, err error) error {
	if info.IsDir() {
		fmt.Printf(" dir : %s\n", path)
		return nil
	}

	r := regexp.MustCompile(`.*ja$`)

	if r.Match([]byte(path)) {
		fileContent, err := os.ReadFile(path)
		handleErr(err)
		f, err := os.OpenFile(OutFileName, os.O_APPEND|os.O_WRONLY, 0644)
		handleErr(err)
		// contentBlock := fmt.Sprintf("%s%s", string(fileContent), "\n")
		_, err = fmt.Fprintln(f, string(fileContent))
		handleErr(err)
	}
	return nil
}

func outFileNotExist() bool {
	outFilePath := fmt.Sprintf("%s%s", currentDir(), OutFileName)
	_, err := os.Stat(outFilePath)
	return os.IsNotExist(err)
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
