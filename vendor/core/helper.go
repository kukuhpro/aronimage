package core

import (
	"encoding/base64"
	"reflect"
	"regexp"
	"strings"
	"time"
)

type replacement struct {
	re *regexp.Regexp
	ch string
}

var (
	rExps = []replacement{
		{re: regexp.MustCompile(`[\xC0-\xC6]`), ch: "A"},
		{re: regexp.MustCompile(`[\xE0-\xE6]`), ch: "a"},
		{re: regexp.MustCompile(`[\xC8-\xCB]`), ch: "E"},
		{re: regexp.MustCompile(`[\xE8-\xEB]`), ch: "e"},
		{re: regexp.MustCompile(`[\xCC-\xCF]`), ch: "I"},
		{re: regexp.MustCompile(`[\xEC-\xEF]`), ch: "i"},
		{re: regexp.MustCompile(`[\xD2-\xD6]`), ch: "O"},
		{re: regexp.MustCompile(`[\xF2-\xF6]`), ch: "o"},
		{re: regexp.MustCompile(`[\xD9-\xDC]`), ch: "U"},
		{re: regexp.MustCompile(`[\xF9-\xFC]`), ch: "u"},
		{re: regexp.MustCompile(`[\xC7-\xE7]`), ch: "c"},
		{re: regexp.MustCompile(`[\xD1]`), ch: "N"},
		{re: regexp.MustCompile(`[\xF1]`), ch: "n"},
	}
	spacereg       = regexp.MustCompile(`\s+`)
	noncharreg     = regexp.MustCompile(`[^A-Za-z0-9-]`)
	minusrepeatreg = regexp.MustCompile(`\-{2,}`)
)

func DateTimeNow() string {
	t := time.Now()
	return t.Format("2006-01-02 03:04:05")
}

func Base64ToByte(base64EncodeString string) []byte {
	e64 := base64.StdEncoding
	enc := []byte(base64EncodeString)
	maxDecLen := e64.DecodedLen(len(enc))
	decBuf := make([]byte, maxDecLen)

	n, err := e64.Decode(decBuf, enc)
	_ = err

	return decBuf[0:n]
}

func ToSlug(s string, lower ...bool) string {
	for _, r := range rExps {
		s = r.re.ReplaceAllString(s, r.ch)
	}

	s = strings.ToLower(s)
	s = spacereg.ReplaceAllString(s, "-")
	s = noncharreg.ReplaceAllString(s, "")
	s = minusrepeatreg.ReplaceAllString(s, "-")

	return s
}

func In_Array(val interface{}, array interface{}) (exists bool, index int) {
	exists = false
	index = -1

	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(array)

		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(val, s.Index(i).Interface()) == true {
				index = i
				exists = true
				return
			}
		}
	}

	return
}

type Response struct {
	HttpStatus     int
	StatusResponse uint
	Data           map[string]interface{}
}

type Exception struct {
	Response                 Response
	MappingExceptionFunction map[string]interface{}
}

func (this *Exception) ErrorException(exceptionname string) Response {
	reflectFunction := reflect.ValueOf(this.MappingExceptionFunction[exceptionname])
	params := []reflect.Value{}
	t := reflectFunction.Call(params)
	resData := t[0].Interface().(Response)

	return resData
}
