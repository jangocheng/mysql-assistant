package foundations

import (
	"crypto/sha1"
	"fmt"
)

func CheckPasswd(hashStr string, passwd string, salt interface{}) bool {
	return hashStr == EncryptWord(passwd, salt)
}

//EncryptWord
func EncryptWord(passwd string, salt interface{}) (hashResult string) {

	h := sha1.New()

	h.Write([]byte(passwd))
	if nil != salt {
		h.Write([]byte(salt.(string)))
	}

	// 16进制输出的结果才和php是一样的。  php默认按16进制进行输出。
	//@see https://segmentfault.com/q/1010000007510284
	hashResult = fmt.Sprintf("%x", h.Sum(nil))
	return
}
