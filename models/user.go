package models

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"strconv"
	"time"

	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Account  string
	Password string
	Salt     string
	Status   int
}

func (this *User) Encrypt() {
	this.Salt = createSalt()
	this.Password = createPassword(this.Password, this.Salt)
}

func (this *User) ClearPassword() {
	this.Password = ""
	this.Salt = ""
}

// EncryptWithSalt 加密密码(使用原有的salt)
func (this *User) EncryptWithSalt() {
	this.Password = createPassword(this.Password, this.Salt)
}

// CreateSalt 创建Salt
func createSalt() string {
	rand.Seed(time.Now().UnixNano())
	return strconv.Itoa(rand.Int())
}

// CreatePassword 创建密码
//  originalPassword:原始密码
func createPassword(originalPassword string, salt string) string {
	var m = md5.New()
	m.Write([]byte(originalPassword + salt))
	return hex.EncodeToString(m.Sum(nil))
}
