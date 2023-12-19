package rocksdbKit

import "github.com/linxGnu/grocksdb"

var (
	// NewDefaultReadOptions
	/*
		PS: 搭配 DB.Get() 使用.
	*/
	NewDefaultReadOptions func() *grocksdb.ReadOptions = grocksdb.NewDefaultReadOptions

	// NewDefaultWriteOptions
	/*
		PS: 搭配 DB.Put() 使用.
	*/
	NewDefaultWriteOptions func() *grocksdb.WriteOptions = grocksdb.NewDefaultWriteOptions
)

func NewDefaultDBOptions() *grocksdb.Options {
	opts := grocksdb.NewDefaultOptions()

	opts.SetCreateIfMissing(true)
	opts.SetCompression(grocksdb.SnappyCompression)

	return opts
}
