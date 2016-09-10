// console 2
package pastries

type MysteryPastry struct {
	PurchasablePastry
}

func (m MysteryPastry) IsBargain() bool {
	return IsBargain(m)
}

type UltimatePastry struct {
	MysteryPastry
}

func (u UltimatePastry) IsBargain() bool {
	return u.Cost() < 1000
}
