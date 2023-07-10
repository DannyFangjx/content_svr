package utils

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"github.com/kardianos/osext"
	"net"
	"reflect"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"time"
	"unicode"
)

func MD5(str string) string {
	data := []byte(str) //切片
	has := md5.Sum(data)
	md5str := fmt.Sprintf("%x", has) //将[]byte转成16进制
	return md5str
}

func GetCurFuncName() string {
	pc, _, line, _ := runtime.Caller(1)
	return runtime.FuncForPC(pc).Name() + ":" + strconv.Itoa(line)
}

func HashCode(s string) int32 {
	var h int32
	for i := 0; i < len(s); i++ {
		h = 31*h + int32(s[i])
	}
	return h
}

func GetCurTsMs() int64 {
	return time.Now().UnixNano() / 1e6
}

// 输入 "/11/16/"  输出 []int64{11,16}
func GetIdListFromFmtPath(input string) []int64 {
	ret := make([]int64, 0)
	if len(input) == 0 {
		return ret
	}
	arrStr := strings.Split(input, "/")
	for _, arr := range arrStr {
		if arr == "" {
			continue
		}
		intArr, err := strconv.Atoi(arr)
		if err != nil {
			continue
		}
		ret = append(ret, int64(intArr))
	}
	return ret
}

// 获取 bin conf log 同级的目录。 eg: /data/content_svr/build
func GetWorkPath() (string, error) {
	folderPath, err := osext.ExecutableFolder() // folderPath = data/content_svr/build/bin
	if err != nil {
		return "", err
	}
	return folderPath[0 : len(folderPath)-len("/bin")], nil
}

func GetLocalIp() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println(err)
		return ""
	}
	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}

func IsNil(i interface{}) bool {
	if i == nil {
		return true
	}
	vi := reflect.ValueOf(i)
	if vi.Kind() == reflect.Ptr {
		return vi.IsNil()
	}
	return false
}

func GetNowUnix() uint64 {
	nowTimeUnix := time.Now().Unix()
	return uint64(nowTimeUnix)
}

func GetNowUnixNano() uint64 {
	nowTimeUnix := time.Now().UnixNano() / 1e6
	return uint64(nowTimeUnix)
}

func InterfaceToString(inter interface{}) string {
	var result string
	switch inter := inter.(type) {
	case string:
		result = inter
	case int:
		result = strconv.Itoa(inter)
	case float64:
		result = strconv.FormatFloat(inter, 'f', -1, 64)
		//result = fmt.Sprintf("%v", inter)
	default:
		result = ""
	}
	return result
}

func MapToJsonString(item map[string]interface{}) (string, error) {
	buf, err := json.Marshal(item)
	if err != nil {
		panic(err)
	}
	return string(buf), err
}

func StructToJsonString(item interface{}) string {
	if item == nil {
		return ""
	}
	buf, err := json.Marshal(item)
	if err != nil {
		return ""
	}
	return string(buf)
}

func JSONToMap(str string) (map[string]interface{}, error) {

	var tempMap map[string]interface{}

	err := json.Unmarshal([]byte(str), &tempMap)
	if err != nil {
		return nil, err
	}
	return tempMap, nil
}

func Int64ListToSetList(itemList []uint64) []uint64 {
	if len(itemList) <= 1 {
		return itemList
	}
	itemMap := make(map[uint64]bool, len(itemList))
	for _, item := range itemList {
		itemMap[item] = true
	}
	result := make([]uint64, 0)
	for key := range itemMap {
		result = append(result, key)
	}
	return result
}
func Uint64ListToSetList(itemList []uint64) []uint64 {
	if len(itemList) <= 1 {
		return itemList
	}
	itemMap := make(map[uint64]bool, len(itemList))
	for _, item := range itemList {
		itemMap[item] = true
	}
	result := make([]uint64, 0)
	for key := range itemMap {
		result = append(result, key)
	}
	return result
}

func StringListToSetList(itemList []string) []string {
	if len(itemList) <= 1 {
		return itemList
	}
	itemMap := make(map[string]bool, len(itemList))
	for _, item := range itemList {
		itemMap[item] = true
	}
	result := make([]string, 0)
	for key := range itemMap {
		result = append(result, key)
	}
	return result
}

func Uint64ToInt64(itemList []uint64) []int64 {
	result := make([]int64, 0)
	for _, item := range itemList {
		result = append(result, int64(item))
	}
	return result
}

