package main

import (
	"fmt"
	self2 "github.com/XieWeiXie/self-system/example/pkg/self"
	self1 "github.com/XieWeiXie/self-system/refactor-system/pkg/self"
)

func main() {
	hello := self1.Hello{}
	fmt.Println(hello.Say())
	world := self2.World{}
	fmt.Println(world.Say())

}
