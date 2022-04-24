package main
import "fmt"

//接口和类型的关系
/* 
一个类型可以实现多个接口
一个接口可以被多个类型实现（多态）

*/

type music interface {
	playMusic()
}

type video interface {
	playVideo()
}

type Phone struct {
	name string
}

type Computers struct {
	name string
}

//一个类型实现多个接口
func (phone Phone) playMusic() {
	fmt.Println(phone.name,"playMusic...")
}

func (phone Phone) playVideo() {
	fmt.Println(phone.name,"playVideo...")
}

//一个接口被多个类型实现
func (com Computers) playMusic() {
	fmt.Println(com.name, "playMusic...")
}

func main() {
	phone := Phone{
		name : "iPhone",
	}
	phone.playMusic()
	phone.playVideo()

	com := Computers {
		name : "MacBookPro",
	}
	com.playMusic()
}