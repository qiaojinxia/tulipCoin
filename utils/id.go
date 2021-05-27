package utils

import "sync/atomic"

/**
 * Created by @CaomaoBoy on 2021/5/25.
 *  email:<115882934@qq.com>
 */

var ID int64 = 1
func GetUserID() int64{
	return atomic.AddInt64(&ID,1)
}