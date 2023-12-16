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

	dbOpts := grocksdb.NewDefaultOptions()
	logrus.Infof("CreateIfMissing: %t", dbOpts.CreateIfMissing())
	dbOpts.SetCreateIfMissing(true)

	compressionType := dbOpts.GetCompression()
	switch compressionType {
	case grocksdb.NoCompression:
		logrus.Info("NoCompression")
	case grocksdb.SnappyCompression:
		logrus.Info("SnappyCompression")
	case grocksdb.ZLibCompression:
		logrus.Info("ZLibCompression")
	case grocksdb.Bz2Compression:
		logrus.Info("Bz2Compression")
	case grocksdb.LZ4Compression:
		logrus.Info("LZ4Compression")
	case grocksdb.LZ4HCCompression:
		logrus.Info("LZ4HCCompression")
	case grocksdb.XpressCompression:
		logrus.Info("XpressCompression")
	case grocksdb.ZSTDCompression:
		logrus.Info("ZSTDCompression")
	default:
		logrus.Fatal("UnknownCompression")
	}

	keyStr := []byte("key")
	valueStr := []byte(idKit.NewULID())

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
