/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/csv"
	"fmt"
	"os"
	"text/tabwriter"
	"time"

	"github.com/spf13/cobra"
)

type Task struct {
	ID          string
	Description string
	CreatedAt   string
	IsComplete  string
}

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List tasks that are not completed",
	Long: `This command is used to print tasks on screen, you have two options,
print by default all the incomplete tasks, or add a flag -a to print them all. 

For example:
CLITodoApp.git list -> will print the default ones or incomplete.
CLITodoApp.git list -a or --all -> will print all the tasks.`,
	Run: func(cmd *cobra.Command, args []string) {

		allFlag, err := cmd.Flags().GetBool("all")

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		file, err := os.Open("task.csv")

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		// create a csv reader
		reader := csv.NewReader(file)

		// read all the csv
		records, err := reader.ReadAll()
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading CSV file:", err)
			os.Exit(-1)
		}

		var tasks []Task
		for i, record := range records {
			if i == 0 {
				continue // skip headers
			}
			tasks = append(tasks, Task{
				ID:          record[0],
				Description: record[1],
				CreatedAt:   record[2],
				IsComplete:  record[3],
			})
		}

		// calculate width necessary for each column
		maxIDWidth := len("ID")
		maxDescWidth := len("Task")
		maxCreatedWidth := len("Created")
		maxDoneWidth := len("Done")

		for _, task := range tasks {
			if len(task.ID) > maxIDWidth {
				maxIDWidth = len(task.ID)
			}
			if len(task.Description) > maxDescWidth {
				maxDescWidth = len(task.Description)
			}

			createdAt, err := time.Parse(time.RFC3339, task.CreatedAt)
			if err != nil {
				fmt.Println("Error parsing date:", err)
				return
			}
			created := time.Since(createdAt).Round(time.Minute)
			createdStr := fmt.Sprintf("%v ago", created)

			if len(createdStr) > maxCreatedWidth {
				maxCreatedWidth = len(createdStr)
			}
			if len(task.IsComplete) > maxDoneWidth {
				maxDoneWidth = len(task.IsComplete)
			}
		}

		// create the tabwriter
		w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)

		// print header based on each width
		fmt.Fprintf(w, "%-*s  %-*s  %-*s  %-*s\n",
			maxIDWidth, "ID",
			maxDescWidth, "Task",
			maxCreatedWidth, "Created",
			maxDoneWidth, "Done",
		)

		// print task based on each width
		for _, task := range tasks {
			if !allFlag && task.IsComplete == "true" {
				continue
			}
			createdAt, err := time.Parse(time.RFC3339, task.CreatedAt)
			if err != nil {
				fmt.Println("Error parsing date:", err)
				return
			}
			created := time.Since(createdAt).Round(time.Minute)
			createdStr := fmt.Sprintf("%v ago", created)

			fmt.Fprintf(w, "%-*s  %-*s  %-*s  %-*s\n",
				maxIDWidth, task.ID,
				maxDescWidth, task.Description,
				maxCreatedWidth, createdStr,
				maxDoneWidth, task.IsComplete,
			)
		}

		w.Flush()

	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	listCmd.PersistentFlags().BoolP("all", "a", false, "flag to show all the tasks, default is false and just show the incomplete ones")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
