package mask

import (
	"fmt"
	"testing"
)

func TestPrefixPreserveMasker(t *testing.T) {
	masker := NewPrefixPreserveMasker()
	fmt.Println(masker.MaskInteger(100100002,2))
	fmt.Println(masker.MaskString("123456789",6))
	fmt.Println(masker.MaskString("123456789大量ldfjsldf",6))
}

