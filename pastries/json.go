// console 2
package pastries

import (
	"encoding/json"
	"time"
)

// type alias
type Berliner Bismarck

type BismarckJSON struct {
	//Bismarck
	//		doesn't work: MarshalJSON infinite cycle

	Berliner

	Filling   byte `json:",omitempty"`
	Sprinkles byte `json:",omitempty"`

	EncodedAt time.Time
}

func (b *Bismarck) MarshalJSON() (data []byte, err error) {
	b2 := (*Berliner)(b)
	return json.Marshal(BismarckJSON{
		Berliner:  *b2,
		EncodedAt: time.Now(),
	})
}
