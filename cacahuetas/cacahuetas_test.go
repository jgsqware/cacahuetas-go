package cacahuetas

import "testing"

var restrictedTest = []struct {
	ina, inb string
	out      bool
}{
	{ina: "Julien", inb: "Céline", out: true},
	{ina: "Céline", inb: "Julien", out: true},
	{ina: "Ludiine", inb: "Jem", out: false},
	{ina: "Julien", inb: "Jem", out: false},
}

func TestIsRestricted(t *testing.T) {
	restrictions = Restrictions{"Julien": "Céline"}

	for _, v := range restrictedTest {
		restricted := isRestricted(v.ina, v.inb)
		if restricted != v.out {
			t.Errorf("Expected {%q,%q} restricted = %v, got %v", v.ina, v.inb, v.out, restricted)
		}
	}

}
