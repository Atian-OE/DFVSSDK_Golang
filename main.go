package main

import (
	"fmt"
	"github.com/Atian-OE/DFVSSDK_Golang/dfvssdk"
	"github.com/Atian-OE/DFVSSDK_Golang/dfvssdk/model"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main()  {

fmt.Println(time.Now())

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	fmt.Println("start")

	client:= dfvssdk.NewDFVSClient("192.168.0.95")

	client.CallConnected(func(addr string) {
		fmt.Println(fmt.Sprintf("连接成功:%s!",addr))
	})
	client.CallDisconnected(func(addr string) {
		fmt.Println(fmt.Sprintf("断开连接:%s!",addr))
	})
	time.Sleep(time.Second*3)



	err:= client.CallAlarmNotify(func(notify *model.AlarmEventNotify, e error) {
		fmt.Println("CallAlarmNotify"+notify.DeviceID)
	})
	fmt.Println("CallAlarmNotify",err)


	err= client.CallFiberStatusNotify(func(notify *model.FiberStateNotify, e error) {
		fmt.Println("CallFiberStatusNotify"+notify.DeviceID)
	})
	fmt.Println("CallFiberStatusNotify",err)





	<-ch
	client.Close()

	fmt.Println("quit")
}