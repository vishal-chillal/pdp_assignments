package main
import "fmt"

type graphicEditor struct {}


func (editor *graphicEditor) draw(shape Shape){
	shape.draw()
}
func (editor *graphicEditor) area(shape Shape){
	shape.area()
}


type Shape interface {
	draw()
	area()
}


type Rectangle struct{}
func (rec *Rectangle) draw() {
	fmt.Println("rectangle Drawn")
}
func (rec *Rectangle) area() {
	fmt.Println("area of triangle : 1/2 *base*height")
}

type Circle struct{}
func (cir *Circle) draw(){
fmt.Println("circle Drawn")	
}
func (cir *Circle) area(){
	fmt.Println("area of circle : Pi*r*r")
}

type Line struct{}
func (ln *Line) draw(){
fmt.Println("line Drawn")	
}
func (ln *Line) area(){
	fmt.Println("No Formula For area")
}


func main(){

	circle := &Circle{}
	rectangle := &Rectangle{}
	line := &Line{}
	gEditor := graphicEditor{}

	gEditor.draw(rectangle)
	gEditor.area(rectangle)

	gEditor.draw(circle)
	gEditor.area(circle)

	gEditor.draw(line)
	gEditor.area(line)
}