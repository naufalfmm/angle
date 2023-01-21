package trig

import (
	"math"

	"github.com/naufalfmm/angle"
)

func Sin(a angle.Angle) float64 {
	return math.Sin(a.ToRadian().ToDecimal().ToFloat())
}

func Cos(a angle.Angle) float64 {
	return math.Cos(a.ToRadian().ToDecimal().ToFloat())
}

func Tan(a angle.Angle) float64 {
	return math.Tan(a.ToRadian().ToDecimal().ToFloat())
}

func Asin(f float64) angle.Angle {
	return angle.NewRadianFromFloat(math.Asin(f))
}

func Acos(f float64) angle.Angle {
	return angle.NewRadianFromFloat(math.Acos(f))
}

func Atan(f float64) angle.Angle {
	return angle.NewRadianFromFloat(math.Atan(f))
}

func Atan2(y, x float64) angle.Angle {
	return angle.NewRadianFromFloat(math.Atan2(y, x))
}

func Cosec(a angle.Angle) float64 {
	return 1. / Sin(a)
}

func Sec(a angle.Angle) float64 {
	return 1. / Cos(a)
}

func Cot(a angle.Angle) float64 {
	return 1. / Tan(a)
}

func Acosec(f float64) angle.Angle {
	return angle.NewRadianFromFloat(math.Asin(1. / f))
}

func Asec(f float64) angle.Angle {
	return angle.NewRadianFromFloat(math.Acos(1. / f))
}

func Acot(f float64) angle.Angle {
	return angle.NewRadianFromFloat(math.Atan(1. / f))
}
