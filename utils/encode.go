package utils

import (
	"bytes"
	"encoding/binary"
)

/**
 * Created by @CaomaoBoy on 2021/5/16.
 *  email:<115882934@qq.com>
 */


//整形转换成字节
func ToBytes(n interface{}) []byte {
	x := n
	bytesBuffer := bytes.NewBuffer([]byte{})
	err := binary.Write(bytesBuffer, binary.BigEndian, x)
	if err != nil{
		panic(ConverErrorWarp(err.Error()))
	}
	return bytesBuffer.Bytes()
}

//整形转换成字节
func Int8ToBytes(n int8) []byte {
	x := n
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, x)
	return bytesBuffer.Bytes()
}

//字节转换成整形
func BytesToInt64(b []byte) int64 {
	bytesBuffer := bytes.NewBuffer(b)

	var x int64
	binary.Read(bytesBuffer, binary.BigEndian, &x)

	return x
}
func BytesToInt16(b []byte) int16 {
	bytesBuffer := bytes.NewBuffer(b)

	var x int16
	binary.Read(bytesBuffer, binary.BigEndian, &x)

	return x
}

func BytesToInt32(b []byte) int32 {
	bytesBuffer := bytes.NewBuffer(b)

	var x int32
	binary.Read(bytesBuffer, binary.BigEndian, &x)

	return x
}



func BytesToInt8(b []byte) int8 {
	bytesBuffer := bytes.NewBuffer(b)

	var x int8
	binary.Read(bytesBuffer, binary.BigEndian, &x)

	return x
}


func BytesToFloat32(b []byte) float32 {
	bytesBuffer := bytes.NewBuffer(b)

	var x float32
	binary.Read(bytesBuffer, binary.BigEndian, &x)

	return x
}
func BytesToFloat64(b []byte) float64 {
	bytesBuffer := bytes.NewBuffer(b)

	var x float64
	binary.Read(bytesBuffer, binary.BigEndian, &x)

	return x
}

func BytesToUInt8(b []byte) uint8 {
	bytesBuffer := bytes.NewBuffer(b)
	var x uint8
	binary.Read(bytesBuffer, binary.BigEndian, &x)
	return x
}


func BytesToUInt16(b []byte) uint16 {
	bytesBuffer := bytes.NewBuffer(b)
	var x uint16
	binary.Read(bytesBuffer, binary.BigEndian, &x)
	return x
}


func BytesToUInt32(b []byte) uint32 {
	bytesBuffer := bytes.NewBuffer(b)
	var x uint32
	binary.Read(bytesBuffer, binary.BigEndian, &x)
	return x
}

func BytesToUInt64(b []byte) uint64 {
	bytesBuffer := bytes.NewBuffer(b)
	var x uint64
	binary.Read(bytesBuffer, binary.BigEndian, &x)
	return x
}
