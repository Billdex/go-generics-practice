package higher_order_functions

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"testing"
)

type TestCase[T any, M any] struct {
	Name     string
	Input    T
	Want     M
	TestFunc func(T) M
}

func runTestCases[T any, M any](t *testing.T, cases []TestCase[T, M]) {
	for _, tt := range cases {
		t.Run(tt.Name, func(t *testing.T) {
			reality := tt.TestFunc(tt.Input)
			if !reflect.DeepEqual(tt.Want, reality) {
				t.Errorf("case %s fail! want: %v, reality: %v\n", tt.Name, tt.Want, reality)
			} else {
				t.Logf("case %s ok\n", tt.Name)
			}
		})
	}
}

func TestFilterSlice(t *testing.T) {
	casesInt := []TestCase[[]int, []int]{
		{
			Name:  "unsigned",
			Input: []int{1, 5, -1, 0, -2, 10, 99, -24, 6},
			Want:  []int{1, 5, 0, 10, 99, 6},
			TestFunc: func(input []int) []int {
				return FilterSlice(input, func(i int) bool { return i >= 0 })
			},
		},
		{
			Name:  "even",
			Input: []int{1, 5, -1, 0, -2, 10, 99, -24, 6},
			Want:  []int{0, -2, 10, -24, 6},
			TestFunc: func(input []int) []int {
				return FilterSlice(input, func(i int) bool { return i%2 == 0 })
			},
		},
	}
	casesString := []TestCase[[]string, []string]{
		{
			Name:  "long str",
			Input: []string{"12321", "222", "abcdef", "1234", "-1234", "%W@#", "你好我是", "咕噜咕噜帕"},
			Want:  []string{"12321", "abcdef", "-1234", "咕噜咕噜帕"},
			TestFunc: func(input []string) []string {
				return FilterSlice(input, func(s string) bool { return len([]rune(s)) >= 5 })
			},
		},
		{
			Name:  "has 0",
			Input: []string{"asdfas", "#$%#GD0", "12301d312", "", "0", "zero", "0000", "O0o0O"},
			Want:  []string{"#$%#GD0", "12301d312", "0", "0000", "O0o0O"},
			TestFunc: func(input []string) []string {
				return FilterSlice(input, func(s string) bool { return strings.Contains(s, "0") })
			},
		},
	}
	type Man struct {
		Name   string
		Height float64
		Weight int
	}
	casesStruct := []TestCase[[]Man, []Man]{
		{
			Name: "High BMI",
			Input: []Man{{Name: "Jack", Height: 1.71, Weight: 65}, {Name: "John", Height: 1.61, Weight: 68}, {Name: "Bob", Height: 1.88, Weight: 106},
				{Name: "Sam", Height: 1.66, Weight: 64}, {Name: "Philip", Height: 1.77, Weight: 87}, {Name: "Candy", Height: 1.72, Weight: 81}},
			Want: []Man{{Name: "John", Height: 1.61, Weight: 68}, {Name: "Bob", Height: 1.88, Weight: 106},
				{Name: "Philip", Height: 1.77, Weight: 87}, {Name: "Candy", Height: 1.72, Weight: 81}},
			TestFunc: func(input []Man) []Man {
				return FilterSlice(input, func(m Man) bool { return float64(m.Weight)/(m.Height*m.Height) > 26.0 })
			},
		},
	}
	runTestCases(t, casesInt)
	runTestCases(t, casesString)
	runTestCases(t, casesStruct)
}

func TestMapSlice(t *testing.T) {
	casesInt2Int := []TestCase[[]int, []int]{
		{
			Name:  "increase",
			Input: []int{1, 5, 2, 8, 32, 456, -22, 0, -1},
			Want:  []int{2, 6, 3, 9, 33, 457, -21, 1, 0},
			TestFunc: func(input []int) []int {
				return MapSlice(input, func(i int) int { return i + 1 })
			},
		},
		{
			Name:  "square",
			Input: []int{1, 4, -1, -5, 0, 33, 5, 2, 678},
			Want:  []int{1, 16, 1, 25, 0, 1089, 25, 4, 459684},
			TestFunc: func(input []int) []int {
				return MapSlice(input, func(i int) int { return i * i })
			},
		},
	}
	casesInt2String := []TestCase[[]int, []string]{
		{
			Name:  "reply",
			Input: []int{1, 2, 3, 0, -1, 10},
			Want:  []string{"11", "22", "33", "00", "-1-1", "1010"},
			TestFunc: func(input []int) []string {
				return MapSlice(input, func(i int) string {
					return fmt.Sprintf("%d%d", i, i)
				})
			},
		},
		{
			Name:  "unsigned",
			Input: []int{1, 5, 0, -2, 6, -100, 99, -1},
			Want:  []string{"1", "5", "0", "0", "6", "0", "99", "0"},
			TestFunc: func(input []int) []string {
				return MapSlice(input, func(i int) string {
					if i < 0 {
						return "0"
					}
					return strconv.Itoa(i)
				})
			},
		},
	}
	type Struct1 struct {
		Name  string
		Score int
		Level int
	}
	type Struct2 struct {
		SumScore int
		Data     string
	}
	casesStruct2Struct := []TestCase[[]Struct1, []Struct2]{
		{
			Name:  "cal score",
			Input: []Struct1{{Score: 2, Level: 1}, {Score: 20, Level: 5}, {Score: -10, Level: 5}},
			Want:  []Struct2{{SumScore: 2}, {SumScore: 100}, {SumScore: 0}},
			TestFunc: func(input []Struct1) []Struct2 {
				return MapSlice(input, func(s Struct1) Struct2 {
					if s.Score < 0 {
						return Struct2{}
					}
					return Struct2{
						SumScore: s.Score * s.Level,
					}
				})
			},
		},
		{
			Name:  "generate data",
			Input: []Struct1{{Name: "Alice", Level: 1}, {Name: "John", Level: 5}},
			Want:  []Struct2{{Data: "Alice lv1"}, {Data: "John lv5"}},
			TestFunc: func(input []Struct1) []Struct2 {
				return MapSlice(input, func(s Struct1) Struct2 {
					return Struct2{
						Data: fmt.Sprintf("%s lv%d", s.Name, s.Level),
					}
				})
			},
		},
	}
	runTestCases(t, casesInt2Int)
	runTestCases(t, casesInt2String)
	runTestCases(t, casesStruct2Struct)
}

func TestReduceSlice(t *testing.T) {
	casesInt := []TestCase[[]int, string]{
		{
			Name:  "sum",
			Input: []int{1, 5, 0, -5, 8, 7},
			Want:  "16",
			TestFunc: func(input []int) string {
				return ReduceSlice(input, "0", func(prev string, current int) string {
					iPrev, _ := strconv.Atoi(prev)
					return strconv.Itoa(iPrev + current)
				})
			},
		},
	}
	runTestCases(t, casesInt)
}
