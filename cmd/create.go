/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) == 0 {
			fmt.Fprintln(os.Stderr, "You need to pass a task")
			os.Exit(-1)
		}
		newTask := args[0]

		// Open CSV file
		file, err := os.OpenFile("task.csv", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(-1)
		}
		defer file.Close()

		// create writer
		writer := csv.NewWriter(file)
		defer writer.Flush()

		// write records
		err = writer.Write([]string{newTask})
		if err != nil {
			panic(err)
		}

		fmt.Println("task crated!")
	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
