package base65536

import(
    "testing"
)

func TestEncode(t *testing.T) {
    s := "hello world"
    expected := "驨ꍬ啯𒁷ꍲᕤ"

    actual := Encode(s)
    if (actual != expected) {
        t.Errorf("This encoder is not correct.\nexpected: %s\nactual: %s", expected, actual)
    }
}

func TestDecode(t *testing.T) {
    s := "驨ꍬ啯𒁷ꍲᕤ"
    expected := "hello world"

    actual := Decode(s)
    if (actual != expected) {
        t.Errorf("This encoder is not correct.\nexpected: %s\nactual: %s", expected, actual)
    }
}

func TestCorrectnessChecking1(t *testing.T) {
    str := "Why Japanese!? おかしいやろぉ！！ 躊躇って書いてる間に躊躇してまぅわ！"
    encode_str := Encode(str)
    decode_str := Decode(encode_str)

    if (str != decode_str) {
        t.Errorf("This encode/decode is not correct.\nbefore: %s\nafter: %s", str, decode_str)
    }
}