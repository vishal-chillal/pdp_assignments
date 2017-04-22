package main
import (
"fmt"
//	"errors"
)
type Vehicle interface{
	Deliver()
	getInfo() string
}
const(
	bike = iota
	cycle
)
var vehicles = []string{"Bike","Cycle"}
func Create_Vehicle(val string) Vehicle {
	for v:=0 ;v<len(vehicles); v++ {
		if(val == vehicles[v]){
			return new(v)
		}
	}
	return nil
}

type Bike struct{
	typeName string
}

func (bk *Bike)Deliver() {
	bk.typeName = " Bike "
}

func (bk *Bike)getInfo() string {
	return "Your " + bk.typeName + " Is Ready..!!! "
}

// type cycle struct{
// 	typeName string
// }

// func (cl *cycle)Deliver {
// 	sv.typeName = " Cycle "
// }

// func (cl *cycle)getInfo() string {
// 	return "Your " + cl.typeName + " Is Ready..!!! "
// }


func main() {
	myCar := Create_Vehicle("Bike")
	myCar.Deliver()
	fmt.Println(myCar.getInfo())
}