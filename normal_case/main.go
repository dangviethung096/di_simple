package main

import "fmt"

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

func (c Customer) bookRealGirl() {
	girl := RealGirl{}
	girl.service()
}

func (c Customer) bookFakeGirl() {
	girl := FakeGirl{}
	girl.service()
}

func main() {
	thanhDan := Customer{}

	thanhDan.bookRealGirl()
	thanhDan.bookFakeGirl()
}
