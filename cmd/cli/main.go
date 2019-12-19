package main

import (
	"fmt"
)

func main() {
	gacha := &Gacha{ticket: "normal_ticket_01"}
	execute(gacha)

	var aaa Drawable

	aaa = gacha

	fmt.Println(aaa)
}

func execute(d Drawer) {
	d.Draw()
}

// Drawable ...
type Drawable interface {
	Draw()
}

// Drawer ...
type Drawer interface {
	Drawable
}

// Gacha ...
type Gacha struct {
	ticket string
}

// Draw ...
func (x *Gacha) Draw() {
	fmt.Println(x.ticket + "で景品ゲット！")
}
