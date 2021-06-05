package utils

import (
	"fmt"
	"github.com/pkg/errors"
	"testing"
)

/**
 * Created by @CaomaoBoy on 2021/5/21.
 *  email:<115882934@qq.com>
 */

func Test_error(t *testing.T) {
	fmt.Println(ConverToJsonInfo(StackErrorWarp(errors.New("123"),"错误了!")))
}