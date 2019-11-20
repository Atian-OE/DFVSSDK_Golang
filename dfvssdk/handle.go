package dfvssdk

import (
	"errors"
	"github.com/Atian-OE/DFVSSDK_Golang/dfvssdk/model"
	"github.com/golang/protobuf/proto"
	"net"
)

func (self*DFVSSDKClient)tcp_handle(msg_id model.MsgID, data []byte,conn net.Conn)  {

	switch (msg_id) {
	case model.MsgID_ConnectID:
		self.connected=true
		if(self._connected_action!=nil){
			go self._connected_action(self.addr)
		}

	case model.MsgID_DisconnectID:
		self.connected=false


		if(self._disconnected_action!=nil){
			go self._disconnected_action(self.addr)
		}

	case model.MsgID_AlarmNotifyID:

		if self._AlarmNotify!=nil{
			reply:=model.AlarmEventNotify{}
			err:=proto.Unmarshal(data[5:], &reply)
			self._AlarmNotify(&reply,err)
		}


	case model.MsgID_FiberStateNotifyID:
		if self._FiberStatusNotify!=nil{
			reply:=model.FiberStateNotify{}
			err:=proto.Unmarshal(data[5:], &reply)
			self._FiberStatusNotify(&reply,err)
		}

	}
}


//回调连接到服务器
func (self*DFVSSDKClient)CallConnected(call func(string))  {
	self._connected_action=call
}

//回调断开连接服务器
func (self*DFVSSDKClient)CallDisconnected(call func(string))  {
	self._disconnected_action=call
}

func (self *DFVSSDKClient) CallAlarmNotify(call func(*model.AlarmEventNotify, error)) error {
	if(call==nil){
		return errors.New("callback func is nil")
	}
	if(!self.connected){
		return errors.New("client not connected")
	}
	self._AlarmNotify=call

	return nil
}

func (self *DFVSSDKClient) CallFiberStatusNotify(call func(*model.FiberStateNotify,error)) error {
	if(call==nil){
		return errors.New("callback func is nil")
	}
	if(!self.connected){
		return errors.New("client not connected")
	}
	self._FiberStatusNotify=call

	return nil
}