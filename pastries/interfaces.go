// console 2
package pastries

//func (b *Bismarck) IsBargain() bool {
//	if b.HasCreamFilling() {
//		return b.Price < 3
//	} else {
//		return b.Price < 2
//	}
//}

type PurchasablePastry interface {
	Cost() int
	HasCreamFilling() bool
}

func (d *Donut) Cost() int {
	return d.Price
}

func IsBargain(p PurchasablePastry) bool {
	if p.HasCreamFilling() {
		return p.Cost() < 10
	} else {
		return p.Cost() < 5
	}
}
