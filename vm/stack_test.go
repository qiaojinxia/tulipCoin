package vm

import (
	"fmt"
	"testing"
)

/**
 * Created by @CaomaoBoy on 2021/5/16.
 *  email:<115882934@qq.com>
 */

func Test_PushBytes(t *testing.T) {
	stack := NewStack()
	stack.PushBytes([]byte("caomaoboy!"))
	fmt.Println(string(stack.PopBytes()))
}

func BenchmarkStack_PopBytes(b *testing.B) {
	stack := NewStack()
	for i:=0;i<b.N;i++{
		stack.PushBytes([]byte("caomaoboy!"))
		stack.PopBytes()
	}
}

func BenchmarkStack_PushInt8(b *testing.B) {
	stack := NewStack()
	for i:=0;i<b.N;i++{
		stack.PushInt8(4)
		stack.PopInt8()
	}
}

func BenchmarkStack_PopInt64(b *testing.B) {
	stack := NewStack()
	for i:=0;i<b.N;i++{
		stack.PushInt64(99999999)
		fmt.Println(stack.PopInt64())
	}
}