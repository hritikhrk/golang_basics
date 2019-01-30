package interfaces

import "fmt"

type geometry interface {
	area() float64;
	perim() float64;
}

type rect struct {
	width, height float64
}
type circle struct {
	radius float64
}

func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}
func main()  {
	r:=rect{width:10,height:18}
	measure(r);
}