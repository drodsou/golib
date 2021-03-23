package aestool

import (
	"testing"
)

func TestCipher(t *testing.T) {
	text := "test text"

	key1 := "123456789012345678901234567890AB"
	if Decipher(Cipher(text, key1), key1) != text {
		t.Errorf("aestool decipher 1 failed")
	}

	key2 := RandString32()
	if Decipher(Cipher(text, key2), key2) != text {
		t.Errorf("aestool decipher 2 failed")
	}

	key3 := RandString32("ABCD")
	if Decipher(Cipher(text, key3), key3) != text {
		t.Errorf("aestool decipher 3 failed")
	}

	t.Log(Cipher(text, key2))
}

func TestRandString(t *testing.T) {
	str1 := RandString32()
	t.Log(str1)
	if len(str1) != 32 {
		t.Error("RandString32 with no param failed")
	}

	str2 := RandString32("ABCDE")
	t.Log(str2)
	if len(str2) != 32 {
		t.Error("RandString32 with param failed")
	}
}
