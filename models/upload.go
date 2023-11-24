package models

import (
	"time"

	"github.com/google/uuid"
)

func GenerateRandomFilename() string {
	timestamp := time.Now().Format("20060102150405")
	uuidStr := uuid.New().String()
	return timestamp + "_" + uuidStr
}
