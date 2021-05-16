package vm

import "main/utils"

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
)

type OperationCode struct {
	Operation //opcode
	args [][]byte //Operating parameters
}

func(s *Stack) Dump(){
	//data := s.GetTop()
	//s.Push(data)
}

//Program runtime stack
type OperationStack struct {
	LocalVariableTable [][]byte         // Program Local variable Data
	Opcode             []*OperationCode //Program OpArg to Analysis
	stack              *Stack           //Program Run stack
	pc                 int              //now OperationCode index
}

func pushBytes(stack *Stack,args [][]byte,localVariableTable [][]byte){
	//for _,arg := range args{
	//	dataIndex := utils.BytesToInt(arg)
	//	if dataIndex > len(localVariableTable){
	//		panic("index overflow")
	//	}
	//	stack.Push(localVariableTable[dataIndex])
	//}
}

func storeBytes(stack *Stack,args [][]byte,localVariableTable [][]byte) {
	dataIndex := utils.BytesToInt64(args[0])
	data := stack.PopBytes()
	if int(dataIndex) > len(localVariableTable) {
		panic("index overflow")
	}
	localVariableTable[dataIndex] = data
}



func(os *OperationStack) Run(){
	for{
		opCode := os.Opcode[os.pc]
		switch opCode.Operation {
		case PushBytes:
			pushBytes(os.stack,opCode.args,os.LocalVariableTable)
		case StoreBytes_t0:
			storeBytes(os.stack,opCode.args,os.LocalVariableTable)
		case CMP_EQ_I64:
			cmpEqualInt64(os.stack)
		}
		os.pc ++
	}

}