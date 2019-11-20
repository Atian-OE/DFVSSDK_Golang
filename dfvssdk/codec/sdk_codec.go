package codec

import (
	"github.com/Atian-OE/DFVSSDK_Golang/dfvssdk/model"
	"github.com/Atian-OE/DFVSSDK_Golang/dfvssdk/utils"
	"github.com/golang/protobuf/proto"
)

//length 4 msg_id 1  包体 。。。
func Encode(msg_obj interface{}) ([]byte,error) {
	data ,err:= proto.Marshal(msg_obj.(proto.Message))
	if(err!=nil){
		return nil,err
	}
	cache := make([]byte, len(data)+5)
	length,_:= utils.IntToBytes(int64(len(data)),4)
	copy(cache,length)
	switch msg_obj.(type) {
	case *model.FiberStateNotify:
		cache[4]=byte(model.MsgID_FiberStateNotifyID)
	case *model.AlarmEventNotify:
		cache[4]=byte(model.MsgID_AlarmNotifyID)
	case *model.HeartBeat:
		cache[4]=byte(model.MsgID_HeartBeatID)
	}
	copy(cache[5:],data)

	return cache,err
}

