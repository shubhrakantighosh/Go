package model

type Grade string

const (
	A Grade = "1"
	B Grade = "2"
	C Grade = "3"
	D Grade = "4"
)

func GradeCheck(grade Grade) Grade {
	switch grade {
	case "A":
		return A
	case "B":
		return B
	case "C":
		return C
	default:
		return D
	}
}
