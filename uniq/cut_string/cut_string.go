package cut_string

import (
	"strings"
	"github.com/mike5535/uniq/uniq_types"
)

//убирает символы из строки
func CutSymbols(inputStr string, numChars int) string {
	if inputStr == "" {
		return inputStr
	}
	buf := strings.Split(inputStr, "")
	if buf[0] == " " {
		copy(buf[:], buf[1:])
		buf[len(buf)-1] = ""
		buf = buf[:len(buf)-1]
	}
	if numChars >= cap(buf) {
		return ""
	}
	buf = buf[numChars:]
	return strings.Join(buf, "")
}

//убирает поля из строки
func CutString(inputStr string, param uniq_types.OptFields) string {
	if inputStr == "" {
		return ""
	}

	if param.AnyRegister {
		inputStr = strings.ToLower(inputStr)
	}

	if param.NumFields != 0 {
		s := strings.Split(inputStr, " ")[param.NumFields:]
		temp := strings.Join(s, " ")
		if param.NumChars != 0 {
			temp = CutSymbols(temp, param.NumChars)
		}
		return temp
	}

	if param.NumChars != 0 {
		return CutSymbols(inputStr, param.NumChars)
	}

	return inputStr
}
