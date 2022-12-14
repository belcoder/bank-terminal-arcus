package contains

import uuid "github.com/satori/go.uuid"

func Int(a []int, x int) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}

func Uuid(a []uuid.UUID, x uuid.UUID) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}

func String(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}
