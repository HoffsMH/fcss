package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/hoffsmh/fcss"
	"github.com/spf13/cobra"
)

var days int
var weeks int
var months int
var dir string

var rootCmd = &cobra.Command{
	Use: "fcss",
	Run: func(cmd *cobra.Command, args []string) {
		scanner := bufio.NewScanner(os.Stdin)

		selectors := []string{}
		for scanner.Scan() {
			selectors = append(selectors, fcss.FindClassesInFile(scanner.Text())...)
		}

		for _ ,selector := range selectors {
			fmt.Println(selector)
		}
	},
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
}
