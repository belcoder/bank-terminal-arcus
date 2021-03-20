package fnSafe

// -----
// Безапосное извлечение значения из ссылки
// -----
func SafeNilFloat64(value *float64) float64 {
	// ..
	if value == nil {
		return 0
	}

	return *value
}
