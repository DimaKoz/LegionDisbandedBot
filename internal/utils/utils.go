package utils

func AppendArgs(target []string, key string, value string) []string {
	if value != "" {
		target = append(target, key)
		target = append(target, value)
	}

	return target
}
