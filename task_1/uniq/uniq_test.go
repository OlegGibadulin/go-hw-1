package uniq

import (
	"github.com/stretchr/testify/require"
	"testing"
)

type TestCase struct {
	name     string
	options  *Options
	src      []string
	expected []string
}

var needCountOption = &Options{
	NeedCount:       true,
	OnlyRepeated:    false,
	OnlyUnique:      false,
	SkipFieldsCount: 0,
	SkipCharsCount:  0,
	IgnoreCase:      false,
}

func TestCheckSuccess(t *testing.T) {
	onlyRepeatedOption := &Options{
		NeedCount:       false,
		OnlyRepeated:    true,
		OnlyUnique:      false,
		SkipFieldsCount: 0,
		SkipCharsCount:  0,
		IgnoreCase:      false,
	}

	onlyUniqueOption := &Options{
		NeedCount:       false,
		OnlyRepeated:    false,
		OnlyUnique:      true,
		SkipFieldsCount: 0,
		SkipCharsCount:  0,
		IgnoreCase:      false,
	}

	ignoreCaseOption := &Options{
		NeedCount:       false,
		OnlyRepeated:    false,
		OnlyUnique:      false,
		SkipFieldsCount: 0,
		SkipCharsCount:  0,
		IgnoreCase:      true,
	}

	skipFieldsOption := func(count int) *Options {
		return &Options{
			NeedCount:       false,
			OnlyRepeated:    false,
			OnlyUnique:      false,
			SkipFieldsCount: count,
			SkipCharsCount:  0,
			IgnoreCase:      true,
		}
	}

	skipCharsOption := func(count int) *Options {
		return &Options{
			NeedCount:       false,
			OnlyRepeated:    false,
			OnlyUnique:      false,
			SkipFieldsCount: 0,
			SkipCharsCount:  count,
			IgnoreCase:      true,
		}
	}

	skipOptions := func(fields int, chars int) *Options {
		return &Options{
			NeedCount:       false,
			OnlyRepeated:    false,
			OnlyUnique:      false,
			SkipFieldsCount: fields,
			SkipCharsCount:  chars,
			IgnoreCase:      true,
		}
	}

	defaultSrc := []string{
		"I love music.",
		"I love music.",
		"I love music.",
		"",
		"I love music of Kartik.",
		"I love music of Kartik.",
		"Thanks.",
		"",
		"Thanks.",
	}

	diffRegSrc := []string{
		"I LOVE MUSIC.",
		"I love music.",
		"I LoVe MuSiC.",
		"",
		"I love MuSIC of Kartik.",
		"I love music of kartik.",
		"Thanks.",
	}

	forSkipSrc := []string{
		"I love music.",
		"A love music.",
		"C love music.",
		"",
		"I love music of Kartik.",
		"We love music of Kartik.",
		"Thanks.",
	}

	cases := []TestCase{
		TestCase{
			name:    "Default options",
			options: new(Options),
			src:     defaultSrc,
			expected: []string{
				"I love music.",
				"",
				"I love music of Kartik.",
				"Thanks.",
				"",
				"Thanks.",
			},
		},
		TestCase{
			name:    "NeedCount option",
			options: needCountOption,
			src:     defaultSrc,
			expected: []string{
				"3 I love music.",
				"1",
				"2 I love music of Kartik.",
				"1 Thanks.",
				"1",
				"1 Thanks.",
			},
		},
		TestCase{
			name:    "OnlyRepeated option",
			options: onlyRepeatedOption,
			src:     defaultSrc,
			expected: []string{
				"I love music.",
				"I love music of Kartik.",
			},
		},
		TestCase{
			name:    "OnlyUnique option",
			options: onlyUniqueOption,
			src:     defaultSrc,
			expected: []string{
				"",
				"Thanks.",
				"",
				"Thanks.",
			},
		},
		TestCase{
			name:    "IgnoreCase option",
			options: ignoreCaseOption,
			src:     diffRegSrc,
			expected: []string{
				"I LOVE MUSIC.",
				"",
				"I love MuSIC of Kartik.",
				"Thanks.",
			},
		},
		TestCase{
			name:    "SkipFields option",
			options: skipFieldsOption(1),
			src:     forSkipSrc,
			expected: []string{
				"I love music.",
				"",
				"I love music of Kartik.",
				"Thanks.",
			},
		},
		TestCase{
			name:    "SkipFields option with few fields",
			options: skipFieldsOption(3),
			src:     forSkipSrc,
			expected: []string{
				"I love music.",
				"I love music of Kartik.",
				"Thanks.",
			},
		},
		TestCase{
			name:    "SkipChars option",
			options: skipCharsOption(1),
			src:     forSkipSrc,
			expected: []string{
				"I love music.",
				"",
				"I love music of Kartik.",
				"We love music of Kartik.",
				"Thanks.",
			},
		},
		TestCase{
			name:    "SkipFields and SkipChars options",
			options: skipOptions(1, 20),
			src:     forSkipSrc,
			expected: []string{
				"I love music.",
				"I love music of Kartik.",
				"Thanks.",
			},
		},
	}

	for _, tc := range cases {
		actual := Execute(tc.options, tc.src)
		require.Equal(t, actual, tc.expected, tc.name)
	}
}

func TestCheckFail(t *testing.T) {
	cases := []TestCase{
		TestCase{
			name:     "Nil options",
			options:  nil,
			src:      []string{},
			expected: nil,
		},
		TestCase{
			name:     "Nil src",
			options:  new(Options),
			src:      nil,
			expected: nil,
		},
		TestCase{
			name:     "Empty input",
			options:  new(Options),
			src:      []string{},
			expected: nil,
		},
		TestCase{
			name:     "Empty input with some option",
			options:  needCountOption,
			src:      []string{},
			expected: nil,
		},
	}

	for _, tc := range cases {
		actual := Execute(tc.options, tc.src)
		require.Equal(t, actual, tc.expected, tc.name)
	}
}
