package svm

/**
 * Created by @CaomaoBoy on 2021/5/16.
 *  email:<115882934@qq.com>
 */

//Get top 2 data exec Cmp operation

func cmpEqualInt64(stack *Stack) {
	data1 := stack.PopInt64()
	data2 := stack.PopInt64()
	stack.PushBool(data1 == data2)
}

func cmpEqualInt32(stack *Stack){
	data1 := stack.PopInt32()
	data2 := stack.PopInt32()
	stack.PushBool(data1 == data2)
}

func cmpEqualInt16(stack *Stack){
	data1 := stack.PopInt16()
	data2 := stack.PopInt16()
	stack.PushBool(data1 == data2)
}


func cmpEqualUint16(stack *Stack){
	data1 := stack.PopInt64()
	data2 := stack.PopInt64()
	stack.PushBool(data1 == data2)
}