package angle

import "github.com/naufalfmm/angle/consts"

func (a Angle) prepareConvertMinuteSecond() Angle {
	if a.second >= consts.TimeFormatConverter {
		a.second = a.second - consts.TimeFormatConverter
		a.minute = a.minute + consts.DecimalOne
	}

	if a.minute >= consts.TimeFormatConverter {
		a.minute = a.minute - consts.TimeFormatConverter
		a.degree = a.degree + consts.DecimalOne
	}

	return a
}
