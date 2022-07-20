//接口
package main

import (
	"fmt"
)

type TapePlayer struct {
	Batteries string
}

func (t TapePlayer) Play(song string) {
	fmt.Println("Playing", song)
}

func (t TapePlayer) Stop() {
	fmt.Println("Stopped!")
}

type TapeRecorder struct {
	Microphones int
}

func (t TapeRecorder) Play(song string) {
	fmt.Println("Playing", song)
}

func (t TapeRecorder) Stop() {
	fmt.Println("Stopped!")
}

func (t TapeRecorder) Record() {
	fmt.Println("Recording")
}

func playList(p Player, songs []string) {
	for _, song := range songs {
		p.Play(song)
	}
	p.Stop()
}

type Player interface {
	Play(string)
	Stop()
}

type Vehicle interface {
	Accelerate()
	Brack()
	Steer(string)
}


func TryOut(p Player) {
	p.Play("Test Track")
	p.Stop()
	recorder, ok := p.(TapeRecorder)
	if ok {
		recorder.Record()
	} else {
		fmt.Println("Player was not a TapeRecorder")
	}
	
}

func main() {
	mixtape := []string{
		"Jessie's Girl",
		"whip It",
		"Love Girl",
	}
	var player Player
	player = TapePlayer{}
	playList(player, mixtape)
	
	player = TapeRecorder{}
	playList(player, mixtape)
	TryOut(player)
	
}