func Int64ToUint64(itemList []int64) []uint64 {
	result := make([]uint64, 0)
	for _, item := range itemList {
		result = append(result, uint64(item))
	}
	return result
}

func HasElement(item interface{}, itemList []interface{}) bool {
	if len(itemList) <= 0 {
		return false
	}
	for _, temItem := range itemList {
		a := reflect.ValueOf(item).Interface()
		b := reflect.ValueOf(temItem).Interface()
		if a == b {
			return true
		}
	}
	return false
}

func HasUint64Element(item uint64, itemList []uint64) bool {
	if len(itemList) <= 0 {
		return false
	}
	for _, temItem := range itemList {
		if item == temItem {
			return true
		}
	}
	return false
}

func HasInt64Element(item int64, itemList []int64) bool {
	if len(itemList) <= 0 {
		return false
	}
	for _, temItem := range itemList {
		if item == temItem {
			return true
		}
	}
	return false
}

func HasInt32Element(item int32, itemList []int32) bool {
	if len(itemList) <= 0 {
		return false
	}
	for _, temItem := range itemList {
		if item == temItem {
			return true
		}
	}
	return false
}

func HasUint8Element(item uint8, itemList []uint8) bool {
	if len(itemList) <= 0 {
		return false
	}
	for _, temItem := range itemList {
		if item == temItem {
			return true
		}
	}
	return false
}

func HasStringElement(item string, itemList []string) bool {
	if len(itemList) <= 0 {
		return false
	}
	for _, temItem := range itemList {
		if item == temItem {
			return true
		}
	}
	return false
}

func GetChanges(itemBase interface{}, itemTarget interface{}) (before map[string]interface{}, after map[string]interface{}) {
	changesIgnoreField := []interface{}{"Ctime", "Mtime", "Operator"}
	baseFields := reflect.TypeOf(itemBase)
	baseValues := reflect.ValueOf(itemBase)
	targetValues := reflect.ValueOf(itemTarget)
	before, after = map[string]interface{}{}, map[string]interface{}{}
	for k := 0; k < baseFields.NumField(); k++ {
		key := baseFields.Field(k).Name
		if HasElement(key, changesIgnoreField) {
			continue
		}
		baseValue := baseValues.Field(k).Interface()
		targetValue := targetValues.Field(k).Interface()
		if IsNil(baseValue) && IsNil(targetValue) {
			continue
		}
		if IsNil(targetValue) {
			continue
		}
		if baseValue == targetValue {
			continue
		}
		before[key] = baseValue
		after[key] = targetValue
	}
	return before, after
}

func RemoveRepeatElem(arr []string) (newArr []string) {
	if len(arr) < 2 {
		return arr
	}

	newArr = make([]string, 0)
	for i := 0; i < len(arr); i++ {
		repeat := false
		for j := i + 1; i < len(arr) && j < len(arr); j++ {
			if arr[i] == arr[j] {
				repeat = true
				break
			}
		}
		if !repeat {
			newArr = append(newArr, arr[i])
		}
	}

	return
}

func LastFileNameByPath(filePath string) (fileName string) {
	_fileNameList := strings.Split(filePath, "/")
	fileName = _fileNameList[len(_fileNameList)-1]
	return
}

// XXXX.jpg  ret jpg
func GetOriginFileType(fileName string) string {
	if fileName == "" {
		return ""
	}
	arr := strings.Split(fileName, ".")
	arrLen := len(arr)
	if arrLen >= 2 {
		typeStr := arr[arrLen-1] // 取用最后一个
		return strings.ToLower(typeStr)
	} else {
		return ""
	}
}

// 0将会跳过
func GetIntListMaxMin(inputSrc []int) (int, int) {
	input := make([]int, 0)
	for _, i := range inputSrc {
		if i != 0 {
			input = append(input)
		}
	}

	if len(input) == 0 {
		return 0, 0
	}
	max := input[0]
	min := input[0]
	for _, i := range input {
		if i == 0 {
			continue
		}
		if i > max {
			max = i
		}
		if i < min {
			min = i
		}
	}
	return max, min
}

// 是否含有中文(包括简体和繁体)
func IsChineseChar(str string) bool {
	for _, r := range str {
		if unicode.Is(unicode.Scripts["Han"], r) || (regexp.MustCompile("[\u3002\uff1b\uff0c\uff1a\u201c\u201d\uff08\uff09\u3001\uff1f\u300a\u300b]").MatchString(string(r))) {
			return true
		}
	}
	return false
}

func BoolToInt(i bool) int {
	if i {
		return 1
	} else {
		return 0
	}
}
