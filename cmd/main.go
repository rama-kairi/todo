package cmd

import (
	"fmt"
	"strconv"

	"github.com/rama-kairi/todo/todo"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "todo [command]",
	Short: "A todo list application",
	Long:  `A Fast and Flexible todo list application built with Go`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
		fmt.Println("Hello World")
	},
}

// AddCmd represents the add command
var AddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new todo",
	Long:  `Add a new todo to the list`,
	Run: func(cmd *cobra.Command, args []string) {
		td := todo.NewTodoService()
		td.AddTodo(args)
	},
}

// RemoveCmd represents the remove command
var RemoveCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove a todo",
	Long:  `Remove a todo from the list`,
	Run: func(cmd *cobra.Command, args []string) {
		td := todo.NewTodoService()
		if len(args[1:]) > 1 {
			fmt.Println("Invalid command")
			return
		}
		id, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Println("Invalid command")
			return
		}
		td.RemoveTodo(id)
	},
}

// ListCmd represents the list command
var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all todo's",
	Long:  `List all todo's from the list`,
	Run: func(cmd *cobra.Command, args []string) {
		td := todo.NewTodoService()
		// if len(args[1:]) > 0 {
		// 	fmt.Println("Invalid command")
		// 	return
		// }
		td.ListTodo()
	},
}

// Add the add command to the root command
func init() {
	RootCmd.AddCommand(AddCmd)
	RootCmd.AddCommand(RemoveCmd)
	RootCmd.AddCommand(ListCmd)
}