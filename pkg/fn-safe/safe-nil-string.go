package fnSafe

// -----
// Безапосное извлечение значения из ссылки
// -----
func SafeNilString(value *string) string {
	// ..
	if value == nil {
		return ""
	}

	return *value
}
