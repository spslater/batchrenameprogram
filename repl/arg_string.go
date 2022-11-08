package repl

import (
    "errors"

    "whatno.io/batchrename/util"
)

type StringArg struct {
    Arg
    Default *string
    Value *string
}

func NewStringArg(name string) *StringArg {
    return &StringArg{Arg: Arg{Name: util.Strptr(name)}}
}

func (a *StringArg) SetRequired(req bool) Args { a.Arg.SetRequired(req); return a }
func (a *StringArg) IsRequired() bool { return a.Arg.IsRequired() }

func (a *StringArg) SetHelp(help string) Args { a.Arg.SetHelp(help); return a }
func (a *StringArg) GetHelp() string { return a.Arg.GetHelp() }
func (a *StringArg) GetUsage() (string, string) { return a.Arg.GetUsage() }
func (a *StringArg) PartialUsage() string { return a.Arg.PartialUsage() }

func (a *StringArg) SetFlag(f bool) Args { a.Arg.SetFlag(f); return a }
func (a *StringArg) IsFlag() bool { return a.Arg.IsFlag() }

func (a *StringArg) SetExactNargs(narg int) Args { a.Arg.SetExactNargs(narg); return a }
func (a *StringArg) SetNargs(narg rune) Args { a.Arg.SetNargs(narg); return a }
func (a *StringArg) GetNargs() Nargs { return a.Nargs }

func (a *StringArg) GetName() *string { return a.Arg.GetName() }

func (a *StringArg) SetAliases(aliases ...string) Args { a.Arg.SetAliases(aliases...); return a }
func (a *StringArg) GetAlias() []*string { return a.Arg.GetAlias() }

func (a *StringArg) SetValue(gather []*string) (Args, error) {
    if len(gather) != 1 {
        return nil, errors.New("too many values passed")
    }
    a.Value = gather[0]
    a.IsSet = true
    return a, nil
}

func (a *StringArg) SetDefault(def any) Args {
    a.Default = def.(*string)
    return a
}

func (a *StringArg) GetValue() any {
    if a.IsSet {
        return a.Value
    }
    return a.Default
}

func (a *StringArg) Reset() error {
    a.IsSet = false
    a.Value = a.Default
    return nil
}
