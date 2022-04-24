//接口嵌套
package main

import "fmt"

type Fly interface {
	fly()
}

type Swimmer interface {
	swimmer()
}

type flySwimmer interface {
	Fly
	Swimmer
}

type Bird struct {
	name string
}

func (bird Bird) fly() {
	fmt.Println(bird.name, "实现fly方法...")

}

func (bird Bird) swimmer() {
	fmt.Println(bird.name, "实现swimmer方法...")

}
	

func main() {
	var bird flySwimmer
	bird = Bird {
		name : "麻雀",
	}
	bird.fly()
	bird.swimmer()
}