package main

type Square struct {
	a Point
	b Point
	c Point
	d Point
}

func (square Square) show()  {
	println("	a:" + square.a.show())
	println("	b:" + square.b.show())
	println("	c:" + square.c.show())
	println("	d:" + square.d.show())
}