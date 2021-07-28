package storage

import (
	"github.com/nats-io/nats.go"
	"gorm.io/gorm"
)

func ListenStorage(db *gorm.DB, nc *nats.Conn) error {
	return nil
}
