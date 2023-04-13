package main

import "fmt"

type Girl interface {
	service()
}

type RealGirl struct{}

func (r RealGirl) service() {
	fmt.Println("Em sẽ phục vụ anh đêm nay!nay")
}

type FakeGirl struct{}

func (r FakeGirl) service() {
	fmt.Println("Em là cú có gai, đi không?")
}

type Customer struct {
}

func (c Customer) book(girl Girl) {
	girl.service()
}

func main() {
	thanhDan := Customer{}

	var girl Girl

	girl = RealGirl{}
	thanhDan.book(girl)

	girl = FakeGirl{}
	thanhDan.book(girl)
}
