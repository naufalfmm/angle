# Angle

Package "Angle" is form of angle data type

## Features
- Have 2 types, that are decimal and degree minute second
- Have 2 units, that are degree and radian
- Automatically convert to exact type and units in additional, substraction, multiplication, division, trigonometry operation, etc.
- Have trigonometry operations, that are sin, cos, tan, arcsin, arccos, arctan, arctan2, cosec, sec, cot, arccosec, arcsec, and arccot.
- Use directly in struct for accepting request from JSON

## Quick Start
Install the library by
```sh
go get github.com/naufalfmm/angle
```

```go
func main() {
    // Initialize angle from float
    decAngle := angle.NewDegreeFromFloat(-6.30286)
    radAngle := angle.NewRadianFromFloat(-11000,56593)
    minSecAngle := angle.NewFromDegreeMinuteSecond(-6, 18, 10.3)

    fmt.Println(decAngle.Add(radAngle))
    fmt.Println(decAngle.AddScalar(5.))
    fmt.Println(radAngle.Mul(2.5))

    fmt.Println(trig.Sin(minSecAngle))
}
```