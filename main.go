package main

import (
	"crypto/rand"
	"log/slog"
	"os"
	"path"

	"github.com/dgraph-io/badger/v4"
	"github.com/google/uuid"
)

func main() {
	location := path.Join(os.TempDir(), uuid.NewString())
	bo := badger.DefaultOptions(location).
		WithValueThreshold(1 << 10).  // 1KB
		WithValueLogFileSize(4 << 20) // 4MB
	if db, err := badger.Open(bo); err != nil {
		slog.Error("DB OPEN", slog.Any("error", err))
		return
	} else {
		for range 1000 {
			if err := db.Update(func(txn *badger.Txn) error {
				key := make([]byte, 200)
				val := make([]byte, 2<<10) // 2KB
				for range 1024 {
					rand.Read(key)
					rand.Read(val)
					if err := txn.Set(key, val); err != nil {
						return err
					}
				}
				return nil
			}); err != nil {
				slog.Error("update error", slog.Any("error", err))
				return
			}
		}
		if err := db.Close(); err != nil {
			slog.Error("DB CLOSE", slog.Any("error", err))
		}
	}
}
