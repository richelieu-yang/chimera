package main

import (
	"github.com/richelieu-yang/chimera/v2/src/idKit"
	"github.com/richelieu-yang/chimera/v2/src/log/logrusKit"
	"github.com/sirupsen/logrus"
	"github.com/tidwall/buntdb"
)

func init() {
	logrusKit.MustSetUp(nil)
}

func main() {
	// Open the data.db file. It will be created if it doesn't exist.
	db, err := buntdb.Open("_test.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	key := "key"
	value := idKit.NewULID()

	// (1) To set a value you must open a read/write transaction:
	err = db.Update(func(tx *buntdb.Tx) error {
		_, _, err := tx.Set(key, value, nil)
		return err
	})
	if err != nil {
		panic(err)
	}
	logrus.Info("Manager to set.")

	// (2) To get the value
	var val string
	err = db.View(func(tx *buntdb.Tx) error {
		var err error
		val, err = tx.Get(key)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
	logrus.WithField("val", val).Info("Manager to get.")
}
