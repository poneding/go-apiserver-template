package util

import (
	"strings"

	"github.com/google/uuid"
)

// UUIDOption is the option for uuid
type UUIDOption struct {
	// RemoveHyphen removes hyphen from uuid
	RemoveHyphen bool
	// UpperCase converts uuid to upper case
	UpperCase bool
}

// NewUUID returns a new uuid
func NewUUID(opts ...UUIDOption) string {
	uuid := uuid.New().String()
	for _, opt := range opts {
		if opt.RemoveHyphen {
			uuid = strings.ReplaceAll(uuid, "-", "")
		}

		if opt.UpperCase {
			uuid = strings.ToUpper(uuid)
		}
	}
	return uuid
}
