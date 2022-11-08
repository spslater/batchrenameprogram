package brp

import (
    "whatno.io/batchrename/repl"
)

func confirm() repl.Args {
    return repl.NewBoolArg("confirm").
        SetAliases("c").
        SetFlag(true).
        SetRequired(false).
        SetDefault(false).
        SetHelp("automatically confirm action")
}

func subCommands() []*repl.Repl {
    return []*repl.Repl{
        HelpParser(),
        SaveParser(),
        QuitParser(),
        WriteParser(),
        ListParser(),
        HistoryParser(),
        UndoParser(),
        ResetParser(),
        AutoParser(),
        ReplaceParser(),
        AppendParser(),
        PrependParser(),
        InsertParser(),
        CaseParser(),
        ExtParser(),
    }
}

func CliParser(args ...string) (*repl.Repl, error) {
    var err error
    // filenames, autos
    parser := repl.NewRepl("batchrenamer", repl.CliCmd).
        SetAliases("brp").
        SetUsage("brp [-h] [-V] [[-a FILE] ...] filename [filename ...]").
        SetDesc("rename batches of files at one time").
        AddArg(repl.NewStringSliceArg("filenames").
            SetNargs('+').
            SetHelp("file to load patterns from")).
        AddArg(repl.NewStringSliceArg("autos").
            SetAliases("a").
            SetFlag(true).
            SetExactNargs(1).
            SetHelp("automated files to run"))
    parser.Cmds = subCommands()
    parser.Idk = UnknownParser()
    if len(args) > 0 {
        err = parser.Parse(args)
    }
    return parser, err
}

func UnknownParser() *repl.Repl {
    return repl.NewRepl("idk", repl.UnknownCmd)
}

func HelpParser() *repl.Repl {
    // small, commands
    return repl.NewRepl("help", repl.HelpCmd).
        SetAliases("h", "?").
        SetUsage("help (h, ?) [-s] [commands ...]").
        SetDesc("display help message").
        AddArg(repl.NewBoolArg("small").
            SetAliases("s").
            SetFlag(true).
            SetHelp("display just the usage messages")).
        AddArg(repl.NewStringSliceArg("commands").
            SetNargs('*').
            SetHelp("commands to get specific info on"))
}

func SaveParser() *repl.Repl {
    // confirm
    return repl.NewRepl("save", repl.SaveCmd).
        SetAliases("s").
        SetUsage("save (s) [-c]").
        SetDesc("save files with current changes").
        AddArg(confirm())
}

func QuitParser() *repl.Repl {
    // confirm
    return repl.NewRepl("quit", repl.QuitCmd).
        SetAliases("q", "exit").
        SetUsage("quit (q, exit) [-c]").
        SetDesc("quit program, don't apply unsaved changes").
        AddArg(confirm())
}

func WriteParser() *repl.Repl {
    // confirm
    return repl.NewRepl("write", repl.WriteCmd).
        SetAliases("w").
        SetUsage("write (w) [-c]").
        SetDesc("write changes and quit program, same as save then quit").
        AddArg(confirm())
}

func ListParser() *repl.Repl {
    return repl.NewRepl("list", repl.ListCmd).
        SetAliases("ls", "l").
        SetUsage("list (ls, l)").
        SetDesc("lists current files being modified")
}

func HistoryParser() *repl.Repl {
    // peak
    return repl.NewRepl("history", repl.HistoryCmd).
        SetAliases("hist", "past").
        SetUsage("history (hist, past) [--peak]").
        SetDesc("print history of changes for all files").
        AddArg(repl.NewBoolArg("peak").
            SetAliases("p").
            SetFlag(true).
            SetHelp("just show single file history"))
}

func UndoParser() *repl.Repl {
    // number
    return repl.NewRepl("undo", repl.UndoCmd).
        SetAliases("u").
        SetUsage("undo (u) [number]").
        SetDesc("undo last change made").
        AddArg(repl.NewIntArg("number").
            SetDefault(1).
            SetExactNargs(1).
            SetHelp("number of changes to undo"))
}

func ResetParser() *repl.Repl {
    // confirm
    return repl.NewRepl("reset", repl.ResetCmd).
        SetAliases("over", "o").
        SetUsage("reset (over, o) [-c]").
        SetDesc("reset changes to original inputs, no undoing").
        AddArg(confirm())
}

