package dfvssdk

import (
	"bytes"
	"fmt"
	"github.com/Atian-OE/DFVSSDK_Golang/dfvssdk/codec"
	"github.com/Atian-OE/DFVSSDK_Golang/dfvssdk/model"
	"github.com/Atian-OE/DFVSSDK_Golang/dfvssdk/utils"
	"github.com/kataras/iris/core/errors"
	"net"
	"time"
)

type WaitPackStr struct {
	Key model.MsgID
	Timeout int64//毫秒
	Call *func(model.MsgID, []byte, net.Conn,error)
}



type DFVSSDKClient struct{
	sess *net.TCPConn
	connected bool

	heart_beat_ticker *time.Ticker//心跳包的发送
	heart_beat_ticker_over  chan interface{} //关闭心跳

	reconnect_ticker *time.Ticker//自动连接
	reconnect_ticker_over  chan interface{} //关闭自动连接



	addr string//地址

	_connected_action         func(string)                               //连接到服务器的回调
	_disconnected_action      func(string)                               //断开连接到服务器的回调
	_AlarmNotify         func(*model.AlarmEventNotify,error)   //分区警报通知
	_FiberStatusNotify       func(*model.FiberStateNotify,error) //设备状态通知
}

func NewDFVSClient(addr string) *DFVSSDKClient {
	conn:= &DFVSSDKClient{}
	conn.init(addr)
	return conn
}

func(self *DFVSSDKClient)init(addr string)  {
	self.addr=addr
	self.heart_beat_ticker= time.NewTicker(time.Second*5)
	self.heart_beat_ticker_over=make(chan interface{})
	self.reconnect_ticker=time.NewTicker(time.Second*10)
	self.reconnect_ticker_over=make(chan interface{})

	go self.heart_beat()
	go self.reconnect()


}


func (self *DFVSSDKClient)connect()  {
	if self.connected {
		return
	}

	conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:17084",self.addr),time.Second*3)
	if err!=nil {
		//fmt.Println("连接服务器失败!")
		return
	}
	tcpConn, ok := conn.(*net.TCPConn)
	if !ok {
		//fmt.Println("连接服务器失败!")
		return
	}
	self.sess=tcpConn
	//禁用缓存
	tcpConn.SetWriteBuffer(5000)
	tcpConn.SetReadBuffer(5000)
	go self.client_handle(tcpConn)
}

func (self *DFVSSDKClient)reconnect()  {
	self.connected=false
	self.connect()

	for{
		select {
		case <-self.reconnect_ticker.C:
			self.connect()

		case <-self.reconnect_ticker_over:
			return

		}
	}
}

//心跳
func (self *DFVSSDKClient)heart_beat() {
	for{
		select {
		case <-self.heart_beat_ticker.C:
			if(self.connected){
				b,_:=codec.Encode(&model.HeartBeat{})
				self.sess.Write(b)
			}

		case <-self.heart_beat_ticker_over:
			return

		}
	}
}


func (self *DFVSSDKClient)client_handle(conn net.Conn)  {
	self.tcp_handle(model.MsgID_ConnectID,nil,conn)
	defer func() {
		if(conn!=nil){
			self.tcp_handle(model.MsgID_DisconnectID,nil,conn)
			conn.Close()
		}
	}()


	buf:=make([]byte,1024)
	var cache bytes.Buffer
	for{
		//cache_index:=0
		n,err:=conn.Read(buf)
		//加上上一次的缓存
		//n=buf_index+n
		if(err!=nil){
			break
		}

		cache.Write(buf[:n])
		for{
			if(self.unpack(&cache,conn)){
				break
			}
		}
	}

}

// true 处理完成 false 循环继续处理
func (self *DFVSSDKClient)unpack(cache *bytes.Buffer,conn net.Conn) bool {
	if(cache.Len()<5){
		return true
	}
	buf:=cache.Bytes()
	pkg_size:= utils.ByteToInt2(buf[:4])
	//长度不够
	if(pkg_size > len(buf)-5){
		return true
	}

	cmd:=buf[4]
	self.tcp_handle(model.MsgID(cmd),buf[:pkg_size+5],conn)
	cache.Reset()
	cache.Write(buf[5+pkg_size:])

	return false
}




//发送消息
func (self*DFVSSDKClient)Send(msg_obj interface{}) error {
	b,err:=codec.Encode(msg_obj)
	if(err!=nil) {
		return err
	}
	if(!self.connected){
		return errors.New("client not connected")
	}
	_,err=self.sess.Write(b)
	return err
}


//关闭
func (self*DFVSSDKClient)Close() {

	self.reconnect_ticker.Stop()
	self.reconnect_ticker_over<-0
	close(self.reconnect_ticker_over)

	self.heart_beat_ticker.Stop()
	self.heart_beat_ticker_over<-0
	close(self.heart_beat_ticker_over)


	if(self.sess!=nil){
		self.sess.Close()
	}

	self.sess=nil
}