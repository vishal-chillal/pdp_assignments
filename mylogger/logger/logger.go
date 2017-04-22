package logger
import (
		"fmt"
		"os"
		"bytes"
		"time"
		)

var fp *os.File

var myLevel = [...]string {"Info","Warning","Error","Fatal"}

const(
    Info_msg = iota
    Warning_msg 
    Error_msg
    Fatal_msg
)

func print_log(level int, msg string) int {
    str := FormatDateTime(time.Now())
    
    for i := level; i < len(myLevel); i++ {
        myMsg := ""+str+"\t"+myLevel[i]+"\t"+msg+"\n"
    	n,err := fp.WriteString(myMsg)
        if err != nil {
            panic(err)
            return n
        }
        n=n+1
    }
    return 0
}

func Close() {
    fp.Close()    
}

func Init(fName string) {

	f,err := os.OpenFile(fName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
    if err != nil {
        panic(err)
    }
    fp = f
}

func Warning(msg string) {
	print_log(Warning_msg,msg)
}

func Info(msg string) {
	print_log(Info_msg,msg)
}
func Error(msg string) {
    print_log(Error_msg,msg)    
}

func Fatal(msg string) {
	print_log(Fatal_msg,msg)
}

func FormatDateTime(t time.Time) string {
    var buffer bytes.Buffer
    buffer.WriteString("["+t.Month().String()[:3])
    buffer.WriteString(fmt.Sprintf(" %2d'%2dat%2d:%2d:%2d]", t.Day(), t.Year(), t.Hour(), t.Minute(), t.Second()))
    return buffer.String()
}
