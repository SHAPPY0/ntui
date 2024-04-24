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
	current := time.Now()
	hours := current.Sub(dateTime).Hours()
	hours, minFraction := math.Modf(hours)
	minutes := minFraction * 60
	minutes, secFraction := math.Modf(minutes)
	seconds := secFraction * 60
	var strBuilder strings.Builder
	if hours > 0 {
		h := fmt.Sprintf("%sh", fmt.Sprint(hours))
		strBuilder.WriteString(h)
	}
	m := fmt.Sprintf("%sm", fmt.Sprint(minutes))
	if minutes < 9 {
		m = fmt.Sprintf("0%sm", fmt.Sprint(minutes))
	}
	strBuilder.WriteString(m)
	
	s := fmt.Sprintf("%ss", fmt.Sprint(math.Floor(seconds)))
	if seconds < 9 {
		s = fmt.Sprintf("0%ss", fmt.Sprint(math.Floor(seconds)))
	}
	strBuilder.WriteString(s)
	return strBuilder.String()
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