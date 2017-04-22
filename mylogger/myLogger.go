package main
import "logger"
func main(){
	msg := "log msg is here"
    logger.Init("output.txt")
//	logger.Error(msg)
//	logger.Fatal(msg)
	logger.Info(msg)	
//	logger.Warning(msg)
	logger.Close()
}