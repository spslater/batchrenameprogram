package repl

import (
	"errors"
	"strconv"

	"whatno.io/batchrename/util"
)

type IntSliceArg struct {
	Arg
	Default []int
	Value   []int
}

func NewIntSliceArg(name string) *IntSliceArg {
	return &IntSliceArg{Arg: Arg{Name: util.Strptr(name)}}
}

func (a *IntSliceArg) SetRequired(req bool) Args { a.Arg.SetRequired(req); return a }
func (a *IntSliceArg) IsRequired() bool          { return a.Arg.IsRequired() }

func (a *IntSliceArg) SetHelp(help string) Args   { a.Arg.SetHelp(help); return a }
func (a *IntSliceArg) GetHelp() string            { return a.Arg.GetHelp() }
func (a *IntSliceArg) GetUsage() (string, string) { return a.Arg.GetUsage() }
func (a *IntSliceArg) PartialUsage() string       { return a.Arg.PartialUsage() }

func (a *IntSliceArg) SetFlag(f bool) Args { a.Arg.SetFlag(f); return a }
func (a *IntSliceArg) IsFlag() bool        { return a.Arg.IsFlag() }

func (a *IntSliceArg) SetExactNargs(narg int) Args { a.Arg.SetExactNargs(narg); return a }
func (a *IntSliceArg) SetNargs(narg rune) Args     { a.Arg.SetNargs(narg); return a }
func (a *IntSliceArg) GetNargs() Nargs             { return a.Nargs }

func (a *IntSliceArg) GetName() *string { return a.Arg.GetName() }

func (a *IntSliceArg) SetAliases(aliases ...string) Args { a.Arg.SetAliases(aliases...); return a }
func (a *IntSliceArg) GetAlias() []*string               { return a.Arg.GetAlias() }

func (a *IntSliceArg) SetValue(gather []*string) (Args, error) {
	if a.Nargs.Rep == 0 && (len(a.Value)+len(gather) > a.Nargs.Num) {
		return a, errors.New("too many values passed in")
	}
	var vals []int
	for _, v := range gather {
		val, err := strconv.Atoi(*v)
		if err != nil {
			return a, err
		}
		vals = append(vals, val)
	}
	a.Value = vals
	a.IsSet = true
	return a, nil
}

func (a *IntSliceArg) SetDefault(def any) Args {
	a.Default = def.([]int)
	return a
}

func (a *IntSliceArg) GetValue() any {
	if a.IsSet {
		return a.Value
	}
	return a.Default
}

func (a *IntSliceArg) Reset() error {
	a.IsSet = false
	a.Value = a.Default
	return nil
}
