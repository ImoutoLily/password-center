package models

import (
	"time"
)

type User struct {
	username         string
	creationDateTime time.Time
	updatedDateTime  time.Time
	Password         []Password
}
