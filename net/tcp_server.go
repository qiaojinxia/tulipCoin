package net

import (
	"fmt"
	"log"
	"main/utils"
	"net"
	"os"
	"os/signal"
	"syscall"
)

/**
 * Created by @CaomaoBoy on 2021/5/21.
 *  email:<115882934@qq.com>
 */


type Server interface {
	Listen()
}
type TcpServer struct {
	Address string
	Port string
}

func(ts *TcpServer) Listen(){
	utils.Try(func() {
		var (
			done = make(chan bool, 1)
			quit = make(chan os.Signal, 1)
		)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		tcpAddr, err := net.ResolveTCPAddr("tcp4", fmt.Sprintf("%s:%s",ts.Address,ts.Port))
		if err != nil{
			panic(utils.NetError(err))
		}
		log.Println(fmt.Sprintf("Listening %s:%s....",ts.Address,ts.Port))
		listener, err := net.ListenTCP("tcp", tcpAddr)
		defer func() {
			log.Println(fmt.Sprintf("Close Server %s:%s....",ts.Address,ts.Port))
			defer listener.Close()
			log.Println(fmt.Sprintf("Server Closed %s:%s....",ts.Address,ts.Port))
		}()

		if err != nil{
			panic(utils.NetError(err))
		}
		for{
			conn,err := listener.Accept()
			if err != nil{
				panic(utils.NetError(err))
			}
			go func() {
				Handle(conn)
			}()
		}
		<- done
	}).Catch(utils.NetErroWarp(""), func(err error) {
		log.Panic(utils.NetErroWarp(err.Error()).ConverToJsonInfo())
	})
}

func Handle(conn net.Conn){
	buff := make([]byte,1024)
	conn.Read(buff)
	fmt.Println(buff)
}