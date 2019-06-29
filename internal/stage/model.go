package stage

import (
	"time"
)

type Stage struct {
	ID       int64 `gorm:"primary_key"`
	ServerID int64 `gorm:"column:server_id"`
	Stage    int64 `gorm:"column:stage"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  *time.Time
}
