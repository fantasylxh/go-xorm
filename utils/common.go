package utils

import (
	"crypto/md5"
	"fmt"
	"reflect"
)

func EncryptPassword(password string,passwordSalt string) string  {
	pwd := []byte(password)
	salt := []byte(passwordSalt)
	pwdStr := fmt.Sprintf("%x", md5.Sum(pwd))
	salStr := fmt.Sprintf("%x",md5.Sum(salt))

	pwdSaltByte := []byte(pwdStr+ salStr)
	md5pwd := fmt.Sprintf("%x",md5.Sum(pwdSaltByte))
	return md5pwd
}
//去除数组重复
func Duplicate(a interface{}) (ret []interface{}) {
	va := reflect.ValueOf(a)
	for i := 0; i < va.Len(); i++ {
		if i > 0 && reflect.DeepEqual(va.Index(i-1).Interface(), va.Index(i).Interface()) {
			continue
		}
		ret = append(ret, va.Index(i).Interface())
	}
	return ret
}