func AutoParser() *repl.Repl {
    return repl.NewRepl("automate", repl.AutoCmd).
        SetAliases("auto", "a").
        SetUsage("automate (a, auto) [filenames ...]").
        SetDesc("automate commands in order to speed up repetative tasks").
        AddArg(repl.NewStringSliceArg("filenames").
            SetNargs('*').
            SetHelp("file names to run commands"))
}

func ReplaceParser() *repl.Repl {
    // find, replace
    return repl.NewRepl("replace", repl.ReplaceCmd).
        SetAliases("r", "re", "reg", "regex").
        SetUsage("replace (r, re, reg, regex) [find] [replace]").
        SetDesc("find and replace based on a regex").
        AddArg(repl.NewStringArg("find").
            SetHelp("pattern to find")).
        AddArg(repl.NewStringArg("replace").
            SetHelp("pattern to insert"))
}

func AppendParser() *repl.Repl {
    // filenames, padding, append, find
    return repl.NewRepl("append", repl.AppendCmd).
        SetAliases("ap").
        SetUsage("append (ap) [-f FILENAMES [FILENAMES ...]] [-p PADDING] [find] [append]").
        SetDesc("pattern and value to append to each file that matches, can be automated with a file").
        AddArg(repl.NewStringSliceArg("filenames").
            SetAliases("f").
            SetFlag(true).
            SetNargs('+').
            SetHelp("file to load patterns from")).
        AddArg(repl.NewStringArg("padding").
            SetAliases("p").
            SetFlag(true).
            SetExactNargs(1).
            SetHelp("string to insert between the end of the filename and the value being appended")).
        AddArg(repl.NewStringArg("append").
            SetHelp("value to append to filename")).
        AddArg(repl.NewStringArg("find").
            SetHelp("regex pattern to match against"))
}

func PrependParser() *repl.Repl {
    // filenames, padding, prepend, find
    return repl.NewRepl("prepend", repl.PrependCmd).
        SetAliases("p", "pre").
        SetUsage("prepend (p, pre) [-f FILENAMES [FILENAMES ...]] [-p PADDING] [find] [prepend]").
        SetDesc("tsv with pattern and value to prepend to each file that matches").
        AddArg(repl.NewStringSliceArg("filenames").
            SetAliases("f").
            SetFlag(true).
            SetNargs('+').
            SetHelp("file to load patterns from")).
        AddArg(repl.NewStringArg("padding").
            SetAliases("p").
            SetFlag(true).
            SetExactNargs(1).
            SetHelp("string to insert between the end of the filename and the value being prepended")).
        AddArg(repl.NewStringArg("prepend").
            SetHelp("value to append to filename")).
        AddArg(repl.NewStringArg("find").
            SetHelp("regex pattern to match against"))
}

func InsertParser() *repl.Repl {
    // insert, index
    return repl.NewRepl("insert", repl.InsertCmd).
        SetAliases("i", "in").
        SetUsage("insert (i, in) [-c] [value] [index]").
        SetDesc("insert string, positive from begining, negative from ending").
        AddArg(repl.NewStringArg("insert").
            SetHelp("value to insert")).
        AddArg(repl.NewIntArg("index").
            SetHelp("index (starting from 0) to insert at, negative numbers will insert counting from the end")).
        AddArg(confirm())
}

func CaseParser() *repl.Repl {
    // styles
    return repl.NewRepl("case", repl.CaseCmd).
        SetAliases("c").
        SetUsage("case (c) [styles ...]").
        SetDesc("change the case (title, upper, lower) of files").
        AddArg(repl.NewStringSliceArg("styles").
            SetNargs('*').
            SetHelp("type of case style (lower, upper, title, camel, kebab, ect) to switch to"))
}

func ExtParser() *repl.Repl {
    // ext, new, pattern
    return repl.NewRepl("extension", repl.ExtCmd).
        SetAliases("x", "ext").
        SetUsage("extension (x, ext) [-e] [new] [pattern]").
        SetDesc("change the extension on all files or files that match pattern").
        AddArg(repl.NewBoolArg("ext").
            SetAliases("e").
            SetFlag(true).
            SetHelp("match against the extensions instead of filename")).
        AddArg(repl.NewBoolArg("nopattern").
            SetAliases("n", "np").
            SetFlag(true).
            SetHelp("change for all files")).
        AddArg(repl.NewStringArg("new").
            SetHelp("change the extension to this")).
        AddArg(repl.NewStringArg("pattern").
            SetHelp("pattern to match against, "))
}
