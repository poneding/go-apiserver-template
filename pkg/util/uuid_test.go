package util

import "testing"

func TestNewUUID(t *testing.T) {
	t.Log(NewUUID())

	t.Log(NewUUID(UUIDOption{
		RemoveHyphen: true,
	}))

	t.Log(NewUUID(UUIDOption{
		RemoveHyphen: true,
		UpperCase:    true,
	}))

	t.Log(NewUUID(UUIDOption{
		RemoveHyphen: true,
	}, UUIDOption{
		UpperCase: true,
	}))
}
