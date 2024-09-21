package commons

import (
	"crypto/aes"
	"encoding/base64"
	"fmt"
)

// PKCS5UnPadding implements PKCS5 unpadding.
func PKCS5UnPadding(src []byte) ([]byte, error) {
	length := len(src)
	unpadding := int(src[length-1])
	if unpadding > length {
		return nil, fmt.Errorf("invalid padding")
	}
	return src[:length-unpadding], nil
}

func AesDecrypt(ciphertext, key string) ([]byte, error) {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return nil, err
	}

	// Base64解码密文
	decodedCiphertext, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return nil, err
	}

	// 创建一个与密文长度相同的新缓冲区用于存储解密后的数据
	decrypted := make([]byte, len(decodedCiphertext))

	// 对每一个块进行解密
	for bs, be := 0, aes.BlockSize; bs < len(decodedCiphertext); bs, be = bs+aes.BlockSize, be+aes.BlockSize {
		block.Decrypt(decrypted[bs:be], decodedCiphertext[bs:be])
	}

	// 移除PKCS5Padding
	decrypted, err = PKCS5UnPadding(decrypted)
	if err != nil {
		return nil, err
	}

	return decrypted, nil
}
