package net

import "testing"

/**
 * Created by @CaomaoBoy on 2021/5/21.
 *  email:<115882934@qq.com>
 */

func TestTcpServer_Listen(t *testing.T) {
	server := &TcpServer{
		Address: "127.0.0.1",
		Port:    "7777",
	}
	server.Listen()
}