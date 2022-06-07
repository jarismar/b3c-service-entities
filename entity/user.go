package entity

import "github.com/google/uuid"

type User struct {
	Id           int64
	UUID         uuid.UUID
	ExternalUUID string
	UserName     string
}
