package pastries

func NewDonut(price int, sprinkles, filling bool) *Donut {
	return &Donut{Price: price, Sprinkles: sprinkles, Filling: filling}
}
func NewBismarck(price, goesBad int) *Bismarck {
	return &Bismarck{Donut{Price: price}, goesBad}
}
func NewGiftBox(d *Donut, b *Bismarck) *GiftBox {
	return &GiftBox{Donut: *d, Bismarck: *b}
}
