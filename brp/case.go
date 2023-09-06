package brp

import (
	"regexp"
	"strings"
	"unicode"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type CaseId int64

const (
	UnknownCase CaseId = iota
	Upper
	Lower
	Title
	Kebab
	Snake
	Dekebab
	Desnake
	Squash
	Trim
	Camel
	Pascal
	Unsquash
	Default
)

var CaseById map[CaseId][2]string
var Cases map[string]CaseId
var toLower func(string) string
var toUpper func(string) string
var toTitle func(string) string

func init() {
	CaseById = map[CaseId][2]string{
		Upper:    [2]string{"upper", "u"},
		Lower:    [2]string{"lower", "l"},
		Title:    [2]string{"title", "t"},
		Kebab:    [2]string{"kebab", "k"},
		Snake:    [2]string{"snake", "s"},
		Dekebab:  [2]string{"dekebab", "dk"},
		Desnake:  [2]string{"desnake", "ds"},
		Squash:   [2]string{"squash", "sq"},
		Trim:     [2]string{"trim", "tr"},
		Camel:    [2]string{"camel", "c"},
		Pascal:   [2]string{"pascal", "p"},
		Unsquash: [2]string{"unsquash", "us"},
		Default:  [2]string{"default", "d"},
	}

	Cases = make(map[string]CaseId)
	for cid, opts := range CaseById {
		for _, opt := range opts {
			Cases[opt] = cid
		}
	}

	toLower = cases.Lower(language.English).String
	toUpper = cases.Upper(language.English).String
	toTitle = cases.Title(language.English).String
}

func ChangeCase(val string, cid CaseId) string {
	var newval string
	switch cid {
	case Upper:
		newval = toUpper(val)
	case Lower:
		newval = toLower(val)
	case Title:
		newval = toTitle(val)
	case Kebab:
		newval = toKebab(val)
	case Dekebab:
		newval = toDekebab(val)
	case Snake:
		newval = toSnake(val)
	case Desnake:
		newval = toDesnake(val)
	case Squash:
		newval = toSquash(val)
	case Unsquash:
		newval = toUnsquash(val)
	case Trim:
		newval = toTrim(val)
	case Camel:
		newval = toCamel(val)
	case Pascal:
		newval = toPascal(val)
	default:
		newval = val
	}
	return newval
}

func toKebab(val string) string {
	return strings.ReplaceAll(val, " ", "-")
}

func toDekebab(val string) string {
	return strings.ReplaceAll(val, "-", " ")
}

func toSnake(val string) string {
	return strings.ReplaceAll(val, " ", "_")
}

func toDesnake(val string) string {
	return strings.ReplaceAll(val, "_", " ")
}

func toSquash(val string) string {
	reg := regexp.MustCompile(`\s+`)
	return reg.ReplaceAllString(val, "")
}

func toTrim(val string) string {
	reg := regexp.MustCompile(`\s+`)
	return reg.ReplaceAllString(val, " ")
}

func toCamel(val string) string {
	toks := strings.Split(val, " ")
	toks[0] = toLower(toks[0])
	for i, v := range toks[1:] {
		toks[i+1] = toTitle(v)
	}
	return strings.Join(toks, "")
}

func toPascal(val string) string {
	toks := strings.Split(val, " ")
	for i, v := range toks {
		toks[i] = toTitle(v)
	}
	return strings.Join(toks, "")
}

func toUnsquash(val string) string {
	var ns strings.Builder
	for _, v := range val {
		if unicode.IsUpper(v) {
			ns.WriteString(" ")
		}
		ns.WriteRune(v)
	}
	reg := regexp.MustCompile(`(^\s+|\s+$)`)
	return reg.ReplaceAllString(ns.String(), "")
}
