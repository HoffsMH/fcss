package main

import (
	"github.com/spf13/cobra"
	"fmt"
	"os"
)

var days int
var weeks int
var months int
var dir string

var rootCmd = &cobra.Command{
	Use: "fcss",
	Run: func(cmd *cobra.Command,  args []string) {
		// fcss.FindClassesInFile(args[0], args[1])
		fmt.Println("hi")
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
