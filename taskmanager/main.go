package main

import (
	"github.com/Sarkk97/gophercises/taskmanager/cmd"
	_ "github.com/Sarkk97/gophercises/taskmanager/database"
)

func main() {
	cmd.Initialize()
	cmd.Execute()
}
