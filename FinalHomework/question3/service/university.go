package service

import (
	"crypto/md5"
	"encoding/hex"
)

func (dc *DataCryptographyMD5)Cryptography(){
	has:=md5.New()
	has.Write([]byte(dc.Data))
	tem:= has.Sum([]byte(""))
	dc.Result=hex.EncodeToString(tem)
}