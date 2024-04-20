package tone

import (
	"math"
	"strconv"

	"github.com/govalues/decimal"
)

// Trunc truncates a  decimal to have no more than n digits.
func Trunc(f float32, n int) float32 {
	if f == 0 || n <= 0 {
		return 0
	}

	s := strconv.FormatFloat(float64(f), 'f', -1, 32)
	d := decimal.MustParse(s)

	numDigits := d.Prec()
	numRight := d.Scale()
	numLeft := numDigits - numRight

	// If the number has more digits in it than we want, we need to truncate it.
	if numDigits > n {

		// If the left side has all the digits that we want, then we can treat the number as an
		// integer and do some easy math to truncate it. Otherwise, we'll continue to parse it as a
		// float and truncate the decimal places.
		if numLeft >= n {
			m := int32(f)
			pow := int32(math.Pow10(numLeft - n))
			m /= pow
			m *= pow
			f = float32(m)
		} else {
			// Calculate how many decimal places we need to keep, and truncate the number.
			decPlaces := n - numLeft
			d = d.Trunc(decPlaces)

			f64, _ := d.Float64()
			f = float32(f64)
		}
	}

	return f
}
