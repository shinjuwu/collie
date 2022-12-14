package g_test

import (
	"fmt"
	"time"

	g "github.com/shinjuwu/collie/go"
)

func Example() {
	d := g.New(10)

	// go 1
	var res int
	d.Go(func() {
		fmt.Println("1 + 1 = ?")
		res = 1 + 1
	}, func() {
		fmt.Println(res)
	})

	d.Cb(<-d.ChanCb)

	// go 2
	d.Go(func() {
		fmt.Print("My name is ")
	}, func() {
		fmt.Println("github.com/shinjuwu/collie")
	})

	d.Close()

	// Output:
	// 1 + 1 = ?
	// 2
	// My name is github.com/shinjuwu/collie
}

func ExampleLinearContext() {
	d := g.New(10)

	// parallel
	d.Go(func() {
		time.Sleep(time.Second / 2)
		fmt.Println("1")
	}, nil)
	d.Go(func() {
		fmt.Println("2")
	}, nil)

	d.Cb(<-d.ChanCb)
	d.Cb(<-d.ChanCb)

	// linear
	c := d.NewLinearContext()
	c.Go(func() {
		time.Sleep(time.Second / 2)
		fmt.Println("1")
	}, nil)
	c.Go(func() {
		fmt.Println("2")
	}, nil)

	d.Close()

	// Output:
	// 2
	// 1
	// 1
	// 2
}
