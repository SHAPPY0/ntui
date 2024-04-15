package utils

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
	"math"
	"path/filepath"
	"encoding/json"
)

const (
	DefaultDirMod os.FileMode = 0755
	DefaultFileMod os.FileMode = 0600
)

func IntToStr(value int) string {
	return strconv.Itoa(value)
}

func StrToInt(s string) int {
	v, err := strconv.Atoi(s)
	if err != nil {
		return 0 
	}
	return v
}

func IntToUint64(v int) uint64 {
	return uint64(v)
}

func Split(str, sep string) []string {
	return strings.Split(str, sep)
}

func Stringify(st interface{}) string {
	str, err := json.Marshal(st)
	if err != nil {
		return ""
	}
	return string(str)
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

func EnsureDirPath(path string, mod os.FileMode) error {
	pathDir := filepath.Dir(path)
	if _, err := os.Stat(pathDir); os.IsNotExist(err) {
		if err = os.MkdirAll(pathDir, mod); err != nil {
			return err
		}
	}
	return nil
}