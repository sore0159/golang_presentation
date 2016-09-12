// console 2
package main

import (
	"log"
	"mule/donut/pastries"
)

func mainPartTwo(d *pastries.Donut, b *pastries.Bismarck) {

	log.Println("Interface IsBargain example:",
		pastries.IsBargain(d), pastries.IsBargain(b))
	// false, true

	b2 := pastries.Berliner(*b)
	// b2.IsBargainNow(0)
	//		type pastries.Berliner has no field or method IsBargainNow

	log.Println("Berliner filling method example:",
		b2.HasCreamFilling(), b2.Donut.HasCreamFilling())
	// false, false

	if j, err := b.MarshalJSON(); err != nil {
		panic("errorchecking is just good sense")
	} else {
		log.Println("Bismarck JSON example:",
			string(j))
		// {"Price":6,"GoesBad":10,"EncodedAt":"2016-09-10T13:52:28.828493223-05:00"}
	}
	if j, err := b2.MarshalJSON(); err != nil {
		panic("errorchecking is just good sense")
	} else {
		log.Println("Berliner Faulty JSON example:",
			string(j))
	}
	// {"Price":6,"Sprinkles":false,"Filling":false}
}
