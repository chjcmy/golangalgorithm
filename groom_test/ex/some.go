package main

import (
	"fmt"
	"time"
)

var scheduler chan string

func consuming(prompt string) {
	fmt.Println("consuming 호출됨")
	select {
	case scheduler <- prompt:
		fmt.Println("이름을 입력받았습니다 : ", <-scheduler)
	case <-time.After(5 * time.Second):
		fmt.Println("시간이 지났습니다.")
	}
}
func producing(console chan string) {
	var name string
	fmt.Print("이름:")
	fmt.Scanln(&name)
	console <- name
}
func main() {
	console := make(chan string, 1)
	scheduler = make(chan string, 1)
	go producing(console)
	go func() { consuming(<-console) }()
	time.Sleep(100 * time.Second)
}

//func (p *MsgPipeRW) WriteMsg(msg Msg) error {
//	if atomic.LoadInt32(p.closed) == 0 {
//		consumed := make(chan struct{}, 1)
//		msg.Payload = &eofSignal{msg.Payload, msg.Size, consumed}
//		select {
//		case p.w <- msg:
//			if msg.Size > 0 {
//				select {
//				case <-consumed:
//				case <-p.closing:
//				}
//			}
//			return nil
//		case <-p.closing:
//		}
//	}
//	return ErrPipeClosed
//}
