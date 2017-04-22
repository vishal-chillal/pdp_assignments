package main
import "fmt"

type graphicEditor struct {}
func (editor *graphicEditor) draw(shape Shape){
	if 1 == shape.getType(){
		fmt.Println("Rectangle Drawn")
	}
	if 2 == shape.getType(){
		fmt.Println("Circle Drawn")
	}
}

type Shape interface {
	getType() int
}

type Rectangle struct{}
func (rec *Rectangle) getType() int {
	return 2
}

type Circle struct{}
func (cir *Circle) getType() int {
	return 1
}
func main(){

	circle := &Circle{}
	rectangle := &Rectangle{}	
	gEditor := graphicEditor{}


	gEditor.draw(rectangle)
	gEditor.draw(circle)
}