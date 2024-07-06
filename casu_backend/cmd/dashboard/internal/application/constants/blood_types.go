package constants

// BloodType
type BloodType int

const (
	PositiveA  BloodType = 1
	NegativeA  BloodType = 2
	PositiveB  BloodType = 3
	NegativeB  BloodType = 4
	Positive0  BloodType = 5
	Negative0  BloodType = 6
	PositiveAB BloodType = 7
	NegativeAB BloodType = 8
)

var bloodTypes = map[string]BloodType{
	"PositiveA":  PositiveA,
	"NegativeA":  NegativeA,
	"PositiveB":  PositiveB,
	"NegativeB":  NegativeB,
	"Positive0":  Positive0,
	"Negative0":  Negative0,
	"PositiveAB": PositiveAB,
	"NegativeAB": NegativeAB,
}

func (s BloodType) String() string {
	switch s {
	case PositiveA:
		return "PositiveA"
	case NegativeA:
		return "NegativeA"
	case PositiveB:
		return "PositiveB"
	case NegativeB:
		return "NegativeB"
	case Positive0:
		return "Positive0"
	case Negative0:
		return "Negative0"
	case PositiveAB:
		return "PositiveAB"
	case NegativeAB:
		return "NegativeAB"

	default:
		return "Unknown"

	}
}

func GetBloodType(sportTypeName string) BloodType {
	return bloodTypes[sportTypeName]
}

type BloodTypeData struct {
	CanGive []BloodType
	CanTake []BloodType
}

var data = map[BloodType]BloodTypeData{
	Positive0:  BloodTypeData{CanGive: []BloodType{Positive0, PositiveA, PositiveB, PositiveAB}, CanTake: []BloodType{Negative0, Positive0}},
	PositiveB:  BloodTypeData{CanGive: []BloodType{PositiveB, PositiveAB}, CanTake: []BloodType{Negative0, Positive0, NegativeB, PositiveB}},
	PositiveA:  BloodTypeData{CanGive: []BloodType{PositiveA, PositiveAB}, CanTake: []BloodType{Negative0, Positive0, NegativeA, PositiveA}},
	NegativeB:  BloodTypeData{CanGive: []BloodType{NegativeB, PositiveB, NegativeAB, PositiveAB}, CanTake: []BloodType{Negative0, NegativeB}},
	Negative0:  BloodTypeData{CanGive: []BloodType{Negative0, Positive0, NegativeA, PositiveA, NegativeB, PositiveB, NegativeAB, PositiveAB}, CanTake: []BloodType{Negative0}},
	NegativeA:  BloodTypeData{CanGive: []BloodType{NegativeA, PositiveA, NegativeAB, PositiveAB}, CanTake: []BloodType{Negative0, NegativeA}},
	NegativeAB: BloodTypeData{CanGive: []BloodType{NegativeAB, PositiveAB}, CanTake: []BloodType{Negative0, NegativeA, NegativeB, NegativeAB}},
	PositiveAB: BloodTypeData{CanGive: []BloodType{PositiveAB}, CanTake: []BloodType{Negative0, Positive0, NegativeA, PositiveA, NegativeB, PositiveB, NegativeAB, PositiveAB}},
}

func GetBloodTypeData(bloodType BloodType) BloodTypeData {
	return data[bloodType]
}
