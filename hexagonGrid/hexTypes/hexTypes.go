package hexTypes

import (
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
	LetterToNumMap = map[string]int{
		"A": 0,
		"B": 1,
		"C": 2,
		"D": 3,
		"E": 4,
		"F": 5,
		"G": 6,
		"H": 7,
		"I": 8,
		"J": 9,
		"K": 10,
		"L": 11,
		"M": 12,
		"N": 13,
		"O": 14,
		"P": 15,
		"Q": 16,
		"R": 17,
		"S": 18,
		"T": 19,
		"U": 20,
		"V": 21,
		"W": 22,
		"X": 23,
		"Y": 24,
		"Z": 25,
	}

	NumToLetterMap = map[int]string{
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

type Location struct {
	Col int
	Row int
}

func (l *Location) GetHexName() string {
	letter, ok := NumToLetterMap[l.Col]
	if !ok {
		panic("Build better letter system")
	}

	return fmt.Sprintf("%s%02d", letter, l.Row + 1)
}

func GetHexName(x, y int) string {
	letter, ok := NumToLetterMap[x]
	if !ok {
		panic("Build better letter system")
	}

	return fmt.Sprintf("%s%02d", letter, y + 1)
}