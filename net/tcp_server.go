package net

import (
	"fmt"
	"io"
	"log"
	"main/utils"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

/**
 * Created by @CaomaoBoy on 2021/5/21.
 *  email:<115882934@qq.com>
 */


type Server interface {
	Listen()
	Close()
}
type TcpServer struct {
	Address string
	Port string
	stopChan chan struct{}
}

func(ts *TcpServer) Close(){
	ts.stopChan <- struct{}{}
}

func(ts *TcpServer) Listen(){
	utils.Try(func() {
		var (
			quit = make(chan os.Signal, 1)
		)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		utils.GO_Func(func() {
			for  {
				select {
				case <-quit:
					ts.Close()
					return
				case <- ts.stopChan:
					return
				}
			}
		})
		utils.GO_Func(func() {
			tcpAddr, err := net.ResolveTCPAddr("tcp4", fmt.Sprintf("%s:%s",ts.Address,ts.Port))
			if err != nil{
				panic(utils.NetErroWarp(err.Error()))
			}
			log.Println(fmt.Sprintf("Listening %s:%s....",ts.Address,ts.Port))
			listener, err := net.ListenTCP("tcp", tcpAddr)
			defer func() {
				log.Println(fmt.Sprintf("Close Server %s:%s....",ts.Address,ts.Port))
				ts.Close()
				listener.Close()
				log.Println(fmt.Sprintf("Server Closed %s:%s....",ts.Address,ts.Port))
			}()
			if err != nil{
				panic(utils.NetErroWarp(err.Error()))
			}
			for{
				conn,err := listener.Accept()
				if err != nil{
					panic(utils.NetErroWarp(err.Error()))
				}
				utils.GO_Func(func() {
					Handle(conn,ts.stopChan)
				})
			}
		})
		time.Sleep(time.Second * 2)
		utils.Wg.Wait()
	}).Catch(utils.NetErroWarp(""), func(err error) {
		log.Panic(utils.ConverToJsonInfo(utils.NetErroWarp(err.Error())))
	}).CatchAll(func(err error) {
		log.Println(err)
	})
}

func Handle(conn net.Conn,stopCh <-chan struct{}){
	se := &Session{conn: conn,ID: time.Now().UnixNano()}
	_SessionManger.AddSession(se)
	log.Printf("Client %d Online!",se.ID)
	for{
		select {
		case <- stopCh:
			conn.Write([]byte("Server Will Be Closed!"))
			conn.Close()
			log.Printf("Conn From %s Close!\n",conn.RemoteAddr().String())
			return
		default:
			buff := make([]byte,1024)
			rlen,err := conn.Read(buff)
			if err != nil{
				_SessionManger.RemoveSession(se.ID)
				if err == io.EOF{
					log.Printf("Client %d Offline!",se.ID)
					return
				}
				utils.NetErroWarp(err.Error())
				return
			}
			cacaheBuffer := make([]byte,0)
			msg,ok := UnPack(buff[:rlen],&cacaheBuffer)
			if ok {
				utils.GO_Func(
					func() {
						utils.Try(func() {
							myfunc,exist := HandlerFunc[int(msg.HandleNo)]
							if exist{
								myfunc(se,msg.Body)
							}
						}).CatchAll(func(err error) {
							log.Printf("catch error %s",err.Error())
						})
					})
			}
		}
	}

}