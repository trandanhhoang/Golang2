package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64
func (e ErrNegativeSqrt) Error() string{
	return fmt.Sprint("cannot revert nega with sqrt",float64(e))
}

const e = 1e-8  // small delta

func Sqrt(x float64) (float64, error) {
	// typ,oke := x.(complex128)
	if x < 0{
		e := ErrNegativeSqrt(-x)
		return float64(e),e
		// Sqrt should return a non-nil error value when given a negative number, as it doesn't support complex numbers.
	}
	// else if oke{
		// return 1.0,nil
	// }
	z := x  // starting point
    for {
        new_z := z - ((z*z - x) / (2*z))
        if math.Abs(new_z - z) < e {
            return new_z, nil
        }
        z = new_z
    }
	// return x, nil

}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
	// fmt.Println(Sqrt(2+2i))
}
// and make it an error by giving it a method 
// such that ErrNegativeSqrt(-2).Error() returns "cannot Sqrt negative number: -2".
// Note: A call to fmt.Sprint(e) inside the Error method will send the program into an infinite loop.
//  You can avoid this by converting e first: fmt.Sprint(float64(e)). Why?