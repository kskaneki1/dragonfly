package prismarine

import "fmt"

type Prismarine struct {
	prismarine
}

type prismarine uint8

func Default() Prismarine {
	return Prismarine{prismarine(0)}
}

func Dark() Prismarine {
	return Prismarine{prismarine(1)}
}

func Bricks() Prismarine {
	return Prismarine{prismarine(2)}
}

func All() []Prismarine {
	return []Prismarine{Default(), Dark(), Bricks()}
}

func (p prismarine) Uint8() uint8 {
	return uint8(p)
}

func (p prismarine) Name() string {
	switch p {
	case 0:
		return "Prismarine"
	case 1:
		return "Dark Prismarine"
	case 2:
		return "Bricks Prismarine"
	}
	panic("unknown prismarine name")
}

// FromString ...
func (p prismarine) FromString(s string) (interface{}, error) {
	switch s {
	case "default":
		return Prismarine{prismarine(0)}, nil
	case "dark":
		return Prismarine{prismarine(1)}, nil
	case "bricks":
		return Prismarine{prismarine(2)}, nil
	}
	return nil, fmt.Errorf("unexpected prismarine type '%v' 'default','dark','brick'", s)
}

// String ...
func (p prismarine) String() string {
	switch p {
	case 0:
		return "default"
	case 1:
		return "dark"
	case 2:
		return "bricks"
	}
}
