package animateDL

import (
	"crypto/aes"
	"crypto/cipher"
)

func AES128Decrypt(crypted, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	mode := cipher.NewCBCDecrypter(block, key[:aes.BlockSize])
	returnData := make([]byte, len(crypted))
	mode.CryptBlocks(returnData, crypted)
	returnData = pkc5Unpadding(returnData)
	return returnData, nil
}
func pkc5Unpadding(returnData []byte) []byte {
	dataLen := len(returnData)
	unPadding := int(returnData[dataLen-1])
	return returnData[:(dataLen - unPadding)]
}
