package angle

import (
	"bytes"
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/naufalfmm/angle/angleType"
	"github.com/naufalfmm/angle/angleUnit"
	"github.com/naufalfmm/angle/consts"
	"github.com/naufalfmm/angle/err"
)

type Angle struct {
	degree float64
	minute float64
	second float64

	neg     bool
	angType angleType.AngleType
	angUnit angleUnit.AngleUnit
}

func (a *Angle) fillBySymbol(src string, symbol rune) error {
	if symbol == consts.DegreeSymbolRune {
		if err := json.Unmarshal([]byte(src), &a.degree); err != nil {
			return err
		}
	}

	if symbol == consts.MinuteSymbolRune {
		if err := json.Unmarshal([]byte(src), &a.minute); err != nil {
			return err
		}
	}

	if err := json.Unmarshal([]byte(src), &a.second); err != nil {
		return err
	}

	return nil
}

func (a *Angle) decideFillBySymbol(src string, symbol rune) error {
	a.angType = angleType.Decimal
	if symbol == consts.MinuteSymbolRune ||
		symbol == consts.SecondSymbolRune {
		a.angType = angleType.DegreeMinuteSecond
	}

	return a.fillBySymbol(src, symbol)
}

func (a *Angle) scanByString(src string) error {
	var (
		buff bytes.Buffer
		s    rune
	)

	for _, s = range src {
		if s != consts.DegreeSymbolRune &&
			s != consts.MinuteSymbolRune &&
			s != consts.SecondSymbolRune {
			if _, err := buff.WriteRune(s); err != nil {
				return err
			}

			continue
		}

		if s == consts.NegativeSymbolRune {
			a.neg = true
			continue
		}

		if err := a.decideFillBySymbol(buff.String(), s); err != nil {
			return err
		}
	}

	return a.decideFillBySymbol(buff.String(), s)
}

func (a *Angle) UnmarshalParam(src string) error {
	return a.scanByString(src)
}

func (a *Angle) UnmarshalJSON(val []byte) error {
	var rawVal string
	if err := json.Unmarshal(val, &rawVal); err != nil {
		return err
	}

	return a.scanByString(rawVal)
}

func (a *Angle) Scan(val interface{}) error {
	rawVal, ok := val.([]byte)
	if !ok {
		return err.ErrConstantParsing
	}
	dbVal := string(rawVal)

	return a.scanByString(dbVal)
}

func (a Angle) String() string {
	var angStr string
	if a.neg {
		angStr = string(consts.NegativeSymbolRune)
	}

	if a.angType == angleType.DegreeMinuteSecond {
		return fmt.Sprintf("%s%s", angStr, strconv.FormatFloat(a.degree, 'f', -1, 64)+string(consts.DegreeSymbolRune)+
			strconv.FormatFloat(a.minute, 'f', -1, 64)+string(consts.MinuteSymbolRune)+
			strconv.FormatFloat(a.second, 'f', -1, 64)+string(consts.SecondSymbolRune))
	}

	angStr = fmt.Sprintf("%s%s", angStr, strconv.FormatFloat(a.degree, 'f', -1, 64))

	if a.angUnit == angleUnit.Degree {
		return fmt.Sprintf("%s%s", angStr, a.angUnit.Unit())
	}

	return fmt.Sprintf("%s%s", angStr, a.angUnit.Unit())
}

func (a Angle) MarshalJSON() ([]byte, error) {
	return json.Marshal(a.String())
}

func (a Angle) Value() (driver.Value, error) {
	return a.String(), nil
}

func (a Angle) AngleType() angleType.AngleType {
	return a.angType
}
