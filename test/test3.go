package main

import (
	"github.com/linxGnu/grocksdb"
	"github.com/richelieu-yang/chimera/v2/src/idKit"
	"github.com/richelieu-yang/chimera/v2/src/log/logrusKit"
	"github.com/sirupsen/logrus"
)

func init() {
	logrusKit.MustSetUp(nil)
}

func main() {
	dirPath := "test_dir"
	keyStr := []byte("key")
	valueStr := []byte(idKit.NewULID())

	dbOpts := grocksdb.NewDefaultOptions()
	dbOpts.SetCreateIfMissing(true)
	dbOpts.SetCompression(grocksdb.SnappyCompression)

	db, err := grocksdb.OpenDb(dbOpts, dirPath)
	if err != nil {
		logrus.Fatal(err)
	}
	defer db.Close()

	writeOpts := grocksdb.NewDefaultWriteOptions()
	err = db.Put(writeOpts, keyStr, valueStr)
	if err != nil {
		logrus.Fatal(err)
	}

	readOpts := grocksdb.NewDefaultReadOptions()
	value, err := db.Get(readOpts, keyStr)
	if err != nil {
		logrus.Fatal(err)
	}
	defer value.Free()
	logrus.Infof("Read value from database: [%s]", string(value.Data()))
}
