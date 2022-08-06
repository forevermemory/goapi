package utils

import (
	"os"
	"regexp"
	"strings"
)

var (
	PACKAGE_CONFIG string = "\"commonapi/config\""
)

var InternelControllerPackages = map[string]int{
	`"strconv"`:                  1,
	`"github.com/gin-gonic/gin"`: 1,
}

var InternelModelPackages = map[string]int{
	`"time"`:         1,
	`"gorm.io/gorm"`: 1,
	`"fmt"`:          1,
	PACKAGE_CONFIG:   1,
}

// 程序工作目录
var Workdir string = func() string {
	p, _ := os.Getwd()
	return p
}()

func ToGoUpper(s string) string {
	s1 := strings.Split(s, "_")
	s2 := make([]string, 0)
	for _, v := range s1 {
		s2 = append(s2, ToFirstUpper(v))
	}

	return strings.Join(s2, "")
}

// ToFirstUpper 把第一个单词变成大写
func ToFirstUpper(s string) string {
	return strings.ToUpper(s[0:1]) + s[1:]
}

func GetImportPositionFromString(s string) []int {
	reg, _ := regexp.Compile(`(?m)^import\s+`)

	res := reg.FindStringIndex(s) // 返回[start end]
	if len(res) == 0 {
		return nil
	}

	var i = 1

	// 两种情况
	a := s[res[1]]

	// 1. import "xxx"
	if a == '"' {
		for {
			b := s[res[1]+i]
			if b == '"' {
				break
			}
			i += 1
		}

	} else {
		// 2. import ()
		for {
			b := s[res[1]+i]
			if b == ')' {
				break
			}
			i += 1
		}
	}
	return []int{res[0], res[1] + i}
}

func GetStructPositionFromString(s string, structName string) []int {

	reg, _ := regexp.Compile(`(?m)^type\s+` + structName + `\s+struct\s+\{`)

	res := reg.FindStringIndex(s) // 返回[start end]
	var i = 1
	var count = 1

	// s2 := []rune(s)
	for {
		b := s[res[1]+i]
		if b == '{' {
			count += 1
		}
		if b == '}' {
			count -= 1
		}
		i += 1
		if count == 0 {
			// log.Println(i)
			break
		}
	}

	return []int{res[0], res[1] + i}
}

func GetStructFromString(s string, structName string) string {

	reg, _ := regexp.Compile(`(?m)^type\s+` + structName + `\s+struct\s+\{`)

	res := reg.FindStringIndex(s) // 返回[start end]
	var i = 1
	var count = 1

	// s2 := []rune(s)
	for {
		b := s[res[1]+i]
		if b == '{' {
			count += 1
		}
		if b == '}' {
			count -= 1
		}
		i += 1
		if count == 0 {
			// log.Println(i)
			break
		}
	}

	return s[res[0] : res[1]+i]
}
