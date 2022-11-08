package repl

import (
    "errors"
    "sort"
    "strings"
    "strconv"

    "whatno.io/batchrename/util"
)


type CmdId int64

const (
    UnknownCmd CmdId = iota
    CliCmd
    HelpCmd
    SaveCmd
    QuitCmd
    WriteCmd
    ListCmd
    HistoryCmd
    UndoCmd
    ResetCmd
    AutoCmd
    ReplaceCmd
    AppendCmd
    PrependCmd
    InsertCmd
    CaseCmd
    ExtCmd
)


type Nargs struct {
    Num int
    Rep rune
}

type Repl struct {
    Id CmdId
    Name *string
    Alias []*string
    Usage *string
    Desc  *string
    Args []Args
    Cmds []*Repl
    Idk *Repl
}

type Arg struct {
    Name *string
    Alias []*string
    Help *string
    Flag bool
    Nargs Nargs
    IsSet bool
    Required bool
}

type Args interface {
    String() string

    GetName() *string
    SetAliases(...string) Args
    GetAlias() []*string
    SetHelp(string) Args
    GetHelp() string
    GetUsage() (string, string)
    PartialUsage() string
    SetFlag(bool) Args
    IsFlag() bool
    SetExactNargs(int) Args
    SetNargs(rune) Args
    GetNargs() Nargs
    SetRequired(bool) Args
    IsRequired() bool

    SetValue([]*string) (Args, error)
    SetDefault(any) Args

    GetValue() any
    Reset() error
}

func NewRepl(name string, id CmdId) *Repl {
    return &Repl{Id: id, Name: util.Strptr(name)}
}

func (repl *Repl) SetAliases(aliases ...string) *Repl {
    repl.Alias = util.Strptrs(aliases...)
    return repl
}

func (repl *Repl) GetAliases() []*string {
    return repl.Alias
}

func (repl *Repl) SetUsage(usage string) *Repl {
    repl.Usage = util.Strptr(usage)
    return repl
}

type ByCmd []string
func (a ByCmd) Len() int           { return len(a) }
func (a ByCmd) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByCmd) Less(i, j int) bool { return len(a[i]) < len(a[j]) || (len(a[i]) == len(a[j]) && a[i] < a[j]) }


func (repl *Repl) GetUsage() (string, string) {
    var names []string
    for _, a := range repl.Alias {
        names = append(names,*a)
    }
    sort.Sort(ByCmd(names))
    alias := strings.Join(names, ", ")
    if len(alias) > 0 {
        alias = "(" + alias + ")"
    }

    for _, a := range repl.Alias {
        names = append(names,*a)
        
    }

    var args string
    for _, arg := range repl.Args {
        args = args + arg.PartialUsage() + " "
    }

    return *repl.Name + " " + alias + " " + args, *repl.Desc
}

func (repl *Repl) GetHelp() string {
    return *repl.Desc
}

func (repl *Repl) SetDesc(desc string) *Repl {
    repl.Desc = util.Strptr(desc)
    return repl
}

func (repl *Repl) GetDesc() *string {
    return repl.Desc
}

func (repl *Repl) AddArg(arg Args) *Repl {
    repl.Args = append(repl.Args, arg)
    return repl
}

func (repl *Repl) GetArgs() ([]Args, []Args) {
    var pos []Args
    var opt []Args
    for _, arg := range repl.Args {
        if arg.IsFlag() {
            opt = append(opt, arg)
        } else {
            pos = append(pos, arg)
        }
    }
    return pos, opt
}

func (repl *Repl) GetArg(name string) Args {
    for _, arg := range repl.Args {
        if name == *arg.GetName() {
            return arg
        }
        for _, alias := range arg.GetAlias() {
            if name == *alias {
                return arg
            }
        }
    }
    return nil
}

func (repl *Repl) GetValue(name string) any {
    // var err error
    var val any = nil
    var arg Args = repl.GetArg(name)
    if arg != nil {
        val = arg.GetValue()
    }
    return val //, err
}

func (repl *Repl) GetCmd(id string) *Repl {
    for _, cmd := range repl.Cmds {
        if id == *cmd.Name {
            return cmd
        }
        for _, alias := range cmd.Alias {
            if id == *alias {
                return cmd
            }
        }
    }
    return repl.Idk
}

func (repl *Repl) Reset() {
    for _, arg := range repl.Args {
        arg.Reset()
    }
}

func (repl *Repl) Matches(name string) bool {
    if *repl.Name == name {
        return true
    }
    for _, alias := range repl.Alias {
        if *alias == name {
            return true
        }
    }
    return false
}

