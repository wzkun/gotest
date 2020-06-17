package decode

import (
	"math/rand"
	"strconv"
	"strings"
	"time"
)

// ToStringArray function.
func ToStringArray(objs ...interface{}) []string {
	records := make([]string, 0, len(objs))
	for _, obj := range objs {
		js, _ := JSON.Marshal(obj)
		record := string(js)
		records = append(records, record)
	}
	return records
}

// ToString function.
func ToString(obj interface{}) string {
	js, _ := JSON.Marshal(obj)
	record := string(js)
	return record
}

// RandomString function.
func RandomString(length int) string {
	// 48 ~ 57 数字
	// 65 ~ 90 A ~ Z
	// 97 ~ 122 a ~ z
	// 一共62个字符，在0~61进行随机，小于10时，在数字范围随机，
	// 小于36在大写范围内随机，其他在小写范围随机
	rand.Seed(time.Now().UnixNano())
	result := make([]string, 0, length)
	for i := 0; i < length; i++ {
		t := rand.Intn(62)
		if t < 10 {
			result = append(result, strconv.Itoa(rand.Intn(10)))
		} else if t < 36 {
			result = append(result, string(rand.Intn(26)+65))
		} else {
			result = append(result, string(rand.Intn(26)+97))
		}
	}

	return strings.Join(result, "")
}
