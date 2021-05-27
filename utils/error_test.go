package utils

import (
	"fmt"
	"testing"
)

/**
 * Created by @CaomaoBoy on 2021/5/21.
 *  email:<115882934@qq.com>
 */

func Test_error(t *testing.T) {
	fmt.Println(ConverToJsonInfo(StackErrorWarp("错误了!")))
}