//------------------------------demo1-------------------------------------
package main

import (
	"fmt"
)

type persion interface {
	address() string
}

//城市信息
type cityInfo struct {
	city string
}

//实现cityInfo结构方法
func (c *cityInfo) address() string {
	return c.city
}

//名字
type info struct {
	name string
}

func (i *info) address() string {
	return i.name
}

func main() {
	c := cityInfo{"朝阳望京"}
	u := info{"张三"}
	p := []persion{&c, &u}

	for _, sh := range p {
		fmt.Println(sh.address())
	}

}


//---------------------------demo2---------------------------------------
package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

func main() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f  // a MyFloat 实现了 Abser
	a = &v // a *Vertex 实现了 Abser

	// 下面一行，v 是一个 Vertex（而不是 *Vertex）
	// 所以没有实现 Abser。
	a = v

	fmt.Println(a.Abs())
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

/**
接口类型是由一组方法定义的集合。

接口类型的值可以存放实现这些方法的任何值。

注意： 示例代码的 22 行存在一个错误。 由于 Abs 只定义在 *Vertex（指针类型）上， 所以 Vertex（值类型）不满足 Abser。
**/
