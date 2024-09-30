package brp

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/google/shlex"

	"whatno.io/batchrename/repl"
)

var pad int

func init() {
	pad = 10
}

type Pend struct {
	side   string
	parser repl.Repl
	input  func(repl.Repl) (string, string)
}

func RunCmd(cmd repl.Repl, args []string, parser repl.Repl, files []*FileHistory) bool {
	var exit bool = false
	var err error = cmd.Parse(args)
	if err != nil {
		fmt.Println(err)
		cmd.Reset()
		return exit
	}
	switch cmd.Id {
	case repl.HelpCmd:
		DoHelp(cmd, parser)
	case repl.SaveCmd:
		DoSave(cmd, files)
	case repl.QuitCmd:
		if DoQuit(cmd) {
			exit = true
		}
	case repl.WriteCmd:
		if DoWrite(cmd, files) {
			exit = true
		}
	case repl.ListCmd:
		DoList(files)
	case repl.HistoryCmd:
		DoHistory(cmd, files)
	case repl.UndoCmd:
		DoUndo(cmd, files)
	case repl.ResetCmd:
		DoReset(cmd, files)
	case repl.AutoCmd:
		DoManualAuto(parser, cmd, files)
	case repl.ReplaceCmd:
		DoReplace(cmd, files)
	case repl.AppendCmd:
		DoAppend(cmd, files)
	case repl.PrependCmd:
		DoPrepend(cmd, files)
	case repl.InsertCmd:
		DoInsert(cmd, files)
	case repl.CaseCmd:
		DoCase(cmd, files)
	case repl.ExtCmd:
		DoExtension(cmd, files)
	default:
		fmt.Printf("Unknown or Invalid Command\n\tPlease try again or type `help` ro `?` to get help\n")
	}
	cmd.Reset()
	return exit
}

func DoHelp(r repl.Repl, parser repl.Repl) {
	var small bool = r.GetValue("small").(bool)
	var reqs []string = r.GetValue("commands").([]string)

	var cmds []repl.Repl
	var errs []string
	if len(reqs) == 0 {
		cmds = parser.Cmds
	} else {
		for _, req := range reqs {
			var cmd repl.Repl = parser.GetCmd(req)
			if cmd.Id == repl.UnknownCmd {
				errs = append(errs, req)
			} else {
				cmds = append(cmds, cmd)
			}
		}
	}
	if len(errs) > 0 {
		fmt.Printf("Unknown commands: %s\n", strings.Join(errs, ", "))
	}

	for _, cmd := range cmds {
		if small {
			var usg, desc string = cmd.GetUsage()
			fmt.Printf("  %s\n      %s\n", usg, desc)
		} else {
			var usg, desc string = cmd.GetUsage()
			fmt.Println(usg)
			fmt.Printf("  %s\n", desc)
			var pos, opt []repl.Args = cmd.GetArgs()
			if len(pos) > 0 {
				fmt.Println("  positional arguments:")
				for _, p := range pos {
					var pu, pd string = p.GetUsage()
					if len(pu) > pad {
						fmt.Printf("    %s\n    %-*s %s\n", pu, pad, " ", pd)
					} else {
						pu = fmt.Sprintf("%-*s", pad, pu)
						fmt.Printf("    %s %s\n", pu, pd)
					}
				}
			}
			if len(opt) > 0 {
				fmt.Println("  optional arguments:")
				for _, o := range opt {
					var ou, od string = o.GetUsage()
					if len(ou) > pad {
						fmt.Printf("    %s\n    %-*s %s\n", ou, pad, " ", od)
					} else {
						ou = fmt.Sprintf("%-*s", pad, ou)
						fmt.Printf("    %s %s\n", ou, od)
					}
				}
			}
			fmt.Println()
		}
	}
}

func DoSave(r repl.Repl, files []*FileHistory) {
	if GetConfirm(r, "Are you sure you want to save? ") {
		for _, file := range files {
			file.Save()
		}
	}
}

func DoQuit(r repl.Repl) bool {
	return GetConfirm(r, "Are you sure you want to quit? ")
}

func DoWrite(r repl.Repl, files []*FileHistory) bool {
	if GetConfirm(r, "Are you sure you want to save and quit? ") {
		for _, file := range files {
			file.Save()
		}
		return true
	}
	return false
}

