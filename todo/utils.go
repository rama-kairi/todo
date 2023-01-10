package todo

import (
	"encoding/json"
	"io"
	"os"
)

// SaveToJson - This function saves the todo list to a json file
func (t todoList) SaveToJson() {
	// Delete the file if it exists
	if _, err := os.Stat("db.json"); !os.IsNotExist(err) {
		os.Remove("db.json")
	}

	// Create a Json file with name `db.json`
	os.Create("db.json")

	// Open the file
	file, err := os.OpenFile("db.json", os.O_RDWR, 0o644)
	if err != nil {
		panic(err)
	}

	// Marshal the data from t.todoStore
	data, err := json.Marshal(t.todoStore)
	if err != nil {
		panic(err)
	}

	// Write the data to the file
	_, err = file.Write([]byte(data))
	if err != nil {
		panic(err)
	}

	// Close the file
	err = file.Close()
	if err != nil {
		panic(err)
	}
}

// LoadFromJson - This function loads the todo list from a json file
func (t todoList) LoadFromJson() {
	// Check if the file exists
	if _, err := os.Stat("db.json"); os.IsNotExist(err) {
		os.Create("db.json")
	}

	// Open the file
	file, err := os.OpenFile("db.json", os.O_RDONLY, 0o644)
	if err != nil {
		panic(err)
	}

	// convert the file to a byte array
	fileByte, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	if len(fileByte) > 0 {
		// Unmarshal the data from the file to t.todoStore
		err = json.Unmarshal([]byte(fileByte), &t.todoStore)
		if err != nil {
			panic(err)
		}
	}

	// Close the file
	err = file.Close()
	if err != nil {
		panic(err)
	}
}