/* Arg Commands */

func (a *Arg) GetName() *string {
    return a.Name
}

func (a *Arg) SetAliases(aliases ...string) Args {
    a.Alias = util.Strptrs(aliases...)
    return a
}

func (a *Arg) GetAlias() []*string {
    return a.Alias
}

func (a *Arg) SetHelp(help string) Args {
    a.Help = util.Strptr(help)
    return a
}

func (a *Arg) GetHelp() string {
    return *a.Help
}

func nargusage(meta string, nargs Nargs, req bool) string {
    var rep string
    if nargs.Rep == 0 && nargs.Num == 0 {
        rep = meta
        if !req {
            rep = "[" + rep + "]"
        }
    } else if nargs.Rep == 0 {
        if nargs.Num == 1 {
            rep = meta
        } else if nargs.Num == 2 || nargs.Num == 3 {
            var tmp []string
            for i := 0; i < nargs.Num; i++ {
                tmp = append(tmp, meta)
            }
            rep = strings.Join(tmp, " ")
            if !req {
                rep = "[" + rep + "]"
            }
        } else if nargs.Num != 0 {
            rep = meta + "{" + strconv.Itoa(nargs.Num) + "}"
        }
    } else if nargs.Rep == '?' {
        rep = "[" + meta + "]"
    } else if nargs.Rep == '+' {
        rep = meta + " [" + meta + " ...]"
    } else if nargs.Rep == '*' {
        rep = "[" + meta + " ...]"
    }

    return rep
}

func (a *Arg) PartialUsage() string {
    var usg string
    if a.IsFlag() {
        var fg string
        as := util.Derefstr(a.Alias)
        sort.Sort(ByCmd(as))
        if len(as) > 0 && len(as[0]) == 1 {
            fg = "-" + as[0]
        } else {
            fg = "--" + *a.Name
        }
        nargs := a.GetNargs()
        if nargs.Rep != 0 || nargs.Num != 0 {
            usg = nargusage(strings.ToUpper(*a.Name), nargs, a.IsRequired())
        }
        if a.IsRequired() && len(usg) > 0 {
            return fg + " " + usg
        } else if a.IsRequired() {
            return fg 
        } else if len(usg) > 0 {
            return "[" + fg + " " + usg + "]"
        } else {
            return "[" + fg + "]"
        }
    }
    return nargusage(*a.Name, a.GetNargs(), a.IsRequired())
}

func (a *Arg) GetUsage() (string, string) {
    if !a.IsFlag() {
        return *a.GetName(), *a.Help
    }

    var nargs Nargs = a.GetNargs()
    var hasargs bool = (nargs.Rep != 0 || nargs.Num != 0)
    var args string = nargusage(strings.ToUpper(*a.Name), nargs, a.IsRequired())

    var short []string
    var long []string

    var tmp string
    tmp = "--" + *a.GetName()
    if hasargs { tmp += " " + args }
    long = append(long, tmp)
    for _, k := range a.Alias {
        tmp = "-"+*k
        if hasargs { tmp += " " + args }
        if len(*k) == 1 {
            short = append(short, tmp)
        } else {
            tmp = "-" + tmp
            long = append(long, tmp)
        }
    }
    sort.Strings(short)
    sort.Strings(long)
    aliases := append(short, long...)
    sort.Sort(ByCmd(aliases))
    alias := strings.Join(aliases, ", ")

    return alias, *a.Help
}

func (a *Arg) SetFlag(f bool) Args {
    a.Flag = f
    return a
}

func (a *Arg) IsFlag() bool {
    return a.Flag
}

func (a *Arg) SetExactNargs(narg int) Args {
    if narg < 1 {
        return nil
    }
    a.Nargs = Nargs{Num: narg}
    return a
}

func (a *Arg) SetNargs(narg rune) Args {
    if narg != '?' && narg != '*' && narg != '+' {
        return nil
    }
    a.Nargs = Nargs{Rep: narg}
    return a
}

func (a *Arg) GetNargs() Nargs {
    return a.Nargs
}

func (a *Arg) SetRequired(req bool) Args {
    a.Required = req
    return a
}

func (a *Arg) IsRequired() bool {
    return a.Required
}

func (a *Arg) SetValue(gather []*string) (Args, error) {
    return nil, errors.New("please implament value")
}

func (a *Arg) SetDefault(_ any) Args {
    return a
}

func (a *Arg) GetValue() any {
    return nil
}

func (a *Arg) Reset() error {
    return nil
}
