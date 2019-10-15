```
package main

import ("fmt"
		"math")

type Vertex struct{
	X, Y float64
} //Make a structure with float64 fields

func (v Vertex) Abs() float64{
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}/*Go has no classes. Instead make a method with a pointer receiver of type Vertex*/

func main(){
	v := Vertex{3, 4}
		fmt.Println(v)
		fmt.Println(v.Abs())
}
```
