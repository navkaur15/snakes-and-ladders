package spot

type Spot interface {
	GetPos() int
	GetNextPos() int
	SetNextPos(pos int, posType string) bool
	GetPosType() string
}
type SpotData struct {
	pos     int
	nextPos int
	posType string
}

func (s *SpotData) GetPos() int {
	return s.pos
}

func (s *SpotData) GetNextPos() int {
	return s.nextPos
}
func (s *SpotData) SetNextPos(pos int, posType string) bool {
	//can add checks here
	//if snake pos < curr pos
	//if ladder,pos > curr pos
	//pos between 1 to 100 only
	if pos >= 1 && pos < 100 {
		s.nextPos = pos
		s.posType = posType
		return true
	}
	return false

}

func (s *SpotData) GetPosType() string {
	return s.posType
}

func CreateNewSpot(pos, nextPos int, posType string) Spot {
	return &SpotData{
		pos:     pos,
		nextPos: nextPos,
		posType: posType,
	}
}
