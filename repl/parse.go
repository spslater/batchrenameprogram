package repl

import (
	"errors"
	"fmt"
	"math"
	"strings"
)

func flagName(raw string) string {
	if len(raw) > 1 && raw[:2] == "--" {
		return raw[2:]
	} else if len(raw) == 2 && raw[:1] == "-" {
		return raw[1:]
	} else {
		return raw
	}
}

func expandArgs(args []string) []string {
	var newargs []string
	for i := 0; i < len(args); i++ {
		var arg string = args[i]
		if len(arg) > 2 && arg[:1] == "-" && arg[1:2] != "-" {
			for _, na := range strings.Split(arg, "")[1:] {
				var temp string = fmt.Sprintf("-%s", na)
				newargs = append(newargs, temp)
			}
		} else {
			newargs = append(newargs, args[i])
		}
	}
	return newargs
}

func gobble(args []string, arg Args) (int, []string, error) {
	var nargs Nargs = arg.GetNargs()
	var got int
	var gather []string
	var err error

	eat := func(num int) {
		if num == -1 {
			num = math.MaxInt
		}
		for i := 0; i < num && i < len(args); i++ {
			var arg string = args[i]
			if arg[0:1] == "-" {
				break
			}
			if arg[0:1] == "/" {
				arg = arg[1:]
				args[i] = arg
			}
			gather = append(gather, args[i])
			got++
		}
	}

	if !arg.IsFlag() && nargs.Rep == 0 && nargs.Num == 0 {
		eat(1)
	} else if nargs.Rep == 0 {
		eat(nargs.Num)
	} else if nargs.Rep == '?' {
		if len(args) > 1 {
			gather = append(gather, args[1])
			got++
		}
	} else if nargs.Rep == '+' {
		if len(args) == 0 {
			err = errors.New("missing arguments")
		} else {
			eat(-1)
		}
	} else if nargs.Rep == '*' {
		eat(-1)
	}
	return got, gather, err
}

func splitArgs(args []Args) (map[string]Args, []Args) {
	var flags = make(map[string]Args)
	var positionals []Args

	for _, arg := range args {
		if arg.IsFlag() {
			flags[arg.GetName()] = arg
			for _, alias := range arg.GetAlias() {
				flags[alias] = arg
			}
		} else {
			positionals = append(positionals, arg)
		}
	}
	return flags, positionals
}

func (repl *Repl) Parse(raw []string) error {
	var args []string = expandArgs(raw)

	flags, positionals := splitArgs(repl.Args)

	var usedFlags []Args
	var usedPos []Args
	var nonflags []string
	var unknown []string

	for i := 0; i < len(args); i++ {
		var arg string = args[i]
		if arg[:1] == "-" {
			var fg string = flagName(arg)
			cur, ok := flags[fg]
			if !ok {
				unknown = append(unknown, arg)
				continue
			}
			got, gather, err := gobble(args[i+1:], cur)
			if err != nil {
				return errors.New(fmt.Sprintf("error gathering args: %s\n", err))
			}
			i += got
			_, err = cur.SetValue(gather)
			if err != nil {
				return errors.New(fmt.Sprintf("error setting value: %s\n", err))
			} else {
				usedFlags = append(usedFlags, cur)
			}
		} else {
			nonflags = append(nonflags, arg)
		}
	}
	for _, pos := range positionals {
		if len(nonflags) == 0 && pos.IsRequired() {
			break
		} else if len(nonflags) == 0 {
			break
		} else {
			got, gather, err := gobble(nonflags, pos)
			if err != nil {
				return errors.New(fmt.Sprintf("error gathering args: %s\n", err))
			}
			_, err = pos.SetValue(gather)
			if err != nil {
				return errors.New(fmt.Sprintf("error setting value: %s\n", err))
			}
			usedPos = append(usedPos, pos)
			nonflags = nonflags[got:]
		}
	}
	unknown = append(unknown, nonflags...)
	return nil
}
