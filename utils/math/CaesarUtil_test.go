package math

import (
	"fmt"
	"testing"
)

func TestCaesarUtil(t *testing.T) {
	// Test the CaesarUtil.Encrypt function
	if CaesarEncrypt("Hello World!", 3) != "Khoor Zruog!" {
		t.Error("CaesarUtil.Encrypt failed")
	}
	// Test the CaesarUtil.Decrypt function
	if CaesarDecrypt("Khoor Zruog!", 3) != "Hello World!" {
		t.Error("CaesarUtil.Decrypt failed")
	}

	fmt.Println(CaesarEncrypt("Hello World!你好，世界", 3))
}
