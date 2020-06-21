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

var rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "delete a task",
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
	Run: func(cmd *cobra.Command, args []string) {
		db := database.GetDB()
		err := db.Update(func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte("tasks"))
			if b == nil {
				return errors.New("Bucket does not exist")
			}
			for _, v := range args {
				rawTaskMap := b.Get([]byte(v))
				if rawTaskMap == nil {
					fmt.Printf("No task id with value %s\n", v)
					continue
				}
				err := b.Delete([]byte(v))
				if err != nil {
					return err
				}
				var taskMap map[string]interface{}
				_ = json.Unmarshal(rawTaskMap, &taskMap)
				fmt.Printf("Removed task: \"%s\" from the db\n", taskMap["task"])
			}
			return nil
		})
		if err != nil {
			log.Fatal(err)
		}
	},
}
