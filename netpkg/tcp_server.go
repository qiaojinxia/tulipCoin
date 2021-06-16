package netpkg

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
	Address  string
	Port     string
	StopChan chan struct{}
}

func(ts *TcpServer) Close(){
	ts.StopChan <- struct{}{}
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
				case <- ts.StopChan:
					return
				}
			}
		})
		utils.GO_Func(func() {
			tcpAddr, err := net.ResolveTCPAddr("tcp4", fmt.Sprintf("%s:%s",ts.Address,ts.Port))
			if err != nil{
				panic(utils.NetErroWarp(err,""))
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
				panic(utils.NetErroWarp(err,""))
			}
			for{
				conn,err := listener.Accept()
				if err != nil{
					panic(utils.NetErroWarp(err,""))
				}
				utils.GO_Func(func() {
					Handle(conn,ts.StopChan)
				})
			}
		})
		time.Sleep(time.Second * 2)
		utils.Wg.Wait()
	}).Catch(utils.NetErroWarp(nil,""), func(err error) {
		log.Println(utils.ConverToJsonInfo(err))
	}).CatchAll(func(err error) {
		log.Println(err)
	})
}

func Handle(conn net.Conn,stopCh <-chan struct{}){
	se := &Session{conn: conn,ID: utils.GetUserID()}
	_SessionManger.AddSession(se)
	cacaheBuffer := make([]byte,0)
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
				panic(utils.NetErroWarp(err,""))
				return
			}
			//First connection verification

			msg := &Msg{}
			for {
				ok := UnPack(buff[:rlen],&cacaheBuffer,msg)
				if ok{
					break
				}
			}
			func() {
				//msg.Body

			}()

			utils.GO_Func(
					func() {
							myfunc,exist := ServerHandlerFunc[int(msg.HandleNo)]
							if exist{
								msgRsp, err := myfunc(se, msg.Body)
								if err != nil{
									panic(utils.BusinessErrorWarp(err,""))
								}
								conn.Write(msgRsp)
							}

					})
			}
		}
	}

