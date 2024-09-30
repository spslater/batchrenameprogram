package repl

import (
	"errors"
	"strconv"
)

type IntArg struct {
	Arg
	Default *int
	Value   *int
}

func NewIntArg(name string) *IntArg {
	return &IntArg{Arg: Arg{Name: name}}
}

func (a *IntArg) SetRequired(req bool) Args { a.Arg.SetRequired(req); return a }
func (a IntArg)  IsRequired() bool          { return a.Arg.IsRequired() }

func (a *IntArg) SetHelp(help string) Args   { a.Arg.SetHelp(help); return a }
func (a IntArg)  GetHelp() string            { return a.Arg.GetHelp() }
func (a IntArg)  GetUsage() (string, string) { return a.Arg.GetUsage() }
func (a IntArg)  PartialUsage() string       { return a.Arg.PartialUsage() }

func (a *IntArg) SetFlag(f bool) Args { a.Arg.SetFlag(f); return a }
func (a IntArg)  IsFlag() bool        { return a.Arg.IsFlag() }

func (a *IntArg) SetExactNargs(narg int) Args { a.Arg.SetExactNargs(narg); return a }
func (a *IntArg) SetNargs(narg rune) Args     { a.Arg.SetNargs(narg); return a }
func (a IntArg)  GetNargs() Nargs             { return a.Nargs }

func (a IntArg)  GetName() string { return a.Arg.GetName() }

func (a *IntArg) SetAliases(aliases ...string) Args { a.Arg.SetAliases(aliases...); return a }
func (a IntArg)  GetAlias() []string                { return a.Arg.GetAlias() }

func (a *IntArg) SetValue(gather []string) (Args, error) {
	if len(gather) != 1 {
		return nil, errors.New("too many values passed")
	}
	val, err := strconv.Atoi(gather[0])
	if err != nil {
		return nil, err
	}
	a.Value = &val
	a.IsSet = true
	return a, nil
}

func (a *IntArg) SetDefault(def any) Args {
	var temp int = def.(int)
	a.Default = &temp
	return a
}

func (a IntArg) GetValue() any {
	if a.IsSet {
		return a.Value
	}
	return a.Default
}

func (a *IntArg) Reset() error {
	a.IsSet = false
	a.Value = a.Default
	return nil
}
