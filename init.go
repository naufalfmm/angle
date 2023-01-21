package angle

import (
	"math"

	"github.com/naufalfmm/angle/angleType"
	"github.com/naufalfmm/angle/angleUnit"
	"github.com/naufalfmm/angle/consts"
)

var Zero = NewDegreeFromFloat(consts.DecimalZero)

func NewDegreeFromFloat(val float64) Angle {
	return Angle{
		degree:  math.Abs(val),
		neg:     val < 0,
		angType: angleType.Decimal,
		angUnit: angleUnit.Degree,
	}
}

func NewRadianFromFloat(val float64) Angle {
	return Angle{
		degree:  math.Abs(val),
		neg:     val < 0,
		angType: angleType.Decimal,
		angUnit: angleUnit.Radian,
	}
}

func NewDegreeFromString(str string) (Angle, error) {
	var deg Angle

	if err := deg.scanByString(str); err != nil {
		return Angle{}, err
	}

	return deg, nil
}

func NewFromDegreeMinuteSecond(degree, minute, second float64) Angle {
	neg := false

	if degree < 0 {
		neg = true
		degree = math.Abs(degree)
	}

	if minute < 0 {
		neg = true
		minute = math.Abs(minute)
	}

	if second < 0 {
		neg = true
		second = math.Abs(second)
	}

	return Angle{
		degree: degree,
		minute: minute,
		second: second,

		neg:     neg,
		angType: angleType.DegreeMinuteSecond,
		angUnit: angleUnit.Degree,
	}
}
