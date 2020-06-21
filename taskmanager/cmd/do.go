package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strconv"

	"github.com/Sarkk97/gophercises/taskmanager/database"
	"github.com/boltdb/bolt"

	"github.com/spf13/cobra"
)

var doCmd = &cobra.Command{
	Use:   "do",
	Short: "marks a task as completed",
	Run: func(cmd *cobra.Command, args []string) {
		db := database.GetDB()
		err := db.Update(func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte("tasks"))
			if b == nil {
				return errors.New("Bucket does not exist")
			}
			var taskMap map[string]interface{}
			for _, val := range args {
				rawTaskMap := b.Get([]byte(val))
				if rawTaskMap == nil {
					fmt.Printf("No task id with value %s\n", val)
					continue
				}

				// b.Delete([]byte(val))
				_ = json.Unmarshal(rawTaskMap, &taskMap)
				if taskMap["completed"].(bool) {
					continue
				}
				taskMap["completed"] = true
				taskMap["completed_at"] = "dummy"

				updatedTask, _ := json.Marshal(taskMap)
				b.Put([]byte(val), updatedTask)
				fmt.Printf("You have completed the \"%s\" task\n", taskMap["task"])
			}
			return nil
		})
		if err != nil {
			log.Fatal(err)
		}
	},
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires task id(s) argument")
		}
		for _, v := range args {
			_, err := strconv.Atoi(v)
			if err != nil {
				return fmt.Errorf("%v is not a valid id", v)
			}
		}
		return nil
	},
}
