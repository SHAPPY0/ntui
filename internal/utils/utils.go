package utils

import (
	"strconv"
	"strings"
	"time"
)

func IntToStr(value int) string {
	return strconv.Itoa(value)
}

func GetID(id string) string {
	if id != nil {
		idParts := strings.Split(id, "-")
		return idParts[0]
	} else {
		return id
	}
}

func ToCapitalize(str string) string {
	if str != "" {
		return strings.Title(strings.ToLower(str))
	} else {
		return str
	}
}

func DateTimeToStr(dateTime time.Time) string {
	return dateTime.Format("2006-01-02 15:04:05")
}

func FormatMemoryUsage(value int) int {
	if value == 0 {
		return value
	}
	return value / 1024 / 1024
}

func SafeDeref[T any](p *T) T {
	if p == nil {
		var v T
		return v
	}
	return *p
}