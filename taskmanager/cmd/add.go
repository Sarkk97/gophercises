package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/Sarkk97/gophercises/taskmanager/database"
	"github.com/boltdb/bolt"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "adds a task to the to-do list",
	Run: func(cmd *cobra.Command, args []string) {
		taskName := strings.Join(args, " ")
		newTask, _ := json.Marshal(map[string]interface{}{
			"task":      taskName,
			"completed": false,
		})
		db := database.GetDB()
		//Add task to db
		err := db.Update(func(tx *bolt.Tx) error {
			//Retrieve tasks bucket
			b := tx.Bucket([]byte("tasks"))
			if b == nil {
				return errors.New("Bucket does not exist")
			}
			//generate ID for the new task
			id, _ := b.NextSequence()

			//add task to bucket with id as key and task as value
			buf, _ := json.Marshal(id)
			err := b.Put(buf, newTask)
			if err != nil {
				return err
			}
			fmt.Printf("Successfully added the task: %v\n", taskName)
			return nil

		})
		if err != nil {
			log.Fatal(err)
		}

	},
	Args: cobra.MinimumNArgs(1),
}
