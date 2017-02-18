package main
import "fmt"

type graphicEditor struct {}


func (editor *graphicEditor) draw(shape Shape){
	shape.draw()
}

type Shape interface {
	draw()
}


type Rectangle struct{}
func (rec *Rectangle) draw() {
	fmt.Println("rectangle Drawn")
}

type Circle struct{}
func (cir *Circle) draw(){
fmt.Println("circlele Drawn")	
}

type Triangle struct{}
func (tri *Triangle) draw(){
fmt.Println("Triangle Drawn")	
}

func main(){

	circle := &Circle{}
	rectangle := &Rectangle{}
	triangle := &Triangle{}	
	gEditor := graphicEditor{}


	gEditor.draw(rectangle)
	gEditor.draw(circle)
	gEditor.draw(triangle)
}