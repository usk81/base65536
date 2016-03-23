package base65536

import (
	// test library
	"testing"
	// other libraries
	"strings"
)

func TestEncode(t *testing.T) {
	s := "hello world"
	expected := "驨ꍬ啯𒁷ꍲᕤ"

	actual := Encode(s)
	if actual != expected {
		t.Errorf("This encoder is not correct.\nexpected: %s\nactual: %s", expected, actual)
	}
}

func TestDecode(t *testing.T) {
	s := "驨ꍬ啯𒁷ꍲᕤ"
	expected := "hello world"

	actual := Decode(s)
	if actual != expected {
		t.Errorf("This encoder is not correct.\nexpected: %s\nactual: %s", expected, actual)
	}
}

func TestCorrectnessCheckingJapaneseAndEnglish(t *testing.T) {
	str := "Why Japanese!? おかしいやろぉ！！ 躊躇って書いてる間に躊躇してまぅわ！"
	encodeStr := Encode(str)
	decodeStr := Decode(encodeStr)

	if str != decodeStr {
		t.Errorf("This encode/decode is not correct.\nbefore: %s\nafter: %s", str, decodeStr)
	}
}

func TestCorrectnessCheckingAllString(t *testing.T) {
	strMap := []string{
		"1234567890",
		"abcdefghijklmnopqrstuvwxyz",
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ",
		"あいうえおかきくけこさしすせそたちつてとなにぬねのはひふへほまみむめもやゐゆゑよらりるれろわをん",
		"アイウエオカキクケコサシスセソタチツテトナニヌネノハヒフヘホマミムメモヤヰユヱヨラリルレロワヲン",
	}
	str := strings.Join(strMap, "")
	encodeStr := Encode(str)
	decodeStr := Decode(encodeStr)

	if str != decodeStr {
		t.Errorf("This encode/decode is not correct.\nbefore: %s\nafter: %s", str, decodeStr)
	}
}
