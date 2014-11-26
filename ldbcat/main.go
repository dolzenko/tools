package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/dolzenko/tools/ldbcat/Godeps/_workspace/src/github.com/syndtr/goleveldb/leveldb"
	"github.com/dolzenko/tools/ldbcat/Godeps/_workspace/src/github.com/syndtr/goleveldb/leveldb/filter"
	"github.com/dolzenko/tools/ldbcat/Godeps/_workspace/src/github.com/syndtr/goleveldb/leveldb/opt"
)

var tableOpts = opt.Options{
	Compression: opt.SnappyCompression,
	Strict:      opt.NoStrict,
	WriteBuffer: 8 * opt.MiB,
	Filter:      filter.NewBloomFilter(10),
}

func main() {
	args := os.Args
	if len(args) < 2 {
		abortOn(errors.New("invalid arguments, require DBFOLDER"))
	}

	db, err := leveldb.OpenFile(args[1], &tableOpts)
	abortOn(err)
	defer db.Close()

	iter := db.NewIterator(nil, nil)
	for iter.Next() {
		fmt.Printf("%s\t%s\n", string(iter.Key()), string(iter.Value()))
	}
	iter.Release()
	abortOn(iter.Error())
}

func abortOn(err error) {
	if err != nil {
		panic(err)
	}
}
