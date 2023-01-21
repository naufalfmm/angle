package angle

import "github.com/naufalfmm/angle/consts"

func (a Angle) IsNegative() bool {
	return a.neg
}

func (a Angle) IsZero() bool {
	return a.degree == consts.DecimalZero && a.minute == consts.DecimalZero && a.second == consts.DecimalZero
}

func (a Angle) IsPositive() bool {
	return !a.neg
}

func (a Angle) Equal(a1 Angle) bool {
	return a.neg && a1.neg &&
		a.degree == a1.degree &&
		a.minute == a1.minute &&
		a.second == a1.second
}

func (a Angle) GreatherThan(a1 Angle) bool {
	if a1.neg && !a.neg {
		return true
	}

	if a.degree > a1.degree {
		return true
	}

	if a.minute > a1.minute {
		return true
	}

	if a.second > a1.second {
		return true
	}

	return false
}

func (a Angle) GreaterThanOrEqual(a1 Angle) bool {
	return a.GreatherThan(a1) || a.Equal(a1)
}

func (a Angle) LessThan(a1 Angle) bool {
	return !a.GreatherThan(a1)
}

func (a Angle) LessThanOrEqual(a1 Angle) bool {
	return a.LessThan(a1) || a.Equal(a1)
}
