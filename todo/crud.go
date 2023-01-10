package todo

import "time"

// AddTodo adds a new todo item to the list
func (t *todoList) AddTodo(newTodo string) error {
	newTodoIns := todo{
		Task:      newTodo,
		Completed: false,
		Added:     time.Now(),
	}
	t.todoStore = append(t.todoStore, newTodoIns)

	// Save the todo list to a json file
	t.SaveToJson()
	return nil
}

// RemoveTodo removes a todo item from the list
func (t *todoList) RemoveTodo(todo string) error {
	for i, v := range t.todoStore {
		if v.Task == todo {
			t.todoStore = append(t.todoStore[:i], t.todoStore[i+1:]...)
		}
	}
	// Save the todo list to a json file
	t.SaveToJson()
	return nil
}

// ListTodo lists all todo items
func (t *todoList) ListTodo() []string {
	var todoList []string
	for _, v := range t.todoStore {
		todoList = append(todoList, v.Task)
	}
	return todoList
}
