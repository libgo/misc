package misc

import (
	"github.com/google/uuid"
)

// UUID returns random generated UUID.
func UUID() string {
	return uuid.New().String()
}
