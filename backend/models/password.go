package models

import "time"

type Password struct {
	encryptedPassword string
	creationDateTime  time.Time
	updatedDateTime   time.Time
}
