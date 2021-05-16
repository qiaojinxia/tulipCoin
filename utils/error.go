package utils

import "fmt"

/**
 * Created by @CaomaoBoy on 2021/5/16.
 *  email:<115882934@qq.com>
 */

func StackError(msg string){
	panic(fmt.Sprintf("StackConvertError:%s",msg))
}