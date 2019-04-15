package misc

import (
	uuid "github.com/satori/go.uuid"
)

// UUID returns random generated UUID.
func UUID() string {
	return uuid.Must(uuid.NewV4()).String()
}
