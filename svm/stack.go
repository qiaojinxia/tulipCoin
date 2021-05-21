package svm

import (
	"main/utils"
)

/**
 * Created by @CaomaoBoy on 2021/5/1.
 *  email:<115882934@qq.com>
 */

var DefeultStackSize = 1024

var MaxBytesInstack = 2

type Stack struct {
	index int
	codes []byte
}

func NewStack() *Stack {
	return &Stack{codes: make(OpData,0,DefeultStackSize)}
}

func(s *Stack) PushBytes(data []byte){
	//data length suport index 65535 size
	lengeth:= int16(len(data))
	if lengeth > int16(1 << 8 * MaxBytesInstack){
		utils.StackError("Max Bytes Size!")
	}
	data = append(data, utils.ToBytes(lengeth)...)
	s.index += len(data)
	s.codes = append(s.codes, data...)
}

func(s *Stack) PushInt8(data int8) {
	s.index += 1
	s.codes = append(s.codes, utils.ToBytes(data)...)
}

func(s *Stack) PushInt16(data int16) {
	s.index += 2
	s.codes = append(s.codes, utils.ToBytes(data)...)
}

func(s *Stack) PushInt32(data int32){
	s.index += 4
	s.codes = append(s.codes, utils.ToBytes(data)...)
}

func(s *Stack) PushInt64(data int64){
	s.index += 8
	s.codes = append(s.codes, utils.ToBytes(data)...)
}


func(s *Stack) PushUint8(data uint8) {
	s.index += 1
	s.codes = append(s.codes, utils.ToBytes(data)...)
}

func(s *Stack) PushUint16(data uint16) {
	s.index += 2
	s.codes = append(s.codes, utils.ToBytes(data)...)
}

func(s *Stack) PushUint32(data uint32) {
	s.index += 4
	s.codes = append(s.codes, utils.ToBytes(data)...)
}

func(s *Stack) PushUint64(data uint64) {
	s.index += 8
	s.codes = append(s.codes, utils.ToBytes(data)...)
}

func(s *Stack) PushBool(data bool){
	s.index += 1
	s.codes = append(s.codes, utils.ToBytes(data)...)
}

func(s *Stack) PushFloat32(data float32){
	s.index += 4
	s.codes = append(s.codes, utils.ToBytes(data)...)
}

func(s *Stack) PushFloat64(data float64){
	s.index += 8
	s.codes = append(s.codes, utils.ToBytes(data)...)
}

func(s *Stack) PopBytes() OpData{
	lengeth := int(utils.BytesToInt16(s.codes[len(s.codes)-MaxBytesInstack:]))
	tmp := s.codes[len(s.codes)- (lengeth + MaxBytesInstack ):len(s.codes) - MaxBytesInstack ]
	s.codes= s.codes[:len(s.codes)-(lengeth + MaxBytesInstack)]
	return tmp
}
func(s *Stack) PopInt8() int8{
	s.index -= 1
	tmp := s.codes[s.index:]
	s.codes = s.codes[:s.index]
	return utils.BytesToInt8(tmp)
}

func(s *Stack) PopInt16() int16{
	s.index -= 2
	tmp := s.codes[s.index:]
	s.codes = s.codes[:s.index]
	return utils.BytesToInt16(tmp)
}

func(s *Stack) PopInt32() int32{
	s.index -= 4
	tmp := s.codes[s.index:]
	s.codes = s.codes[:s.index]
	return utils.BytesToInt32(tmp)
}

func(s *Stack) PopInt64() int64{
	s.index -= 8
	tmp := s.codes[s.index:]
	s.codes = s.codes[:s.index]
	return utils.BytesToInt64(tmp)
}

func(s *Stack) PopUint8() uint8{
	s.index -= 1
	tmp := s.codes[s.index:]
	s.codes = s.codes[:s.index]
	return utils.BytesToUInt8(tmp)
}

func(s *Stack) PopUint16() uint16{
	s.index -= 2
	tmp := s.codes[s.index:]
	s.codes = s.codes[:s.index]
	return utils.BytesToUInt16(tmp)
}

func(s *Stack) PopUInt32() uint32{
	s.index -= 4
	tmp := s.codes[s.index:]
	s.codes = s.codes[:s.index]
	return utils.BytesToUInt32(tmp)
}

func(s *Stack) PopUInt64() uint64{
	s.index -= 8
	tmp := s.codes[s.index:]
	s.codes = s.codes[:s.index]
	return utils.BytesToUInt64(tmp)
}


func(s *Stack) PopFloat32() float32{
	s.index -= 4
	tmp := s.codes[s.index:]
	s.codes = s.codes[:s.index]
	return utils.BytesToFloat32(tmp)
}

func(s *Stack) PopFloat64() float64{
	s.index -= 4
	tmp := s.codes[s.index:]
	s.codes = s.codes[:s.index]
	return utils.BytesToFloat64(tmp)
}
