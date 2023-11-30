package main

import "fmt"

type HomeMovieFacade struct {
	tv       *TVScreen
	internet *Internet
}

type TVScreen struct {
}

type Internet struct {
}

func (i *Internet) connectToWifi() {

}

func (t *TVScreen) display() {

}

func (h *HomeMovieFacade) PlayMovie(name string) {
	h.tv.display()
	h.internet.connectToWifi()
	fmt.Printf("BE READY TO WATCH 1XBET AD!!!!!! before the %s starts\n", name)
}

func main() {
	tv := &TVScreen{}
	i := &Internet{}
	homemovie := &HomeMovieFacade{tv, i}
	homemovie.PlayMovie("INTERSTELLAR")
}
