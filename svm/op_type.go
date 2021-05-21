package svm

import (
	"bytes"
	"main/utils"
	"strconv"
	"unsafe"
)


/**
 * Created by @CaomaoBoy on 2021/5/16.
 *  email:<115882934@qq.com>
 */
type OpData []byte

func(od OpData) Equals(toCompareData []byte) bool{
	return bytes.Equal(od,toCompareData)
}
func(od OpData) ConvertToInt()  int{
	 return int(utils.BytesToInt64(od))
}
func(od OpData) ConvertToBool() bool{
	val := *(*string)(unsafe.Pointer(&od))
	bo,err :=  strconv.ParseBool(val)
	if err != nil{
		utils.StackError("stackError")
	}
	return bo
}

func(od OpData) ConvertTostring() string{
	return string(od)
}
