package uniq

import (
	"strconv"
	"strings"
)

type Options struct {
	NeedCount       bool
	OnlyRepeated    bool
	OnlyUnique      bool
	SkipFieldsCount int
	SkipCharsCount  int
	IgnoreCase      bool
}

func Execute(src []string, options *Options) ([]string, error) {
	knownStrings := make(map[string]int)
	var uniqStrings []string

	for _, s := range src {
		template := getStringTemplate(options, s)
		if _, ok := knownStrings[template]; ok {
			knownStrings[template]++
		} else {
			uniqStrings = append(uniqStrings, s)
			knownStrings[template] = 1
		}
	}

	uniqStrings = checkOptions(options, knownStrings, uniqStrings)

	return uniqStrings, nil
}

func getStringTemplate(options *Options, s string) string {
	template := s

	if options.IgnoreCase {
		template = strings.ToLower(template)
	}

	if options.SkipFieldsCount != 0 {
		words := strings.Split(template, " ")
		if len(words) <= options.SkipFieldsCount {
			template = ""
		} else {
			template = strings.Join(words[options.SkipFieldsCount:], " ")
		}
	}

	if options.SkipCharsCount != 0 {
		if len(template) <= options.SkipCharsCount {
			template = ""
		} else {
			template = template[options.SkipCharsCount:]
		}
	}
	return template

}

func getStringCount(options *Options, knownStrings map[string]int, s string) int {
	template := getStringTemplate(options, s)
	if options.IgnoreCase {
		template = strings.ToLower(s)
	}
	return knownStrings[template]
}

func checkOptions(options *Options, knownStrings map[string]int, uniqStrings []string) []string {
	uniqStrings = checkNeedCount(options, knownStrings, uniqStrings)
	uniqStrings = checkOnlyRepeated(options, knownStrings, uniqStrings)
	uniqStrings = checkOnlyUnique(options, knownStrings, uniqStrings)
	return uniqStrings
}

func checkNeedCount(options *Options, knownStrings map[string]int, uniqStrings []string) []string {
	var res []string
	if options.NeedCount {
		for _, s := range uniqStrings {
			count := getStringCount(options, knownStrings, s)
			resString := strconv.Itoa(count) + " " + s
			res = append(res, resString)
		}
		return res
	}
	return uniqStrings
}

func checkOnlyRepeated(options *Options, knownStrings map[string]int, uniqStrings []string) []string {
	var res []string
	if options.OnlyRepeated {
		for _, s := range uniqStrings {
			count := getStringCount(options, knownStrings, s)
			if count > 1 {
				res = append(res, s)
			}
		}
		return res
	}
	return uniqStrings
}

func checkOnlyUnique(options *Options, knownStrings map[string]int, uniqStrings []string) []string {
	var res []string
	if options.OnlyUnique {
		for _, s := range uniqStrings {
			count := getStringCount(options, knownStrings, s)
			if count == 1 {
				res = append(res, s)
			}
		}
		return res
	}
	return uniqStrings
}
