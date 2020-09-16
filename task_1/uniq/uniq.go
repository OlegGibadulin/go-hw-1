package uniq

import ()

type Options struct {
	NeedCount       bool
	OnlyRepeated    bool
	OnlyUnique      bool
	SkipFieldsCount uint
	SkipCharsCount  uint
	IgnoreCase      bool
}

func Execute(src []string, options *Options) ([]string, error) {
	return nil, nil
}
