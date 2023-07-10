package main

import "fmt"

type VideoPlayer interface {
	Play(videoName string)
}

type oldPlayer struct {
}

func (o *oldPlayer) Play(videoName string) {
	fmt.Println("Playing with an old player: ", videoName)
}

type newPlayer struct {
}

func (n *newPlayer) Start(videoName string) {
	fmt.Println("Playing with a new player: ", videoName)
}

type newPlayerAdapter struct {
	newPlayer *newPlayer
}

func (a *newPlayerAdapter) Play(videoName string) {
	a.newPlayer.Start(videoName)
}

func main() {
	a := &newPlayerAdapter{
		newPlayer: &newPlayer{},
	}
	a.Play("123")
}
