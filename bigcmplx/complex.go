package bigcmplx

import (
	"fmt"
	"math/big"
)

// A big rational complex number type, copying the big.Rat API.
// Only Cartesian operations are supported.
type Complex struct {
	real, imag big.Rat
}

func NewComplex(real, imag *big.Rat) *Complex {
	return new(Complex).SetRat(real, imag)
}

// Real returns the real part of z.
func (z *Complex) Real() *big.Rat {
	return &z.real
}

// Imag returns the imaginary part of z.
func (z *Complex) Imag() *big.Rat {
	return &z.imag
}

// SetRat sets z to real + imag*i and returns z.
func (z *Complex) SetRat(real, imag *big.Rat) *Complex {
	z.real.Set(real)
	z.imag.Set(imag)
	return z
}

// SetComplex128 sets z to exactly x and returns z.
// If x is not finite, SetComplex128 panics.
func (z *Complex) SetComplex128(x complex128) *Complex {
	z.real.SetFloat64(real(x))
	z.imag.SetFloat64(imag(x))
	return z
}

// Conj sets z to the complex conjugate of x and returns z.
func (z *Complex) Conj(x *Complex) *Complex {
	z.real.Set(&x.real)
	z.imag.Neg(&x.imag)
	return z
}

// Neg sets z to -x and returns x.
func (z *Complex) Neg(x *Complex) *Complex {
	z.real.Neg(&x.real)
	z.imag.Neg(&x.imag)
	return z
}

// SqAbs sets z to |x|^2, which will have imaginary part 0, and returns z.
func (z *Complex) SqAbs(x *Complex) *Complex {
	var conj Complex
	conj.Conj(x)
	z.Mul(z, &conj)
	return z
}

// Inv sets z to x^-1 and returns z.
// If |x| == 0, Inv panics.
func (z *Complex) Inv(x *Complex) *Complex {
	var a Complex
	a.SqAbs(x)
	a.real.Inv(&a.real)
	z.Conj(z)
	z.Mul(z, &a)
	return z
}

// Add sets z to the sum x+y and returns z.
func (z *Complex) Add(x, y *Complex) *Complex {
	z.real.Add(&x.real, &y.real)
	z.imag.Add(&x.imag, &y.imag)
	return z
}

// Sub sets z to the difference x-y and returns z.
func (z *Complex) Sub(x, y *Complex) *Complex {
	z.Neg(y)
	z.Add(x, z)
	return z
}

// Mul sets z to the product x*y and returns z.
func (z *Complex) Mul(x, y *Complex) *Complex {
	var ac, ad, bc, bd big.Rat
	ac.Mul(&x.real, &y.real)
	bc.Mul(&x.imag, &y.real)
	ad.Mul(&x.real, &y.imag)
	bd.Mul(&x.imag, &y.imag)
	z.real.Sub(&ac, &bd)
	z.imag.Add(&bc, &ad)
	return z
}

// Div sets z to the quotient x/y and returns z.
// If |y| == 0, Inv panics.
func (z *Complex) Quo(x, y *Complex) *Complex {
	z.Inv(y)
	z.Mul(x, z)
	return z
}

// String returns a string representation of z in the form "a/b + c/d i".
func (z *Complex) String() string {
	return fmt.Sprintf("%s + %s i", z.real.String(), z.imag.String())
}
