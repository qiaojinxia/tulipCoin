package svm

import (
	"bytes"
	"encoding/hex"
	"errors"
	"fmt"
	"main/utils"
	"strings"
)

/**
 * Created by @CaomaoBoy on 2021/5/1.
 *  email:<115882934@qq.com>
 */
type MStack interface {
	Run()
}

type Operation uint8

const(
	PushBytes Operation = 1
	StoreBytes_t0 Operation = 2 //Bytes Pop From Stack Top 0 Into Localvariable Table
	StoreBytes_t1 Operation = 3 //Bytes Pop From Stack Top 1 Into Localvariable Table
	PoPInt Operation = 4
	DUMP Operation = 5
	CMP_EQ_I64  Operation = 6
	OP_HASH160 Operation = 7
	OP_EQUALVERIFY Operation = 8
	OP_CHECKSIG Operation = 9
	OP_DUP = 10
)

type OperationCode struct {
	Operation //opcode
	args [][]byte //Operating parameters
}

func(s *Stack) Dump(){
	data := s.GetTop()
	s.PushBytes(data)
}

//Program runtime stack
type OperationStack struct {
	LocalVariableTable [][]byte         // Program Local variable Data
	Opcode             []*OperationCode //Program OpArg to Analysis
	stack              *Stack           //Program Run stack
	pc                 int              //now OperationCode index
}

func NewOperationStack(script string,txHash []byte) *OperationStack{
	codes := strings.Split(script," ")
	lovalVariable := make([][]byte,0,len(codes))
	opCode := make([]*OperationCode,0,len(codes))
	for _,code := range codes{
		switch code {
		case "OP_HASH160":
			opcode := &OperationCode{Operation: OP_HASH160}
			opCode = append(opCode,opcode)
		case "OP_EQUALVERIFY":
			opcode := &OperationCode{
				Operation: OP_EQUALVERIFY}
			opCode = append(opCode,opcode)
		case "OP_CHECKSIG":
			opcode := &OperationCode{
				Operation: OP_CHECKSIG,
				args: [][]byte{
					txHash,
				}}
			opCode = append(opCode,opcode)
		case "OP_DUP":
			opcode := &OperationCode{
				Operation: OP_DUP}
			opCode = append(opCode,opcode)
		case "":
			continue
		default:
			lovalVariable = append(lovalVariable, []byte(code))
			index := len(lovalVariable)
			opcode := &OperationCode{
				Operation: PushBytes,
				args:      [][]byte{
					utils.ToBytes(int32(index-1)),
				},
			}
			opCode = append(opCode,opcode)
		}
	}
	return &OperationStack{
		LocalVariableTable: lovalVariable,
		Opcode:             opCode,
		stack:              NewStack(),
		pc:                 0,
	}
}

func opHash160(stack *Stack,args [][]byte,localVariableTable [][]byte){
	d1,err := hex.DecodeString(string(stack.PopBytes()))
	if err != nil{
		panic(utils.ConverErrorWarp(err.Error()))
	}
	d2 := utils.GeneratePublicKeyHash(d1)
	d3 := fmt.Sprintf("%x",d2)
	stack.PushBytes([]byte(d3))
}

func opEqualverify(stack *Stack,args [][]byte,localVariableTable [][]byte){
	d1 := stack.PopBytes()
	d2 := stack.PopBytes()
	res := bytes.Equal(d1,d2)
	if !res {
		utils.StackErrorWarp("opEqualverify Error!")
	}
}

func opChecksig(stack *Stack,args [][]byte,localVariableTable [][]byte){
	dataHash,err := hex.DecodeString(string(args[0]))
	if err != nil{
		panic(utils.ConverErrorWarp(err.Error()))
	}
	pubKey := stack.PopBytes()
	signature := stack.PopBytes()
	bPubKey := make([]byte,len(pubKey))
	bSignature := make([]byte,len(signature))
	n,err := hex.Decode(bPubKey,pubKey)
	bPubKey = bPubKey[:n]
	if err != nil{
		panic(utils.ConverErrorWarp(err.Error()))
	}
	n,err = hex.Decode(bSignature,signature)
	bSignature = bSignature[:n]
	if err != nil{
		panic(utils.ConverErrorWarp(err.Error()))
	}
	if utils.Verify(bSignature,bPubKey,dataHash){
		stack.PushBool(true)
		return
	}
	stack.PushBool(false)
}

func opDUP(stack *Stack,args [][]byte,localVariableTable [][]byte){
	stack.Dump()
}

func pushBytes(stack *Stack,args [][]byte,localVariableTable [][]byte){
	for _,arg := range args{
		dataIndex := utils.BytesToInt32(arg)
		if int(dataIndex) > len(localVariableTable){
			utils.StackErrorWarp("index overflow")
		}
		stack.PushBytes(localVariableTable[dataIndex])
	}
}

func storeBytes(stack *Stack,args [][]byte,localVariableTable [][]byte) {
	dataIndex := utils.BytesToInt32(args[0])
	data := stack.PopBytes()
	if int(dataIndex) > len(localVariableTable) {
		panic("index overflow")
	}
	localVariableTable[dataIndex] = data
}



func(os *OperationStack) Run() error {
	for os.pc < len(os.Opcode){
		opCode := os.Opcode[os.pc]
		switch opCode.Operation {
		case PushBytes:
			pushBytes(os.stack,opCode.args,os.LocalVariableTable)
		case StoreBytes_t0:
			storeBytes(os.stack,opCode.args,os.LocalVariableTable)
		case CMP_EQ_I64:
			cmpEqualInt64(os.stack)
		case OP_DUP:
			opDUP(os.stack,opCode.args,os.LocalVariableTable)
		case OP_HASH160:
			opHash160(os.stack,opCode.args,os.LocalVariableTable)
		case OP_EQUALVERIFY:
			opEqualverify(os.stack,opCode.args,os.LocalVariableTable)
		case OP_CHECKSIG:
			opChecksig(os.stack,opCode.args,os.LocalVariableTable)
		}
		os.pc ++
	}
	if !os.stack.PopBool(){
		return errors.New("Valid Script Failed!")
	}
	return nil

}