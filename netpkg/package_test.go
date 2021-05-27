package netpkg

import (
	"encoding/json"
	"fmt"
	"testing"
)

/**
 * Created by @CaomaoBoy on 2021/5/22.
 *  email:<115882934@qq.com>
 */

func Test_Package(t *testing.T) {
	Mpackage := Pack([]byte("8767687"),0)
	//fmt.Println(Mpackage)
	cacheBuff := make( []byte,0)
	//rand.Seed(time.Now().UnixNano())
	//rand.Intn(10)
	t1 := Mpackage[:3]
	t2 := Mpackage[3:6]
	t3 := Mpackage[6:len(Mpackage)]
	t3 = append(t3, []byte("ASDASD")...)
	xxxx1 := make(map[int][]byte)
	xxxx1[0] = t1
	xxxx1[1] = t2
	xxxx1[2] = t3
	var msg *Msg
	for i:=0;i<3;i++{
		ok := UnPack(xxxx1[i],&cacheBuff,msg )
		if ok{
			dt,err := json.Marshal(msg)
			if err != nil{
				panic(err)
			}
			fmt.Printf("%s",dt)
			break
		}
	}


}