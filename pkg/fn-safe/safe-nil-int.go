package fnSafe

// -----
// Безапосное извлечение значения из ссылки
// -----
func SafeNilInt(value *int) int {
	// ..
	if value == nil {
		return 0
	}

	return *value
}
