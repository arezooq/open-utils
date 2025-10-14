package uuid

import "github.com/google/uuid"

func UUIDString() string {
	return uuid.New().String()
}
