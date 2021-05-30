package netpkg

import (
	"bytes"
	"main/utils"
)

/**
 * Created by @CaomaoBoy on 2021/5/22.
 *  email:<115882934@qq.com>
 */

const  MagicNum = 0x77
const HeaderLen = 5

type Msg struct {
	MagicNum int8 //MagicNum
	//Vout int32  //msg Vout
	Len int32 //Data Len maxLen
	HandleNo int32 //handle protol
	Body []byte //Body data
}

func Pack(body []byte,HandleNo int) []byte{
	return bytes.Join(
		[][]byte{
			{MagicNum},
			utils.ToBytes(int32(4 + len(body))),
			utils.ToBytes(int32(HandleNo)),
			body},
		[]byte{})
}
func UnPack(data []byte,cacheBuffer *[]byte,msg *Msg) bool{
	var newData []byte
	var length int32
	if len(*cacheBuffer) == 0{
		if len(data) <= HeaderLen{
			*cacheBuffer = append(*cacheBuffer,data...)
			return false
		}
		validMagic := data[0]
		if validMagic != MagicNum{
			utils.NetErroWarp("Invalid Data Flow")
		}
		length = utils.BytesToInt32(data[1:HeaderLen])
		surPlus := int(length) - (len(data) - HeaderLen)
		if surPlus > 0 {
			*cacheBuffer = append(*cacheBuffer,data...)
			return false
		}else if surPlus == 0 {
			newData = data
		}else if surPlus < 0 {
			*cacheBuffer = append(*cacheBuffer, data[length:]...)
			newData = data[:length]
		}
	}else{
		validMagic := (*cacheBuffer)[0]
		if validMagic != MagicNum{
			utils.NetErroWarp("Invalid Data Flow")
		}
		if len(*cacheBuffer) < HeaderLen{
			n := len(*cacheBuffer)
			*cacheBuffer = append(*cacheBuffer, data[:HeaderLen-n]...)
			data = data[HeaderLen-n:]
		}
		length = utils.BytesToInt32((*cacheBuffer)[1:HeaderLen])
		totalLen := len(data) + len(*cacheBuffer)
		if totalLen < int(length){
			*cacheBuffer = append(*cacheBuffer, data...)
			return false
		}else if totalLen >= int(length){
			nlen := int(length)  - len(*cacheBuffer) + HeaderLen
			needData := data[:nlen]
			tmp := append(*cacheBuffer,needData...)
			*cacheBuffer = data[nlen:]
			newData = tmp
		}
	}
	msg.MagicNum = MagicNum
	msg.Len = length
	msg.HandleNo = utils.BytesToInt32(newData[HeaderLen:HeaderLen+4])
	msg.Body = newData[HeaderLen + 4:]
	return true
}