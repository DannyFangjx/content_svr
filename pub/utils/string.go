package utils

import (
	"fmt"
	"math/rand"
	"regexp"
	"strconv"
	"strings"

	"github.com/pkg/errors"

	errors2 "content_svr/pub/errors"
)

var (
	phoneRegexp           = regexp.MustCompile(`^(\\+86|\\+852|\\+886|\\+62|\\+60|\\+84|\\+63|\\+66|\\+82|\\+65|\\+55|\\+91|\\+52)(\\d.*)`)
	InviteCodeLetterRunes = []rune("123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnpqrstuvwxyz")
)

func GenRandomCkCode() string {
	buf := make([]rune, 6)
	for i := range buf {
		buf[i] = InviteCodeLetterRunes[rand.Intn(len(InviteCodeLetterRunes))]
	}
	return string(buf)
}

func MaskString(str string) string {
	length := len([]rune(str))

	if length <= 5 {
		return string([]rune(str)[:1]) + "****"
	} else if length <= 10 {
		return string([]rune(str)[:1]) + "****" + string([]rune(str)[length-1:])
	} else {
		return string([]rune(str)[:2]) + "****" + string([]rune(str)[length-2:])
	}
}

func Str2Int64s(str, sep string) ([]int64, error) {
	if str == "" {
		return []int64{}, nil
	}
	elements := strings.Split(str, sep)
	items := make([]int64, 0, len(elements))
	for _, ele := range elements {
		num, err := strconv.ParseInt(ele, 10, 64)
		if err != nil {
			return nil, errors.Wrap(err, errors2.GetStack())
		}
		items = append(items, num)
	}
	return items, nil
}

func JoinInt64s(items []int64, sep string) string {
	elements := make([]string, 0, len(items))
	for _, i := range items {
		elements = append(elements, strconv.FormatInt(i, 10))
	}
	return strings.Join(elements, sep)
}

func JoinInt(items []int, sep string) string {
	elements := make([]string, 0, len(items))
	for _, i := range items {
		elements = append(elements, strconv.Itoa(i))
	}
	return strings.Join(elements, sep)
}

func StrIn(target string, items ...string) bool {
	for _, item := range items {
		if item == target {
			return true
		}
	}
	return false
}

func Str2Int64(str string) (int64, error) {
	num, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return 0, errors.Wrap(err, errors2.GetStack())
	}
	return num, nil
}

func Str(s string) *string {
	return &s
}

func MaskPhone(phone string) string {
	if len(phone) < 7 {
		return phone
	}
	matchs := phoneRegexp.FindStringSubmatch(phone)
	prefix := phone[:3]
	if len(matchs) > 0 {
		prefix = matchs[1]
	}
	return prefix + "****" + phone[len(phone)-4:]
}

func MaskEmail(email string) string {
	strs := strings.Split(email, "@")
	if len(strs) != 2 {
		return email
	}
	prefix := strs[0]
	if len(prefix) > 3 {
		prefix = prefix[:3]
	}
	return prefix + "****@" + strs[1]
}

func StrSet(strs []string) []string {
	strMap := make(map[string]bool)
	for _, _s := range strs {
		strMap[_s] = true
	}
	result := make([]string, 0, len(strMap))
	for k := range strMap {
		result = append(result, k)
	}
	return result
}

func FormatFloat(f interface{}) string {
	str := fmt.Sprintf("%f", f)
	hasPoint := false
	for _, s := range str {
		if s == '.' {
			hasPoint = true
			break
		}
	}
	if !hasPoint {
		return str
	}
	var i int
	for i = len(str) - 1; i >= 0; i-- {
		if str[i] != '0' {
			break
		}
	}
	if str[i] == '.' {
		return str[:i]
	}
	return str[:i+1]
}

// 转化失败则返回0
func String2int64(src string) int {
	verInt64, _ := strconv.Atoi(src)
	return verInt64
}

// 123123,3435435 转化为 []int64{123123,3435435}
func Str2Int64List(input string) []int64 {
	parts := strings.Split(input, ",")
	ints := make([]int64, 0)
	for _, part := range parts {
		val, err := strconv.ParseInt(strings.Trim(part, " "), 10, 64)
		if err != nil {
			continue
		}
		ints = append(ints, val)
	}
	return ints
}

func Str2StringList(input string) []string {
	parts := strings.Split(input, ",")
	retStrList := make([]string, 0)
	for _, part := range parts {
		val := strings.Trim(part, " ")
		if val == "" {
			continue
		}
		retStrList = append(retStrList, val)
	}
	return retStrList
}
