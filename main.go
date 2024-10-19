package main

import (
	"fmt"
	"os"

	"whatno.io/batchrename/brp"
	"whatno.io/batchrename/repl"
	"whatno.io/batchrename/util"
)

func main() {
	var files []*brp.FileHistory
	parser, err := brp.CliParser(os.Args[1:]...)
	if err != nil {
		fmt.Println(err)
		return
	}

	if parser.GetValue("help").(bool) {
		fmt.Println("usage:", parser.Usage)
		return
	}

	if parser.GetValue("version").(bool) {
		fmt.Println(util.Version)
		return
	}

	var filenames []string = parser.GetValue("filenames").([]string)
	if len(filenames) == 0 {
		fmt.Println("need files to rename... silly")
		return
	}
	for _, filename := range filenames {
		files = append(files, brp.NewFileHistory(filename))
	}

	var autofiles []string = parser.GetValue("autos").([]string)
	brp.DoAutofiles(parser, autofiles, files)

	var exit bool = false
	for {
		cmd, args := brp.Command(parser)
		exit = brp.RunCmd(cmd, args, parser, files)
		if exit {
			break
		}
		switch cmd.Id {
		case repl.CaseCmd, repl.AppendCmd, repl.PrependCmd, repl.UndoCmd, repl.ReplaceCmd, repl.ExtCmd, repl.InsertCmd, repl.AutoCmd, repl.DeleteCmd:
			brp.DoList(files)
		}
	}
}
