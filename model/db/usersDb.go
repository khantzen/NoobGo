package db

import "time"

type UserDb struct {
	Id           int
	Username     string
	CreationDate time.Time
}
