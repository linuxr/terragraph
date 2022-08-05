package utils

import (
	"github.com/google/uuid"
)

func GenUUID() string {
	u4 := uuid.New()
	return u4.String()
}
