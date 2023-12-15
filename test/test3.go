package main

import (
	"fmt"
	"github.com/linxGnu/grocksdb"
)

func main() {
	dirPath := "test_dir"

	dbOpts := grocksdb.NewDefaultOptions()
	dbOpts.SetCreateIfMissing(true)
	db, err := grocksdb.OpenDb(dbOpts, dirPath)
	if err != nil {
		fmt.Println("Error opening database: ", err)
		return
	}
	defer db.Close()

	writeOpts := grocksdb.NewDefaultWriteOptions()
	readOpts := grocksdb.NewDefaultReadOptions()
	err = db.Put(writeOpts, []byte("key"), []byte("value"))
	if err != nil {
		fmt.Println("Error writing to database: ", err)
		return
	}
	value, err := db.Get(readOpts, []byte("key"))
	if err != nil {
		fmt.Println("Error reading from database: ", err)
		return
	}
	defer value.Free()
	fmt.Println("Read value from database: ", string(value.Data()))
}
