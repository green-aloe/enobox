package tone

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

// Test_Consts tests any package-level constant values.
func Test_Consts(t *testing.T) {
	require.Equal(t, 20, NumHarmGains)
}

// Test_NewTone tests that NewTone returns a Tone that has been initialized correctly.
func Test_NewTone(t *testing.T) {
	tone := NewTone()
	require.NotEmpty(t, tone)
	require.IsType(t, Tone{}, tone)
	require.Equal(t, float32(0), tone.Frequency)
	require.Equal(t, float32(0), tone.Gain)
	require.Len(t, tone.HarmonicGains, NumHarmGains)
	for _, gain := range tone.HarmonicGains {
		require.Equal(t, float32(0), gain)
	}
}

// Test_NewToneAt tests that NewToneAt returns a Tone that has been initialized with the correct
// fundamental frequency.
func Test_NewToneAt(t *testing.T) {
	type subtest struct {
		frequency float32
		name      string
	}

	subtests := []subtest{
		{0, "zero frequency"},
		{-10, "negative frequency"},
		{10, "positive frequency"},
		{23.1, "non-integer frequency"},
		{440, "A4 frequency"},
	}

	for _, subtest := range subtests {
		t.Run(subtest.name, func(t *testing.T) {
			tone := NewToneAt(subtest.frequency)
			require.NotEmpty(t, tone)
			require.IsType(t, Tone{}, tone)
			require.Equal(t, subtest.frequency, tone.Frequency)
			require.Equal(t, float32(0), tone.Gain)
			require.Len(t, tone.HarmonicGains, NumHarmGains)
			for _, gain := range tone.HarmonicGains {
				require.Equal(t, float32(0), gain)
			}
		})
	}
}

