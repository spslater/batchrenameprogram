package brp

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/google/shlex"
	"github.com/chzyer/readline"

	"whatno.io/batchrename/repl"
)

type ResId int64

const (
	Invalid ResId = iota
	Confirm
	Deny
	Back
)

var response = map[string]ResId{
	"yes":  Confirm,
	"y":    Confirm,
	"no":   Deny,
	"n":    Deny,
	"back": Back,
	"b":    Back,
	"quit": Back,
	"q":    Back,
}

var reader *readline.Instance

func init() {
	var err error
	reader, err = readline.New("?? ")
	if err != nil {
		panic(err)
	}
}

func readrepl(prefix string) string {
	for {
		reader.SetPrompt(prefix)
		input, err := reader.Readline()
		if err != nil {
			fmt.Printf("error: %v\n", err)
			continue
		}
		return input
	}
}

func getstring(val *string, msg string) string {
	if val == nil {
		return readrepl(msg)
	}
	return *val
}

func GetResponse(r repl.Repl, msg string) ResId {
	if r.GetValue("confirm").(bool) {
		return Confirm
	}

	if msg == "" {
		msg = "Yes or No? "
	}
	var res string
	var ans ResId
	var valid bool
	for {
		res = strings.ToLower(readrepl(msg))
		ans, valid = response[res]
		if !valid {
			fmt.Println("Please enter a valid response.")
			continue
		}
		return ans
	}
}

func GetConfirm(r repl.Repl, msg string) bool {
	return Confirm == GetResponse(r, msg)
}

func Command(parser repl.Repl) (repl.Repl, []string) {
	for {
		var input = readrepl(">> ")
		toks, _ := shlex.Split(input)
		if len(toks) == 0 {
			continue
		}
		var command repl.Repl = parser.GetCmd(strings.ToLower(toks[0]))
		return command, toks[1:]
	}
}

func GetReplace(args repl.Repl) (string, string) {
	var raw_find *string = args.GetValue("find").(*string)
	var raw_repl *string = args.GetValue("replace").(*string)

	var find string = getstring(raw_find, "Find: ")
	var repl string = getstring(raw_repl, "Repl: ")

	return find, repl
}

func GetExtension(args repl.Repl) (string, string) {
	var nopattern bool = args.GetValue("nopattern").(bool)

	var raw_new *string = args.GetValue("new").(*string)
	var new string = strings.TrimSpace(getstring(raw_new, "New Ext: "))

	var pattern string = ""
	if !nopattern {
		var raw_pattern *string = args.GetValue("pattern").(*string)
		pattern = getstring(raw_pattern, "Match Pattern (Leave blank for no pattern): ")
	}

	return new, pattern
}

func validateCases(args []string) ([]string, []CaseId) {
	var errs []string
	var good []CaseId
	for _, arg := range args {
		style, valid := Cases[arg]
		if !valid {
			errs = append(errs, arg)
		} else {
			good = append(good, style)
		}
	}
	return errs, good
}

func GetCases(args repl.Repl) []CaseId {
	var styles []string = args.GetValue("styles").([]string)
	var errs []string
	var good []CaseId
	if len(styles) > 0 {
		errs, good = validateCases(styles)
		if len(errs) > 0 {
			fmt.Printf("Invalid cases requested: %v\nPlease enter valid cases\n", errs)
		}
	}
	if len(styles) == 0 || len(errs) > 0 {
		for {
			styles = strings.Split(readrepl("Styles? "), " ")
			errs, good = validateCases(styles)
			if len(errs) > 0 {
				fmt.Printf("Invalid cases requested: %v\n", errs)
			} else {
				break
			}
		}
	}

	return good
}

func getPend(args repl.Repl, pend string) (string, string, string) {
	var raw_pend *string = args.GetValue(pend).(*string)
	var raw_find *string = args.GetValue("find").(*string)
	var padding *string = args.GetValue("padding").(*string)

	var pad string = " "
	if padding != nil {
		pad = *padding
	}

	var add string = getstring(raw_pend, strings.Title(pend)+": ")
	var find string = getstring(raw_find, "Find: ")
	return pad, add, find
}

func GetAppend(args repl.Repl) (string, string) {
	var pad, pend, find string = getPend(args, "append")
	return pad + pend, find
}

func GetPrepend(args repl.Repl) (string, string) {
	var pad, pend, find string = getPend(args, "prepend")
	return pend + pad, find
}

func GetInsert(args repl.Repl, test string) (int, string) {
	var raw_ins *string = args.GetValue("insert").(*string)
	var raw_idx *int = args.GetValue("index").(*int)
	var confirm bool = args.GetValue("confirm").(bool)

	var insert string = getstring(raw_ins, "Insert: ")
	for {
		var num int
		var err error
		if raw_idx == nil {
			num, err = strconv.Atoi(readrepl("Index: "))
			if err != nil {
				fmt.Println("please enter positive or negative integer")
				continue
			}
		} else {
			num = *raw_idx
		}

		var tmp string = insertinto(test, num, insert)
		if confirm {
			return num, insert
		}
		fmt.Println("Example:", tmp)
		switch GetResponse(args, "Right index? ") {
		case Confirm:
			return num, insert
		case Deny:
			raw_idx = args.GetValue("index").(*int)
		case Back:
			return 0, ""
		}
	}
}

func GetAutofiles(args repl.Repl) []string {
	var rawnames []string = args.GetValue("filenames").([]string)

	if len(rawnames) > 0 {
		return rawnames
	}
	toks, _ := shlex.Split(readrepl("Filenames: "))
	return toks
}
