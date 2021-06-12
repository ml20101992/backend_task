package processing

type BoxType int

const (
	FTYP BoxType = iota
	MOOV
	OTHER
)

//Method used to return the name of BoxType variables
func (bt BoxType) Name() string {
	names := []string{"ftyp", "moov", "other"}

	return names[bt]
}
