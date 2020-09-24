package calc

import (
	"errors"
	"github.com/stretchr/testify/require"
	"testing"
)

type TestCase struct {
	name string
	expr string
	res  float64
	err  error
}

func TestCheckSuccess(t *testing.T) {
	cases := []TestCase{
		TestCase{
			name: "Empty input",
			expr: "",
			res:  0,
			err:  nil,
		},
		TestCase{
			name: "Input contains only spaces",
			expr: "",
			res:  0,
			err:  nil,
		},
		TestCase{
			name: "Single value",
			expr: "11",
			res:  11,
			err:  nil,
		},
		TestCase{
			name: "Sum",
			expr: "2+2",
			res:  4,
			err:  nil,
		},
		TestCase{
			name: "Sub",
			expr: "10-2",
			res:  8,
			err:  nil,
		},
		TestCase{
			name: "Mult",
			expr: "2*2",
			res:  4,
			err:  nil,
		},
		TestCase{
			name: "Del",
			expr: "100/25",
			res:  4,
			err:  nil,
		},
		TestCase{
			name: "Priority",
			expr: "2*(3+1)+4",
			res:  12,
			err:  nil,
		},
		TestCase{
			name: "Brackets",
			expr: "((((2)*()(3+1))+4))",
			res:  12,
			err:  nil,
		},
		TestCase{
			name: "Spaces",
			expr: "   1 + 2  + 3 / 3   + 7   ",
			res:  11,
			err:  nil,
		},
		TestCase{
			name: "Float output",
			expr: "1 / 2 * 4 / 10 + 1",
			res:  1.2,
			err:  nil,
		},
		TestCase{
			name: "Float input",
			expr: "(1.2 + 2.3 - 4.5) * 10",
			res:  -10,
			err:  nil,
		},
		TestCase{
			name: "All operators",
			expr: "1/2*4+(4-5*3)/2*(3+20)+124.5",
			res:  0,
			err:  nil,
		},
	}

	for _, tc := range cases {
		res, err := Calculate(tc.expr)
		require.Equal(t, res, tc.res, tc.name)
		require.Equal(t, err, tc.err, tc.name)
	}
}

func TestCheckFail(t *testing.T) {
	cases := []TestCase{
		TestCase{
			name: "Only operators",
			expr: "+-*/",
			err:  errors.New("Wrong input data"),
		},
		TestCase{
			name: "Wrong count of operators",
			expr: "12+-1",
			err:  errors.New("Wrong input data"),
		},
		TestCase{
			name: "Not binary operation",
			expr: "-12",
			err:  errors.New("Wrong input data"),
		},
		TestCase{
			name: "Wrong count of operators into brackets",
			expr: "(1 + * 1)",
			err:  errors.New("Wrong input data"),
		},
		// TestCase{
		// 	name: "",
		// 	expr: "1 1",
		// 	err:  errors.New("Wrong input data"),
		// },
		TestCase{
			name: "A few points in float number",
			expr: "1 + 1.2.2",
		},
		TestCase{
			name: "Comma instead dot in number",
			expr: "1 + 1,2",
		},
	}

	for _, tc := range cases {
		res, err := Calculate(tc.expr)
		require.Equal(t, res, float64(0), tc.name)
		if tc.err == nil {
			require.NotEqual(t, err, tc.err, tc.name)
		} else {
			require.Equal(t, err, tc.err, tc.name)
		}

	}
}
