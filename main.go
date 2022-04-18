package main

import "fmt"

// var c, python, java bool
// var i, j int = 1, 2

// var (
// 	ToBe   bool       = false
// 	MaxInt uint64     = 1<<64 - 1
// 	z      complex128 = cmplx.Sqrt(-5 + 12i)
// )

// const Pi = 3.14

const (
	Big   = 1 << 100
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}
func main() {
	// var c, python, java = true, false, "no!"
	// var i, j int = 1, 2
	// c, python, java := true, false, "no!"
	// fmt.Println(i, j, c, python, java)
	// fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	// fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	// fmt.Printf("Type: %T Value: %v\n", z, z)

	// var i int
	// var f float64
	// var b bool
	// var s string
	// fmt.Printf("%v %v %v %q\n", i, f, b, s)

	// var x, y int = 3, 4
	// var f float64 = math.Sqrt(float64(x*x + y*y))
	// var z uint = uint(f)
	// fmt.Println(x, y, z)

	// v := 42
	// fmt.Printf("v is of type %T\n", v)
	// const World = "世界"
	// fmt.Println("Hello", World)
	// fmt.Println("Happy", Pi, "Day")
	// const Truth = true
	// fmt.Println("Go rules?", Truth)

	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}
