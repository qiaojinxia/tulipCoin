package netpkg

import (
	"testing"
	"time"
)

/**
 * Created by @CaomaoBoy on 2021/5/21.
 *  email:<115882934@qq.com>
 */

func TestTcpServer_Listen(t *testing.T) {
	server := &TcpServer{
		Address: "127.0.0.1",
		Port:    "7777",
		stopChan: make(chan struct{},1),
	}
	go func() {
		time.Sleep(time.Second * 7)
		server.Close()
	}()
	server.Listen()

}