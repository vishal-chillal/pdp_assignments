package main
import "fmt"

type msgserver struct {
	msg sms
	mail email

}
func (srv *msgserver) sendMSG(){
	fmt.Println(srv.msg.send()+" is SENT")
}

func (srv *msgserver) sendEmail(){
	fmt.Println(srv.mail.send()+" is SENT")
	
}

func (srv *msgserver) setEmailType(myType email){
	srv.mail = myType
}
func (srv *msgserver) setMsgType(myType sms){
	srv.msg = myType
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

	Email := email{}
	SMS := sms{}
	srver := msgserver{}

	srver.setEmailType(Email)
	srver.sendEmail()

	srver.setMsgType(SMS)
	srver.sendMSG()

}