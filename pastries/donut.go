package pastries

type Donut struct {
	Price     int
	Sprinkles bool
	Filling   bool
}

func (d *Donut) HasCreamFilling() bool {
	return d.Filling
}
func (d *Donut) IsBargain() bool {
	if d.HasCreamFilling() {
		return d.Price < 10
	} else {
		return d.Price < 5
	}
}

type Bismarck struct {
	Donut
	GoesBad int
}

func (b *Bismarck) HasCreamFilling() bool {
	return true
}
func (b *Bismarck) IsBargainNow(now int) bool {
	if !b.StillGood(now) {
		return false
	} else if b.HasCreamFilling() {
		return b.Price < 10
	} else {
		return b.Price < 5
	}
}
func (b *Bismarck) StillGood(now int) bool {
	return now < b.GoesBad
}

type GiftBox struct {
	Donut
	Bismarck
}
