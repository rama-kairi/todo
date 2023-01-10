package main

import (
	"fmt"

	"github.com/rama-kairi/todo/todo"
)

func main() {
	newTodoSer := todo.NewTodoService()
	newTodoSer.LoadFromJson()

	fmt.Println(newTodoSer.AddTodo("Learn Go"))
	fmt.Println(newTodoSer.AddTodo("Learn Rust"))
	fmt.Println(newTodoSer.AddTodo("Learn Python"))

	fmt.Println(newTodoSer.ListTodo())
}
