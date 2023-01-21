package err

import "errors"

var (
	ErrShouldBeDecimal = errors.New("type should be decimal")
	ErrShouldBeDegree  = errors.New("unit should be degree")
	ErrConstantParsing = errors.New("expected string for the constant")
)
