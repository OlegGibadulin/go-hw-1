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

func Execute(options *Options, src []string) []string {
	if options == nil || src == nil {
		return nil
	}
	uniqStrings, stringsCount := getUniqStrings(options, src)
	uniqStrings = checkOptions(options, uniqStrings, stringsCount)
	return uniqStrings
}

func getUniqStrings(options *Options, src []string) ([]string, map[string]int) {
	if len(src) == 0 {
		return nil, nil
	}
	var uniqStrings []string
	stringsCount := make(map[string]int)

	prevString := src[0]
	prevTemplate := getStringTemplate(options, src[0])
	uniqCount := 1
	for _, s := range src[1:] {
		curTemplate := getStringTemplate(options, s)
		if prevTemplate == curTemplate {
			uniqCount++
		} else {
			uniqStrings = append(uniqStrings, prevString)
			stringsCount[prevString] = uniqCount
			prevString = s
			prevTemplate = curTemplate
			uniqCount = 1
		}
	}
	uniqStrings = append(uniqStrings, prevString)
	stringsCount[prevString] = uniqCount
	return uniqStrings, stringsCount
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

func checkOptions(options *Options, uniqStrings []string, stringsCount map[string]int) []string {
	uniqStrings = checkNeedCount(options, uniqStrings, stringsCount)
	uniqStrings = checkOnlyRepeated(options, uniqStrings, stringsCount)
	uniqStrings = checkOnlyUnique(options, uniqStrings, stringsCount)
	return uniqStrings
}

func checkNeedCount(options *Options, uniqStrings []string, stringsCount map[string]int) []string {
	var res []string
	if options.NeedCount {
		for _, s := range uniqStrings {
			count := stringsCount[s]
			resString := strconv.Itoa(count)
			if s != "" {
				resString += " "
			}
			resString += s
			res = append(res, resString)
		}
		return res
	}
	return uniqStrings
}

func checkOnlyRepeated(options *Options, uniqStrings []string, stringsCount map[string]int) []string {
	var res []string
	if options.OnlyRepeated {
		for _, s := range uniqStrings {
			count := stringsCount[s]
			if count > 1 {
				res = append(res, s)
			}
		}
		return res
	}
	return uniqStrings
}

func checkOnlyUnique(options *Options, uniqStrings []string, stringsCount map[string]int) []string {
	var res []string
	if options.OnlyUnique {
		for _, s := range uniqStrings {
			count := stringsCount[s]
			if count == 1 {
				res = append(res, s)
			}
		}
		return res
	}
	return uniqStrings
}
