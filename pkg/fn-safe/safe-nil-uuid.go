package fnSafe

import uuid "github.com/satori/go.uuid"

// -----
// Безапосное извлечение значения из ссылки
// -----
func SafeNilUuid(value *uuid.UUID) uuid.UUID {
	// ..
	if value == nil {
		return uuid.FromStringOrNil("")
	}

	return *value
}
