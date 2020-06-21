package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"github.com/Sarkk97/gophercises/taskmanager/database"
	"github.com/boltdb/bolt"
	"github.com/spf13/cobra"
)

var completeCmd = &cobra.Command{
	Use:   "completed",
	Short: "lists all completed tasks in the to-do list",
	Run: func(cmd *cobra.Command, args []string) {
		db := database.GetDB()
		err := db.View(func(tx *bolt.Tx) error {
			//get task bucket
			b := tx.Bucket([]byte("tasks"))
			if b == nil {
				return errors.New("Bucket does not exist")
			}
			c := b.Cursor()

			count := 0
			var taskMap map[string]interface{}
			for k, v := c.First(); k != nil; k, v = c.Next() {
				count++
				_ = json.Unmarshal(v, &taskMap)
				if taskMap["completed"].(bool) {
					fmt.Printf("%d. ID: %s | Task: %s\n", count, k, taskMap["task"])
				}
			}
			return nil
		})
		if err != nil {
			log.Fatal(err)
		}
	},
}
