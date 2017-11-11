package triangle

import "math"

const testVersion = 3

func KindFromSides(a, b, c float64) Kind {
	if IsATriangle(a, b, c) {
		if IsEquilateral(a, b, c) {
			return Equ
		}
		if IsIsoceles(a, b, c) {
			return Iso
		}
		return Sca
	}
	return NaT

}

func IsATriangle(a, b, c float64) bool {
	if math.IsInf(a, 0) || math.IsInf(b, 0) || math.IsInf(c, 0) {
		return false
	}
	if a <= 0 || b <= 0 || c <= 0 {
		return false
	}

	var aIsLesser, bIsLesser, cIsLesser bool
	if b+c >= a {
		aIsLesser = true
	} else {
		aIsLesser = false
	}
	if a+c >= b {
		bIsLesser = true
	} else {
		bIsLesser = false
	}
	if a+b >= c {
		cIsLesser = true
	} else {
		cIsLesser = false
	}
	return (aIsLesser && bIsLesser && cIsLesser)

}

func IsEquilateral(a, b, c float64) bool {
	return (a == b && b == c && a == c)
}

func IsIsoceles(a, b, c float64) bool {
	var bcEqual, acEqual, abEqual bool
	if b == c {
		bcEqual = true
	} else {
		bcEqual = false
	}
	if a == c {
		acEqual = true
	} else {
		acEqual = false
	}
	if a == b {
		abEqual = true
	} else {
		abEqual = false
	}
	return (bcEqual || acEqual || abEqual)
}

// Notice KindFromSides() returns this type. Pick a suitable data type.
type Kind string

const (
	NaT Kind = "NaT" // not a triangle
	Equ Kind = "Equ" // equilateral
	Iso Kind = "Iso" // isosceles
	Sca Kind = "Sca" // scalene
)
