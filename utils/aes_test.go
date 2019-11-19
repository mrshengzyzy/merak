package utils

import (
	"fmt"
	"testing"
)

func TestAesDecryptWithBase64(t *testing.T) {

	salt := "B1827B657FFF9232"
	data := "A3DD5Y6JJA"

	fmt.Println("Origin:", data)

	en, err := AesEncryptWithBase64(data, salt)
	fmt.Println("Encrypt:", en)

	if err != nil {
		t.Error(err.Error())
	}

	unEn, err := AesDecryptWithBase64(en, salt)
	fmt.Println("UnEncrypt:", unEn)

	if err != nil {
		t.Error(err.Error())
	}

	if data != unEn {
		t.Error("No Pass")
	}
}
