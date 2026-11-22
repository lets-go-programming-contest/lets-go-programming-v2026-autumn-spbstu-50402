package codetoint

import (
	"fmt"
	"strconv"
	"strings"
)

func CodeToInt(raw string) (int, error) {
	trimmed := strings.TrimSpace(raw)
	if trimmed == "" {
		return 0, nil
	}

	numCode, err := strconv.Atoi(trimmed)
	if err != nil {
		return 0, fmt.Errorf("parse num code failed: %w", raw, err)
	}

	return numCode, nil
}
