package goe

func BoolOrDefault(value bool, defaultValue bool) bool {
	if !value {
		return defaultValue
	}

	return value
}

func BoolOrElse(value bool, defaultFunc func() (bool, error)) (bool, error) {
	if !value {
		return defaultFunc()
	}

	return value, nil
}

func IntOrDefault(value int, defaultValue int) int {
	if value == 0 {
		return defaultValue
	}

	return value
}

func IntOrElse(value int, defaultFunc func() (int, error)) (int, error) {
	if value == 0 {
		return defaultFunc()
	}

	return value, nil
}

func UintOrDefault(value uint, defaultValue uint) uint {
	if value == 0 {
		return defaultValue
	}

	return value
}

func UintOrElse(value uint, defaultFunc func() (uint, error)) (uint, error) {
	if value == 0 {
		return defaultFunc()
	}

	return value, nil
}

func Float32OrDefault(value float32, defaultValue float32) float32 {
	if value == 0 {
		return defaultValue
	}

	return value
}

func Float32OrElse(value float32, defaultFunc func() (float32, error)) (float32, error) {
	if value == 0 {
		return defaultFunc()
	}

	return value, nil
}

func Float64OrDefault(value float64, defaultValue float64) float64 {
	if value == 0 {
		return defaultValue
	}

	return value
}

func Float64OrElse(value float64, defaultFunc func() (float64, error)) (float64, error) {
	if value == 0 {
		return defaultFunc()
	}

	return value, nil
}

func StringOrDefault(value string, defaultValue string) string {
	if value == "" {
		return defaultValue
	}

	return value
}

func StringOrElse(value string, defaultFunc func() (string, error)) (string, error) {
	if value == "" {
		return defaultFunc()
	}

	return value, nil
}

func BytesOrDefault(value []byte, defaultValue []byte) []byte {
	if value == nil {
		return defaultValue
	}

	return value
}

func BytesOrElse(value []byte, defaultFunc func() ([]byte, error)) ([]byte, error) {
	if value == nil {
		return defaultFunc()
	}

	return value, nil
}
