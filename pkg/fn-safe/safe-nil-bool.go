package fnSafe

// -----
// Безапосное извлечение значения из ссылки
// -----
func SafeNilBool(value *bool) bool {
	// ..
	if value == nil {
		return false
	}

	return *value
}
