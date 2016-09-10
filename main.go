package main

import (
	"log"
	"mule/donut/pastries"
)

func main() {

	d := pastries.NewDonut(6, true, false) // price, sprinkles, filling
	log.Println("Simple Donut IsBargain example:",
		d.IsBargain())
	// false

	b := pastries.NewBismarck(6, 10) // price, goesBad
	log.Println("Bismarck filling method example:",
		b.HasCreamFilling(), b.Donut.HasCreamFilling())
	// true, false

	log.Println("Bismarck Price, IsBargain example:",
		b.Price, b.IsBargain())
	// 6, false

	log.Println("Bismarck IsBargainNow example:",
		b.IsBargainNow(0))
	// true

	g := pastries.NewGiftBox(d, b)
	// g.HasCreamFilling()
	//		ambiguous selector g.HasCreamFilling()

	log.Println("GiftBox StillGood example:",
		g.StillGood(0))
	// true

	mainPartTwo(d, b)
}
