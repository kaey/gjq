package main

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

func JSON(v interface{}) (string, error) {
	b := new(strings.Builder)
	if err := json.NewEncoder(b).Encode(v); err != nil {
		return "", err
	}

	return b.String(), nil
}

func Duration(v interface{}) (time.Duration, error) {
	s, ok := v.(string)
	if !ok {
		return 0, fmt.Errorf("duration: expected string, got %v with type %t", v, v)
	}

	return time.ParseDuration(s)
}

var funcMap = map[string]interface{}{
	"duration": Duration,
	"json":     JSON,
}
