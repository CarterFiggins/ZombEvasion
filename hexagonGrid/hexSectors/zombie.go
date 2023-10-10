package hexSectors

import (
	"image/color"

	"github.com/tdewolff/canvas"
)

type Zombie struct {
	Name string
	*Location
}

func (s Zombie) GetColor() color.Color {
	return canvas.Lawngreen 
}

func (s Zombie) GetStrokeColor() color.Color {
	return canvas.Black
}

func (s *Zombie) GetSectorName() string {
	s.Name = ZombieName
	return ZombieName
}

func (s Zombie) GetText() (*canvas.Text, error) {
	fontFamily := canvas.NewFontFamily("times")
	if err := fontFamily.LoadSystemFont("Nimbus Roman, serif", canvas.FontRegular); err != nil {
		return nil, err
	}

	face := fontFamily.Face(8.0, canvas.Black, canvas.FontRegular, canvas.FontNormal)
	return canvas.NewTextLine(face, "Z", canvas.Center), nil
}

func (s *Zombie) SetLocation(col int, row int) {
	s.Location = &Location{
		Col: col,
		Row: row,
	}
}


func (s *Zombie) CanMoveHere() bool {
	return false
}
