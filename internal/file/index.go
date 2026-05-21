package file

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func GetEnvToBool(key string) bool {
	var value = GetEnv(key)

	if value == "" {
		return false
	}

	boolValue, error := strconv.ParseBool(value)

	if error != nil {
		return false
	}

	return boolValue
}

func GetEnv(key string) string {
	file, err := os.Open(".env")
	if err != nil {
		return ""
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}

		currentKey := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		if currentKey == key {
			return value
		}
	}

	if err := scanner.Err(); err != nil {
		return ""
	}

	return ""
}
