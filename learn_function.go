package main

import (
	"errors"
	"fmt"
	"math"
	"strconv"
)

func divAndRemainder(a int, b int) (int, int) {
	return a / b, a % b
}

func doubleFail(a int, arr [2]int, s string) {
	a = a * 2
	for i := 0; i < len(arr); i++ {
		arr[i] = arr[i] * 2
	}
	s = s + s

	fmt.Println("in DoubleFail:", a, arr, s)
}
func addNumbers(a int, b int) int {
	return a + b
}

func addOne(a int) int {
	return a + 1
}
func addTwo(a int) int {
	return a + 2
}
func printOperation(a int, f func(int) int) {
	fmt.Println(f(a))
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

const Delta = 0.0001

func isConverged(d float64) bool {
	if d < 0.0 {
		d = -d
	}
	if d < Delta {
		return true
	}
	return false
}

func Sqrt(x float64) float64 {
	z := 1.0
	tmp := 0.0
	for {
		tmp = z - (z*z-x)/2*z
		if d := tmp - z; isConverged(d) {
			return tmp
		}
		z = tmp
	}
	return z
}

func pow(x, n, lim float64) float64 {

	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >=%g\n", v, lim)
	}

	return lim
}

type Vertex struct {
	x int
	y int
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func exactIndex(s []int, k int) int {
	for i := 0; i < len(s); i++ {
		if s[i] == k {
			return i
		}
	}
	return 0
}

type Item struct {
	lat, long float64
}

var m = map[string]Item{
	"Bell Labs": Item{
		lat: 40.68433, long: -74.39967,
	},
	"Google": Item{
		lat: 37.42202, long: -122.08408,
	},
}

func WordCount(s string) map[string]int {
	return map[string]int{"x": 1}
}

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci(n int) int {

	firstNumber := 0
	secondNumber := 1
	result := 0

	if n == 0 {
		return 0
	}

	if n == 1 {
		return 1
	}

	for i := 2; i <= n; i++ {
		result = firstNumber + secondNumber
		firstNumber = secondNumber
		secondNumber = result

	}

	return result
}

type Foo struct {
	A int
	B string
}

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func (f Foo) String() string {
	return f.B + "Result"
}
func (f *Foo) Double() {
	f.A = f.A * 2
}

type myInt int

func (mi myInt) IsEven() bool {
	return mi%2 == 0
}

func (mi *myInt) Double() {
	*mi = *mi * 2
}

type Abser interface {
	abs() float64
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type I interface {
	M()
}

type T struct {
	S string
}

func (t T) M() {
	fmt.Println(t.S)
}
func divAndMod(a int, b int) (int, int, error) {
	if b == 0 {
		return 0, 0, errors.New("Cannot divide by zero")
	}
	return a / b, a % b, nil
}

type SuperHero struct {
	Name    string
	Age     int
	Address Address
}

type Address struct {
	Number int
	Street string
	City   string
}

type Movie struct {
	Name   string
	Rating float64
}

func (m *Movie) summary() string {
	r := strconv.FormatFloat(m.Rating, 'f', 1, 64)
	return m.Name + "," + r
}

type Sphere struct {
	Radious float64
}

func (s *Sphere) SurfaceArea() float64 {
	return float64(4) * math.Pi * (s.Radious * s.Radious)
}

func (s *Sphere) Volume() float64 {
	radiousCubed := s.Radious * s.Radious * s.Radious
	return (float64(4) / float64(3)) * math.Pi * radiousCubed
}

type Triangle struct {
	base   float64
	height float64
}

func (t *Triangle) Area() float64 {
	return 0.5 * (t.base * t.height)
}

func (t *Triangle) ChangeBase(f float64) {
	t.base = f
	return
}

func main() {

	t := Triangle{base: 3, height: 1}
	fmt.Println("Area", t.Area())
	t.ChangeBase(4)
	fmt.Println("base", t.base)

	s := Sphere{
		Radious: 5,
	}
	fmt.Println("SurfaceArea", s.SurfaceArea())
	fmt.Println("Volume", s.Volume())
	// a := Drink{
	// 	Name: "Lemonade",
	// 	Ice:  true,
	// }
	// b := Drink{
	// 	Name: "Lemonade",
	// 	Ice:  true,
	// }

	//fmt.Printf(reflect.TypeOf(a))

	//fmt.Printf(reflect.TypeOf(b))

	// if b == a {
	// 	fmt.Println("they are both equal")
	// }

	// fmt.Printf("%+v\n", a)
	// fmt.Printf("%+v\n", b)

	// m := Movie{
	// 	Name:   "Hot movie",
	// 	Rating: 10,
	// }

	// fmt.Println(m.Name, m.Rating)
	// fmt.Printf("%+v\n", m)
	// fmt.Println(m.summary())
	// e := SuperHero{
	// 	Name: "batman",
	// 	Age:  32,
	// 	Address: Address{
	// 		Number: 10007,
	// 		Street: "Mountain drive",
	// 		City:   "Gotham",
	// 	},
	// }
	// fmt.Printf("%+v\n", e)
	// b := &e
	// b.Name = "Moses"
	// fmt.Printf("%+v\n", b)

	// if len(os.Args) > 3 || len(os.Args) < 2 {
	// 	fmt.Println("Expected two Input Parameters")
	// 	os.Exit(1)
	// }
	// a, err := strconv.ParseInt(os.Args[1], 10, 64)
	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }
	// fmt.Println(a)

	// m := myInt(10)
	// fmt.Println(m)
	// f := Foo{
	// 	A: 10,
	// 	B: "Hello",
	// }

	// fmt.Println(f.String())
	// f.Double()
	// fmt.Println(f.String())
	// bob := `{"name": "Bob", "age": 30}`
	// var b Person

	// json.Unmarshal([]byte(bob), &b)
	// fmt.Println(b)
	// bob2, _ := json.Marshal(b)
	// fmt.Println(bob2)
	// fmt.Println(fibonacci(3))

	// delete(m, "Google")
	// fmt.Println(m)

	// elem, ok := m["Google"]

	// if ok {
	// 	fmt.Println(elem)
	// } else {
	// 	fmt.Println("not present")
	// }
	//Create a tic-tac-toe board
	// var pow = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	// for i, v := range pow {
	// 	fmt.Printf("2**%d = %d\n", i, v)
	// }
	// fmt.Println(exactIndex(pow, 5))
	// q := []int{2, 3, 5, 7, 11, 13}
	// fmt.Println(q)

	// r := []bool{true, false, true, true, false, true}
	// fmt.Println(r)

	// t := []struct {
	// 	i int
	// 	b bool
	// }{
	// 	{2, true},
	// 	{3, true},
	// 	{5, false},
	// 	{7, true},
	// 	{11, false},
	// }

	// fmt.Println(t)

	// names := [4]string{
	// 	"John",
	// 	"George",
	// 	"Paul",
	// 	"Ringo",
	// }

	// fmt.Println(names)
	// a := names[0:2]
	// b := names[1:3]
	// c := names[0:1]

	// fmt.Println(a)
	// fmt.Println(b)
	// fmt.Println(c)

	// primes := [6]int{2, 3, 5, 7, 11, 13}
	// var s []int = primes[1:4]
	// fmt.Println(s)
	// fmt.Println(primes)
	// b[0] = "XXX"
	// fmt.Println(a, b)
	// fmt.Println(names)
}
