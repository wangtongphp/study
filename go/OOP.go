package main

type Handler interface {
	F1()
}

type C1 struct {
}

func (c *C1) F1() {

}

func main() {
	c1 := new(C1)
	dev(c1)
}

func dev(Handler) {

}
