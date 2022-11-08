package repl

import (
    "errors"
    "strconv"

    "whatno.io/batchrename/util"
)

type BoolArg struct {
    Arg
    Default bool
    Value bool
}

func NewBoolArg(name string) Args {
    return &BoolArg{Arg: Arg{Name: util.Strptr(name)}}
}

func (a *BoolArg) SetRequired(req bool) Args { a.Arg.SetRequired(req); return a }
func (a *BoolArg) IsRequired() bool { return a.Arg.IsRequired() }

func (a *BoolArg) SetHelp(help string) Args { a.Arg.SetHelp(help); return a }
func (a *BoolArg) GetHelp() string { return a.Arg.GetHelp() }
func (a *BoolArg) GetUsage() (string, string) { return a.Arg.GetUsage() }
func (a *BoolArg) PartialUsage() string { return a.Arg.PartialUsage() }

func (a *BoolArg) SetFlag(f bool) Args { a.Arg.SetFlag(f); return a }
func (a *BoolArg) IsFlag() bool { return a.Arg.IsFlag() }

func (a *BoolArg) SetExactNargs(narg int) Args { a.Arg.SetExactNargs(narg); return a }
func (a *BoolArg) SetNargs(narg rune) Args { a.Arg.SetNargs(narg); return a }
func (a *BoolArg) GetNargs() Nargs { return a.Nargs }

func (a *BoolArg) GetName() *string { return a.Arg.GetName() }

func (a *BoolArg) SetAliases(aliases ...string) Args { a.Arg.SetAliases(aliases...); return a }
func (a *BoolArg) GetAlias() []*string { return a.Arg.GetAlias() }

func (a *BoolArg) SetValue(gather []*string) (Args, error) {
    if len(gather) > 1 {
        return nil, errors.New("too many values passed")
    }

    var val bool
    var err error

    if len(gather) == 0 {
        val = true
    } else {
        val, err = strconv.ParseBool(*gather[0])
        if err != nil {
            return nil, err
        }
    }

    a.Value = val
    a.IsSet = true
    return a, nil
}

func (a *BoolArg) SetDefault(def any) Args {
    a.Default = def.(bool)
    return a
}

func (a *BoolArg) GetValue() any {
    if a.IsSet {
        return a.Value
    }
    return a.Default
}

func (a *BoolArg) Reset() error {
    a.IsSet = false
    a.Value = a.Default
    return nil
}