// Test_Trunc tests that Trunc correctly truncates various numbers.
func Test_Trunc(t *testing.T) {
	type subtest struct {
		f    float32
		n    int
		want string
	}

	subtests := []subtest{
		{0, -1, "0"}, {0, 0, "0"}, {0, 1, "0"},
		{0.0, -1, "0"}, {0.0, 0, "0"}, {0.0, 1, "0"},
		{3, -1, "0"}, {3, 0, "0"}, {3, 1, "3"}, {3, 2, "3"}, {3, 3, "3"}, {3, 4, "3"}, {3, 5, "3"}, {3, 6, "3"}, {3, 7, "3"}, {3, 8, "3"},
		{3.1, -1, "0"}, {3.1, 0, "0"}, {3.1, 1, "3"}, {3.1, 2, "3.1"}, {3.1, 3, "3.1"}, {3.1, 4, "3.1"}, {3.1, 5, "3.1"}, {3.1, 6, "3.1"}, {3.1, 7, "3.1"}, {3.1, 8, "3.1"},
		{3.14, -1, "0"}, {3.14, 0, "0"}, {3.14, 1, "3"}, {3.14, 2, "3.1"}, {3.14, 3, "3.14"}, {3.14, 4, "3.14"}, {3.14, 5, "3.14"}, {3.14, 6, "3.14"}, {3.14, 7, "3.14"}, {3.14, 8, "3.14"},
		{3.141, -1, "0"}, {3.141, 0, "0"}, {3.141, 1, "3"}, {3.141, 2, "3.1"}, {3.141, 3, "3.14"}, {3.141, 4, "3.141"}, {3.141, 5, "3.141"}, {3.141, 6, "3.141"}, {3.141, 7, "3.141"}, {3.141, 8, "3.141"},
		{3.1415, -1, "0"}, {3.1415, 0, "0"}, {3.1415, 1, "3"}, {3.1415, 2, "3.1"}, {3.1415, 3, "3.14"}, {3.1415, 4, "3.141"}, {3.1415, 5, "3.1415"}, {3.1415, 6, "3.1415"}, {3.1415, 7, "3.1415"}, {3.1415, 8, "3.1415"},
		{3.14159, -1, "0"}, {3.14159, 0, "0"}, {3.14159, 1, "3"}, {3.14159, 2, "3.1"}, {3.14159, 3, "3.14"}, {3.14159, 4, "3.141"}, {3.14159, 5, "3.1415"}, {3.14159, 6, "3.14159"}, {3.14159, 7, "3.14159"}, {3.14159, 8, "3.14159"},
		{3.141592, -1, "0"}, {3.141592, 0, "0"}, {3.141592, 1, "3"}, {3.141592, 2, "3.1"}, {3.141592, 3, "3.14"}, {3.141592, 4, "3.141"}, {3.141592, 5, "3.1415"}, {3.141592, 6, "3.14159"}, {3.141592, 7, "3.141592"}, {3.141592, 8, "3.141592"},
		{3.1415926, -1, "0"}, {3.1415926, 0, "0"}, {3.1415926, 1, "3"}, {3.1415926, 2, "3.1"}, {3.1415926, 3, "3.14"}, {3.1415926, 4, "3.141"}, {3.1415926, 5, "3.1415"}, {3.1415926, 6, "3.14159"}, {3.1415926, 7, "3.141592"}, // {3.1415926, 8, "3.1415926"},
		{0.1, -1, "0"}, {0.1, 0, "0"}, {0.1, 1, "0.1"}, {0.1, 2, "0.1"}, {0.1, 3, "0.1"}, {0.1, 4, "0.1"}, {0.1, 5, "0.1"}, {0.1, 6, "0.1"}, {0.1, 7, "0.1"}, {0.1, 8, "0.1"},
		{0.01, -1, "0"}, {0.01, 0, "0"}, {0.01, 1, "0.01"}, {0.01, 2, "0.01"}, {0.01, 3, "0.01"}, {0.01, 4, "0.01"}, {0.01, 5, "0.01"}, {0.01, 6, "0.01"}, {0.01, 7, "0.01"}, {0.01, 8, "0.01"},
		{0.001, -1, "0"}, {0.001, 0, "0"}, {0.001, 1, "0.001"}, {0.001, 2, "0.001"}, {0.001, 3, "0.001"}, {0.001, 4, "0.001"}, {0.001, 5, "0.001"}, {0.001, 6, "0.001"}, {0.001, 7, "0.001"}, {0.001, 8, "0.001"},
		{0.0001, -1, "0"}, {0.0001, 0, "0"}, {0.0001, 1, "0.0001"}, {0.0001, 2, "0.0001"}, {0.0001, 3, "0.0001"}, {0.0001, 4, "0.0001"}, {0.0001, 5, "0.0001"}, {0.0001, 6, "0.0001"}, {0.0001, 7, "0.0001"}, {0.0001, 8, "0.0001"},
		{0.00001, -1, "0"}, {0.00001, 0, "0"}, {0.00001, 1, "0.00001"}, {0.00001, 2, "0.00001"}, {0.00001, 3, "0.00001"}, {0.00001, 4, "0.00001"}, {0.00001, 5, "0.00001"}, {0.00001, 6, "0.00001"}, {0.00001, 7, "0.00001"}, {0.00001, 8, "0.00001"},
		{0.000001, -1, "0"}, {0.000001, 0, "0"}, {0.000001, 1, "0.000001"}, {0.000001, 2, "0.000001"}, {0.000001, 3, "0.000001"}, {0.000001, 4, "0.000001"}, {0.000001, 5, "0.000001"}, {0.000001, 6, "0.000001"}, {0.000001, 7, "0.000001"}, {0.000001, 8, "0.000001"},
		{0.0000001, -1, "0"}, {0.0000001, 0, "0"}, {0.0000001, 1, "0.0000001"}, {0.0000001, 2, "0.0000001"}, {0.0000001, 3, "0.0000001"}, {0.0000001, 4, "0.0000001"}, {0.0000001, 5, "0.0000001"}, {0.0000001, 6, "0.0000001"}, {0.0000001, 7, "0.0000001"}, {0.0000001, 8, "0.0000001"},
		{0.00000001, -1, "0"}, {0.00000001, 0, "0"}, {0.00000001, 1, "0.00000001"}, {0.00000001, 2, "0.00000001"}, {0.00000001, 3, "0.00000001"}, {0.00000001, 4, "0.00000001"}, {0.00000001, 5, "0.00000001"}, {0.00000001, 6, "0.00000001"}, {0.00000001, 7, "0.00000001"}, {0.00000001, 8, "0.00000001"},
		{0.000000001, -1, "0"}, {0.000000001, 0, "0"}, {0.000000001, 1, "0.000000001"}, {0.000000001, 2, "0.000000001"}, {0.000000001, 3, "0.000000001"}, {0.000000001, 4, "0.000000001"}, {0.000000001, 5, "0.000000001"}, {0.000000001, 6, "0.000000001"}, {0.000000001, 7, "0.000000001"}, {0.000000001, 8, "0.000000001"},
		{0.0000000001, -1, "0"}, {0.0000000001, 0, "0"}, {0.0000000001, 1, "0.0000000001"}, {0.0000000001, 2, "0.0000000001"}, {0.0000000001, 3, "0.0000000001"}, {0.0000000001, 4, "0.0000000001"}, {0.0000000001, 5, "0.0000000001"}, {0.0000000001, 6, "0.0000000001"}, {0.0000000001, 7, "0.0000000001"}, {0.0000000001, 8, "0.0000000001"},
		{1000.0001, -1, "0"}, {1000.0001, 0, "0"}, {1000.0001, 1, "1000"}, {1000.0001, 2, "1000"}, {1000.0001, 3, "1000"}, {1000.0001, 4, "1000"}, {1000.0001, 5, "1000"}, {1000.0001, 6, "1000"}, {1000.0001, 7, "1000"}, {1000.0001, 8, "1000.0001"}, {1000.0001, 9, "1000.0001"}, {1000.0001, 10, "1000.0001"},
		{10000.00001, -1, "0"}, {10000.00001, 0, "0"}, {10000.00001, 1, "10000"}, {10000.00001, 2, "10000"}, {10000.00001, 3, "10000"}, {10000.00001, 4, "10000"}, {10000.00001, 5, "10000"}, {10000.00001, 6, "10000"}, {10000.00001, 7, "10000"}, //{10000.00001, 8, "10000.00001"}, {10000.00001, 9, "10000.00001"}, {10000.00001, 10, "10000.00001"},
		{10000000.00000001, -1, "0"}, {10000000.00000001, 0, "0"}, {10000000.00000001, 1, "10000000"}, {10000000.00000001, 2, "10000000"}, {10000000.00000001, 3, "10000000"}, {10000000.00000001, 4, "10000000"}, {10000000.00000001, 5, "10000000"}, {10000000.00000001, 6, "10000000"}, {10000000.00000001, 7, "10000000"}, {10000000.00000001, 8, "10000000"}, {10000000.00000001, 9, "10000000"}, {10000000.00000001, 10, "10000000"}, {10000000.00000001, 11, "10000000"}, {10000000.00000001, 12, "10000000"}, {10000000.00000001, 13, "10000000"}, {10000000.00000001, 14, "10000000"}, {10000000.00000001, 15, "10000000"}, // {10000000.00000001, 16, "10000000.00000001"}, {10000000.00000001, 17, "10000000.00000001"}, {10000000.00000001, 18, "10000000.00000001"},
		{.00000000123456780, -1, "0"}, {.00000000123456780, 0, "0"}, {.00000000123456780, 1, "0.000000001"}, {.00000000123456780, 2, "0.0000000012"}, {.00000000123456780, 3, "0.00000000123"}, {.00000000123456780, 4, "0.000000001234"}, {.00000000123456780, 5, "0.0000000012345"}, {.00000000123456780, 6, "0.00000000123456"}, {.00000000123456780, 7, "0.000000001234567"}, {.00000000123456780, 8, "0.0000000012345678"},
		{0.0000000123456780, -1, "0"}, {0.0000000123456780, 0, "0"}, {0.0000000123456780, 1, "0.00000001"}, {0.0000000123456780, 2, "0.000000012"}, {0.0000000123456780, 3, "0.0000000123"}, {0.0000000123456780, 4, "0.00000001234"}, {0.0000000123456780, 5, "0.000000012345"}, {0.0000000123456780, 6, "0.0000000123456"}, {0.0000000123456780, 7, "0.00000001234567"}, {0.0000000123456780, 8, "0.000000012345678"},
		{00.000000123456780, -1, "0"}, {00.000000123456780, 0, "0"}, {00.000000123456780, 1, "0.0000001"}, {00.000000123456780, 2, "0.00000012"}, {00.000000123456780, 3, "0.000000123"}, {00.000000123456780, 4, "0.0000001234"}, {00.000000123456780, 5, "0.00000012345"}, {00.000000123456780, 6, "0.000000123456"}, {00.000000123456780, 7, "0.0000001234567"}, // {00.000000123456780, 8, "0.00000012345678"},
		{000.00000123456780, -1, "0"}, {000.00000123456780, 0, "0"}, {000.00000123456780, 1, "0.000001"}, {000.00000123456780, 2, "0.0000012"}, {000.00000123456780, 3, "0.00000123"}, {000.00000123456780, 4, "0.000001234"}, {000.00000123456780, 5, "0.0000012345"}, {000.00000123456780, 6, "0.00000123456"}, {000.00000123456780, 7, "0.000001234567"}, {000.00000123456780, 8, "0.0000012345678"},
		{0000.0000123456780, -1, "0"}, {0000.0000123456780, 0, "0"}, {0000.0000123456780, 1, "0.00001"}, {0000.0000123456780, 2, "0.000012"}, {0000.0000123456780, 3, "0.0000123"}, {0000.0000123456780, 4, "0.00001234"}, {0000.0000123456780, 5, "0.000012345"}, {0000.0000123456780, 6, "0.0000123456"}, {0000.0000123456780, 7, "0.00001234567"}, {0000.0000123456780, 8, "0.000012345678"},
		{00000.000123456780, -1, "0"}, {00000.000123456780, 0, "0"}, {00000.000123456780, 1, "0.0001"}, {00000.000123456780, 2, "0.00012"}, {00000.000123456780, 3, "0.000123"}, {00000.000123456780, 4, "0.0001234"}, {00000.000123456780, 5, "0.00012345"}, {00000.000123456780, 6, "0.000123456"}, {00000.000123456780, 7, "0.0001234567"}, {00000.000123456780, 8, "0.00012345678"},
		{000000.00123456780, -1, "0"}, {000000.00123456780, 0, "0"}, {000000.00123456780, 1, "0.001"}, {000000.00123456780, 2, "0.0012"}, {000000.00123456780, 3, "0.00123"}, {000000.00123456780, 4, "0.001234"}, {000000.00123456780, 5, "0.0012345"}, {000000.00123456780, 6, "0.00123456"}, {000000.00123456780, 7, "0.001234567"}, {000000.00123456780, 8, "0.0012345678"},
		{0000000.0123456780, -1, "0"}, {0000000.0123456780, 0, "0"}, {0000000.0123456780, 1, "0.01"}, {0000000.0123456780, 2, "0.012"}, {0000000.0123456780, 3, "0.0123"}, {0000000.0123456780, 4, "0.01234"}, {0000000.0123456780, 5, "0.012345"}, {0000000.0123456780, 6, "0.0123456"}, {0000000.0123456780, 7, "0.01234567"}, {0000000.0123456780, 8, "0.012345678"},
		{00000000.123456780, -1, "0"}, {00000000.123456780, 0, "0"}, {00000000.123456780, 1, "0.1"}, {00000000.123456780, 2, "0.12"}, {00000000.123456780, 3, "0.123"}, {00000000.123456780, 4, "0.1234"}, {00000000.123456780, 5, "0.12345"}, {00000000.123456780, 6, "0.123456"}, {00000000.123456780, 7, "0.1234567"}, {00000000.123456780, 8, "0.12345678"},
		{000000001.23456780, -1, "0"}, {000000001.23456780, 0, "0"}, {000000001.23456780, 1, "1"}, {000000001.23456780, 2, "1.2"}, {000000001.23456780, 3, "1.23"}, {000000001.23456780, 4, "1.234"}, {000000001.23456780, 5, "1.2345"}, {000000001.23456780, 6, "1.23456"}, {000000001.23456780, 7, "1.234567"}, {000000001.23456780, 8, "1.2345678"},
		{0000000012.3456780, -1, "0"}, {0000000012.3456780, 0, "0"}, {0000000012.3456780, 1, "10"}, {0000000012.3456780, 2, "12"}, {0000000012.3456780, 3, "12.3"}, {0000000012.3456780, 4, "12.34"}, {0000000012.3456780, 5, "12.345"}, {0000000012.3456780, 6, "12.3456"}, {0000000012.3456780, 7, "12.34567"}, {0000000012.3456780, 8, "12.345678"},
		{00000000123.456780, -1, "0"}, {00000000123.456780, 0, "0"}, {00000000123.456780, 1, "100"}, {00000000123.456780, 2, "120"}, {00000000123.456780, 3, "123"}, {00000000123.456780, 4, "123.4"}, {00000000123.456780, 5, "123.45"}, {00000000123.456780, 6, "123.456"}, {00000000123.456780, 7, "123.4567"}, {00000000123.456780, 8, "123.45678"},
		{000000001234.56780, -1, "0"}, {000000001234.56780, 0, "0"}, {000000001234.56780, 1, "1000"}, {000000001234.56780, 2, "1200"}, {000000001234.56780, 3, "1230"}, {000000001234.56780, 4, "1234"}, {000000001234.56780, 5, "1234.5"}, {000000001234.56780, 6, "1234.56"}, {000000001234.56780, 7, "1234.567"}, // {000000001234.56780, 8, "1234.5678"},
		{0000000012345.6780, -1, "0"}, {0000000012345.6780, 0, "0"}, {0000000012345.6780, 1, "10000"}, {0000000012345.6780, 2, "12000"}, {0000000012345.6780, 3, "12300"}, {0000000012345.6780, 4, "12340"}, {0000000012345.6780, 5, "12345"}, {0000000012345.6780, 6, "12345.6"}, {0000000012345.6780, 7, "12345.67"}, {0000000012345.6780, 8, "12345.678"},
		{00000000123456.780, -1, "0"}, {00000000123456.780, 0, "0"}, {00000000123456.780, 1, "100000"}, {00000000123456.780, 2, "120000"}, {00000000123456.780, 3, "123000"}, {00000000123456.780, 4, "123400"}, {00000000123456.780, 5, "123450"}, {00000000123456.780, 6, "123456"}, {00000000123456.780, 7, "123456.7"}, {00000000123456.780, 8, "123456.78"},
		{000000001234567.80, -1, "0"}, {000000001234567.80, 0, "0"}, {000000001234567.80, 1, "1000000"}, {000000001234567.80, 2, "1200000"}, {000000001234567.80, 3, "1230000"}, {000000001234567.80, 4, "1234000"}, {000000001234567.80, 5, "1234500"}, {000000001234567.80, 6, "1234560"}, {000000001234567.80, 7, "1234567"}, {000000001234567.80, 8, "1234567.8"},
		{0000000012345678.0, -1, "0"}, {0000000012345678.0, 0, "0"}, {0000000012345678.0, 1, "10000000"}, {0000000012345678.0, 2, "12000000"}, {0000000012345678.0, 3, "12300000"}, {0000000012345678.0, 4, "12340000"}, {0000000012345678.0, 5, "12345000"}, {0000000012345678.0, 6, "12345600"}, {0000000012345678.0, 7, "12345670"}, {0000000012345678.0, 8, "12345678"},
		{1, -1, "0"}, {1, 0, "0"}, {1, 1, "1"}, {1, 2, "1"}, {1, 3, "1"}, {1, 4, "1"}, {1, 5, "1"}, {1, 6, "1"}, {1, 7, "1"}, {1, 8, "1"}, {1, 9, "1"}, {1, 10, "1"},
		{12, -1, "0"}, {12, 0, "0"}, {12, 1, "10"}, {12, 2, "12"}, {12, 3, "12"}, {12, 4, "12"}, {12, 5, "12"}, {12, 6, "12"}, {12, 7, "12"}, {12, 8, "12"}, {12, 9, "12"}, {12, 10, "12"},
		{123, -1, "0"}, {123, 0, "0"}, {123, 1, "100"}, {123, 2, "120"}, {123, 3, "123"}, {123, 4, "123"}, {123, 5, "123"}, {123, 6, "123"}, {123, 7, "123"}, {123, 8, "123"}, {123, 9, "123"}, {123, 10, "123"},
		{1234, -1, "0"}, {1234, 0, "0"}, {1234, 1, "1000"}, {1234, 2, "1200"}, {1234, 3, "1230"}, {1234, 4, "1234"}, {1234, 5, "1234"}, {1234, 6, "1234"}, {1234, 7, "1234"}, {1234, 8, "1234"}, {1234, 9, "1234"}, {1234, 10, "1234"},
		{12345, -1, "0"}, {12345, 0, "0"}, {12345, 1, "10000"}, {12345, 2, "12000"}, {12345, 3, "12300"}, {12345, 4, "12340"}, {12345, 5, "12345"}, {12345, 6, "12345"}, {12345, 7, "12345"}, {12345, 8, "12345"}, {12345, 9, "12345"}, {12345, 10, "12345"},
		{123456, -1, "0"}, {123456, 0, "0"}, {123456, 1, "100000"}, {123456, 2, "120000"}, {123456, 3, "123000"}, {123456, 4, "123400"}, {123456, 5, "123450"}, {123456, 6, "123456"}, {123456, 7, "123456"}, {123456, 8, "123456"}, {123456, 9, "123456"}, {123456, 10, "123456"},
		{1234567, -1, "0"}, {1234567, 0, "0"}, {1234567, 1, "1000000"}, {1234567, 2, "1200000"}, {1234567, 3, "1230000"}, {1234567, 4, "1234000"}, {1234567, 5, "1234500"}, {1234567, 6, "1234560"}, {1234567, 7, "1234567"}, {1234567, 8, "1234567"}, {1234567, 9, "1234567"}, {1234567, 10, "1234567"},
		{12345678, -1, "0"}, {12345678, 0, "0"}, {12345678, 1, "10000000"}, {12345678, 2, "12000000"}, {12345678, 3, "12300000"}, {12345678, 4, "12340000"}, {12345678, 5, "12345000"}, {12345678, 6, "12345600"}, {12345678, 7, "12345670"}, {12345678, 8, "12345678"}, {12345678, 9, "12345678"}, {12345678, 10, "12345678"},
		{123456789, -1, "0"}, {123456789, 0, "0"}, {123456789, 1, "100000000"}, {123456789, 2, "120000000"}, {123456789, 3, "123000000"}, {123456789, 4, "123400000"}, {123456789, 5, "123450000"}, {123456789, 6, "123456000"}, {123456789, 7, "123456700"}, // {123456789, 8, "123456780"}, {123456789, 9, "123456789"}, {123456789, 10, "123456789"},
		{1.1, -1, "0"}, {1.1, 0, "0"}, {1.1, 1, "1"}, {1.1, 2, "1.1"}, {1.1, 3, "1.1"}, {1.1, 4, "1.1"}, {1.1, 5, "1.1"}, {1.1, 6, "1.1"}, {1.1, 7, "1.1"}, {1.1, 8, "1.1"}, {1.1, 9, "1.1"}, {1.1, 10, "1.1"},
		{12.1, -1, "0"}, {12.1, 0, "0"}, {12.1, 1, "10"}, {12.1, 2, "12"}, {12.1, 3, "12.1"}, {12.1, 4, "12.1"}, {12.1, 5, "12.1"}, {12.1, 6, "12.1"}, {12.1, 7, "12.1"}, {12.1, 8, "12.1"}, {12.1, 9, "12.1"}, {12.1, 10, "12.1"},
		{123.1, -1, "0"}, {123.1, 0, "0"}, {123.1, 1, "100"}, {123.1, 2, "120"}, {123.1, 3, "123"}, {123.1, 4, "123.1"}, {123.1, 5, "123.1"}, {123.1, 6, "123.1"}, {123.1, 7, "123.1"}, {123.1, 8, "123.1"}, {123.1, 9, "123.1"}, {123.1, 10, "123.1"},
		{1234.1, -1, "0"}, {1234.1, 0, "0"}, {1234.1, 1, "1000"}, {1234.1, 2, "1200"}, {1234.1, 3, "1230"}, {1234.1, 4, "1234"}, {1234.1, 5, "1234.1"}, {1234.1, 6, "1234.1"}, {1234.1, 7, "1234.1"}, {1234.1, 8, "1234.1"}, {1234.1, 9, "1234.1"}, {1234.1, 10, "1234.1"},
		{12345.1, -1, "0"}, {12345.1, 0, "0"}, {12345.1, 1, "10000"}, {12345.1, 2, "12000"}, {12345.1, 3, "12300"}, {12345.1, 4, "12340"}, {12345.1, 5, "12345"}, {12345.1, 6, "12345.1"}, {12345.1, 7, "12345.1"}, {12345.1, 8, "12345.1"}, {12345.1, 9, "12345.1"}, {12345.1, 10, "12345.1"},
		{123456.1, -1, "0"}, {123456.1, 0, "0"}, {123456.1, 1, "100000"}, {123456.1, 2, "120000"}, {123456.1, 3, "123000"}, {123456.1, 4, "123400"}, {123456.1, 5, "123450"}, {123456.1, 6, "123456"}, {123456.1, 7, "123456.1"}, {123456.1, 8, "123456.1"}, {123456.1, 9, "123456.1"}, {123456.1, 10, "123456.1"},
		{1234567.1, -1, "0"}, {1234567.1, 0, "0"}, {1234567.1, 1, "1000000"}, {1234567.1, 2, "1200000"}, {1234567.1, 3, "1230000"}, {1234567.1, 4, "1234000"}, {1234567.1, 5, "1234500"}, {1234567.1, 6, "1234560"}, {1234567.1, 7, "1234567"}, {1234567.1, 8, "1234567.1"}, {1234567.1, 9, "1234567.1"}, {1234567.1, 10, "1234567.1"},
		{12345678.1, -1, "0"}, {12345678.1, 0, "0"}, {12345678.1, 1, "10000000"}, {12345678.1, 2, "12000000"}, {12345678.1, 3, "12300000"}, {12345678.1, 4, "12340000"}, {12345678.1, 5, "12345000"}, {12345678.1, 6, "12345600"}, {12345678.1, 7, "12345670"}, {12345678.1, 8, "12345678"}, //{12345678.1, 9, "12345678.1"}, {12345678.1, 10, "12345678.1"},
		{123456789.1, -1, "0"}, {123456789.1, 0, "0"}, {123456789.1, 1, "100000000"}, {123456789.1, 2, "120000000"}, {123456789.1, 3, "123000000"}, {123456789.1, 4, "123400000"}, {123456789.1, 5, "123450000"}, {123456789.1, 6, "123456000"}, {123456789.1, 7, "123456700"}, // {123456789.1, 8, "123456780"}, {123456789.1, 9, "123456789"}, {123456789.1, 10, "123456789.1"},

		// Test the commented-out values above that are too large to be truncated and actually end
		// up getting rounded a bit because of precision loss.
		{3.1415926, 8, "3.1415925"},
		{10000.00001, 8, "10000"}, {10000.00001, 9, "10000"}, {10000.00001, 10, "10000"},
		{10000000.00000001, 16, "10000000"}, {10000000.00000001, 17, "10000000"}, {10000000.00000001, 18, "10000000"},
		{00.000000123456780, 8, "0.00000012345679"},
		{000000001234.56780, 8, "1234.5677"},
		{123456789, 8, "123456790"}, {123456789, 9, "123456790"}, {123456789, 10, "123456790"},
		{12345678.1, 9, "12345678"}, {12345678.1, 10, "12345678"},
		{123456789.1, 8, "123456790"}, {123456789.1, 9, "123456790"}, {123456789.1, 10, "123456790"},
	}

	for _, subtest := range subtests {
		from := strconv.FormatFloat(float64(subtest.f), 'f', -1, 32)
		t.Run(fmt.Sprintf("%v|%v|%v", from, subtest.n, subtest.want), func(t *testing.T) {
			// Test the positive value.
			want := subtest.want
			have := Trunc(subtest.f, subtest.n)
			s := strconv.FormatFloat(float64(have), 'f', -1, 32)
			require.Equal(t, want, s)

			// Test the negative value.
			if want != "0" {
				want = "-" + subtest.want
			}
			have = Trunc(-subtest.f, subtest.n)
			s = strconv.FormatFloat(float64(have), 'f', -1, 32)
			require.Equal(t, want, s)
		})
	}
}
