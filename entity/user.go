package entity

import "github.com/google/uuid"

type User struct {
	UserID           int64
	UserUUID         uuid.UUID
	UserExternalUUID string
	UserName         string
}