func DoList(files []*FileHistory) {
	for _, file := range files {
		fmt.Println(file.Display())
	}
}

func DoHistory(r repl.Repl, files []*FileHistory) {
	var peak bool = r.GetValue("peak").(bool)
	if peak {
		fmt.Println(files[0].History())
		return
	}
	for _, file := range files {
		fmt.Println(file.History())
	}
}

func DoUndo(r repl.Repl, files []*FileHistory) {
	var raw *int = r.GetValue("number").(*int)
	var num int = 0
	if raw != nil {
		num = *raw
	}
	for i := 0; i < num; i++ {
		for _, file := range files {
			if !file.Undo() {
				return
			}
		}
	}
}

func DoReset(r repl.Repl, files []*FileHistory) {
	if GetConfirm(r, "Are you sure you want to reset? ") {
		for _, file := range files {
			file.Reset()
		}
	}
}

func readfile(filename string) []string {
	var lines []string
	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		fmt.Println(err)
		return []string{}
	}

	var scanner *bufio.Scanner = bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		return []string{}
	}
	return lines

}

func DoAutofiles(parser repl.Repl, autofiles []string, files []*FileHistory) {
	for _, autofile := range autofiles {
		var lines []string = readfile(autofile)
		for _, line := range lines {
			toks, _ := shlex.Split(line)
			if len(toks) == 0 {
				continue
			}
			var cmd repl.Repl = parser.GetCmd(strings.ToLower(toks[0]))
			toks = toks[1:]
			switch cmd.Id {
			case repl.HelpCmd, repl.SaveCmd, repl.QuitCmd, repl.WriteCmd:
				continue
			default:
				RunCmd(cmd, toks, parser, files)
			}
		}
	}
}

func DoManualAuto(parser repl.Repl, r repl.Repl, files []*FileHistory) {
	var filenames []string = GetAutofiles(r)
	DoAutofiles(parser, filenames, files)
}

func DoReplace(r repl.Repl, files []*FileHistory) {
	var find, repl string = GetReplace(r)
	for _, file := range files {
		file.Replace(find, repl)
	}
}

func filePend(filenames []string, files []*FileHistory, kind Pend) {
	for _, filename := range filenames {
		var lines []string = readfile(filename)
		for _, line := range lines {
			toks, _ := shlex.Split(line)
			if len(toks) == 0 {
				continue
			}
			kind.parser.Parse(toks)
			var pend, find string = kind.input(kind.parser)
			manualPend(pend, find, kind, files)
		}
	}
}

func manualPend(pend string, find string, kind Pend, files []*FileHistory) {
	re, err := regexp.Compile(find)
	if err != nil {
		fmt.Printf("regex err: %v\n", err)
		return
	}
	for _, file := range files {
		if file.MatchName(re) {
			file.Replace(kind.side, pend)
		} else {
			file.Noop()
		}
	}
}

func doPend(r repl.Repl, files []*FileHistory, kind Pend) {
	var filenames []string = r.GetValue("filenames").([]string)

	if len(filenames) > 0 {
		filePend(filenames, files, kind)
	} else {
		var pend, find string = kind.input(r)
		manualPend(pend, find, kind, files)
	}
}

func DoAppend(r repl.Repl, files []*FileHistory) {
	var append Pend = Pend{side: "$", parser: AppendParser(), input: GetAppend}
	doPend(r, files, append)
}

func DoPrepend(r repl.Repl, files []*FileHistory) {
	var prepend Pend = Pend{side: "^", parser: PrependParser(), input: GetPrepend}
	doPend(r, files, prepend)
}

func DoInsert(r repl.Repl, files []*FileHistory) {
	idx, ins := GetInsert(r, files[0].PeakName())
	if ins == "" {
		return
	}
	for _, file := range files {
		file.Insert(idx, ins)
	}
}

func DoCase(r repl.Repl, files []*FileHistory) {
	var req_cases []CaseId = GetCases(r)
	for _, file := range files {
		file.ChangeCase(req_cases)
	}
}

func DoExtension(r repl.Repl, files []*FileHistory) {
	var ext bool = r.GetValue("ext").(bool)
	var new, patt = GetExtension(r)
	for _, file := range files {
		file.ChangeExt(new, patt, ext)
	}
}
