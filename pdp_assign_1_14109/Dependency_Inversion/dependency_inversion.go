package main
import "fmt"

type msgserver struct {
	MSG messege
}
func (srv *msgserver) sendMSG(){

	fmt.Println(srv.MSG.send()+" is SENT")
}

func (srv *msgserver) setType(myType messege){
	srv.MSG = myType
}

type messege interface {
	send() string
}

type email struct{}
func ( *email) send() string{
	return "email"
}

type sms struct{}
func (*sms) send() string{
	return "msg"
}

func main(){

	Email := &email{}
	SMS := &sms{}
	srver := msgserver{}

	srver.setType(Email)
	srver.sendMSG()
	srver.setType(SMS)
	srver.sendMSG()

}