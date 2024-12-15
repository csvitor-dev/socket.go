package utils

import (
	"fmt"
	"time"
)

func LogDate() string {
	now := time.Now()
	hour, min, sec := now.Clock()

	return fmt.Sprintf("%vh%vmin%vs", hour, min, sec)
}