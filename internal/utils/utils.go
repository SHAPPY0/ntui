package utils

import (
	"fmt"
	"strconv"
	"strings"
	"time"
	"math"
)

func IntToStr(value int) string {
	return strconv.Itoa(value)
}

func GetID(id string) string {
	if id != "" {
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

func DateTimeDiff(dateTime time.Time) string {
	Current := time.Now()
	Hours := Current.Sub(dateTime).Hours()
	Hours, MinFraction := math.Modf(Hours)
	Minutes := MinFraction * 60
	Minutes, SecFraction := math.Modf(Minutes)
	Seconds := SecFraction * 60
	Diff := fmt.Sprintf("%shrs %smins %ssecs ago", fmt.Sprint(Hours), fmt.Sprint(Minutes), fmt.Sprint(math.Floor(Seconds)))
	return Diff
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