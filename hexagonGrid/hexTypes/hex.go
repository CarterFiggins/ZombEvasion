package hexTypes

import (
	"errors"
	"fmt"
)

const (
	DangerousSectorName string = "Dangerous Sector"
	EscapeSectorName = "Escape Sector"
	HumanSectorName = "Human Sector"
	SecureSectorName = "Secure Sector"
	WallSectorName = "Wall Sector"
	ZombieSectorName = "Zombie Sector"
)

var (
	LetterMap = map[int]string{
		0: "A",
		1: "B",
		2: "C",
		3: "D",
		4: "E",
		5: "F",
		6: "G",
		7: "H",
		8: "I",
		9: "J",
		10: "K",
		11: "L",
		12: "M",
		13: "N",
		14: "O",
		15: "P",
		16: "Q",
		17: "R",
		18: "S",
		19: "T",
		20: "U",
		21: "V",
		22: "W",
		23: "X",
		24: "Y",
		25: "Z",
	}
)


func HexName(x, y int) (string, error) {
	letter, ok := LetterMap[x]
	if !ok {
		return "", errors.New("no letter found in map!")
	}

	return fmt.Sprintf("%s%02d", letter, y), nil
}