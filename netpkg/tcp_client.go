package netpkg

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

/**
 * Created by @CaomaoBoy on 2021/5/21.
 *  email:<115882934@qq.com>
 */


func Start() {
	conn, err := net.Dial("tcp4", "127.0.0.1:7777")
	defer conn.Close()
	if err != nil {
		fmt.Printf("connect failed, err : %v\n", err.Error())
		return
	}
	inputReader := bufio.NewReader(os.Stdin)
	for {
		input, err := inputReader.ReadString('\n')
		if err != nil {
			fmt.Printf("read from console failed, err: %v\n", err)
			break
		}
		trimmedInput := strings.TrimSpace(input)
		if trimmedInput == "Q" {
			break
		}
		_, err = conn.Write(Pack([]byte(trimmedInput),0))

		if err != nil {
			fmt.Printf("write failed , err : %v\n", err)
			break
		}
	}
}