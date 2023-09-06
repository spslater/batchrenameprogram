package repl

import (
	"errors"

	"whatno.io/batchrename/util"
)

type StringSliceArg struct {
	Arg
	Default []*string
	Value   []*string
}

func NewStringSliceArg(name string) *StringSliceArg {
	return &StringSliceArg{Arg: Arg{Name: util.Strptr(name)}}
}

func (a *StringSliceArg) SetRequired(req bool) Args { a.Arg.SetRequired(req); return a }
func (a *StringSliceArg) IsRequired() bool          { return a.Arg.IsRequired() }

func (a *StringSliceArg) SetHelp(help string) Args   { a.Arg.SetHelp(help); return a }
func (a *StringSliceArg) GetHelp() string            { return a.Arg.GetHelp() }
func (a *StringSliceArg) GetUsage() (string, string) { return a.Arg.GetUsage() }
func (a *StringSliceArg) PartialUsage() string       { return a.Arg.PartialUsage() }

func (a *StringSliceArg) SetFlag(f bool) Args { a.Arg.SetFlag(f); return a }
func (a *StringSliceArg) IsFlag() bool        { return a.Arg.IsFlag() }

func (a *StringSliceArg) SetExactNargs(narg int) Args { a.Arg.SetExactNargs(narg); return a }
func (a *StringSliceArg) SetNargs(narg rune) Args     { a.Arg.SetNargs(narg); return a }
func (a *StringSliceArg) GetNargs() Nargs             { return a.Nargs }

func (a *StringSliceArg) GetName() *string { return a.Arg.GetName() }

func (a *StringSliceArg) SetAliases(aliases ...string) Args { a.Arg.SetAliases(aliases...); return a }
func (a *StringSliceArg) GetAlias() []*string               { return a.Arg.GetAlias() }

func (a *StringSliceArg) SetValue(gather []*string) (Args, error) {
	if a.Nargs.Rep == 0 && (len(a.Value)+len(gather) > a.Nargs.Num) {
		return a, errors.New("too many values passed in")
	}
	a.Value = append(a.Value, gather...)
	a.IsSet = true
	return a, nil
}

func (a *StringSliceArg) SetDefault(def any) Args {
	a.Default = def.([]*string)
	return a
}

func (a *StringSliceArg) GetValue() any {
	if a.IsSet {
		return a.Value
	}
	return a.Default
}

func (a *StringSliceArg) Reset() error {
	a.IsSet = false
	a.Value = a.Default
	return nil
}
