package test

import (
	"fmt"
	"qiyuesuo/sdk/commons"
	"testing"
)

/*
*
aes 解密
*/
func TestAes(t *testing.T) {
	ciphertext := "UoYYii0CN0bJyDcdkgf/iA=="
	key := "5f6db7ec8325aAGF"

	decrypted, err := commons.AesDecrypt(ciphertext, key)
	if err != nil {
		fmt.Println("Error decrypting:", err)
		return
	}
	fmt.Println("Decrypted text:", string(decrypted))

}
