package database

import (
	"fmt"
	"log"
	"os"

	"github.com/boltdb/bolt"
	homedir "github.com/mitchellh/go-homedir"
)

var db *bolt.DB

func init() {
	var err error
	dbLocation := configureDBLocation()
	db, err = bolt.Open(dbLocation, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}

	//Create task bucket if not exist
	err = db.Update(func(tx *bolt.Tx) error {
		// _ = tx.DeleteBucket([]byte("tasks"))
		_, err := tx.CreateBucketIfNotExists([]byte("tasks"))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
}

func configureDBLocation() string {
	dir, _ := homedir.Expand("~/.boltdb")
	if _, err := os.Lstat(dir); os.IsNotExist(err) {
		os.Mkdir(dir, 0700)
	}
	return dir + "/my.db"
}

//GetDB returns the boltDB connection
func GetDB() *bolt.DB {
	return db
}
