package repl

import (
	"fmt"
	"strconv"
	"strings"
)

func (n *Nargs) String() string {
	if n.Rep == 0 {
		return fmt.Sprintf("%v", n.Num)
	}
	return fmt.Sprintf("%c", n.Rep)
}

func (a *Arg) String() string {
	var out strings.Builder
	var alias []string
	for _, a := range a.Alias {
		alias = append(alias, *a)
	}
	out.WriteString(fmt.Sprintf("Name: %s, ", *a.Name))
	out.WriteString(fmt.Sprintf("Alias: %s, ", alias))
	out.WriteString(fmt.Sprintf("Flag: %t, ", a.Flag))
	out.WriteString(fmt.Sprintf("Nargs: %s, ", a.Nargs.String()))
	out.WriteString(fmt.Sprintf("IsSet: %t, ", a.IsSet))
	out.WriteString(fmt.Sprintf("Required: %t, ", a.Required))
	return out.String()
}

func (r *Repl) String() string {
	var out strings.Builder
	out.WriteString("{\n")
	out.WriteString(fmt.Sprintf("  Name: %s\n", *r.Name))
	out.WriteString(fmt.Sprintf("  Usage: %s\n", *r.Usage))
	out.WriteString(fmt.Sprintf("  Desc: %s\n", *r.Desc))
	out.WriteString("  Args: [\n")
	for _, arg := range r.Args {
		out.WriteString(fmt.Sprintf("    %s\n", arg.String()))
	}
	out.WriteString("  ]\n}")
	out.WriteString("}\n")
	return out.String()
}

func (a *IntArg) String() string {
	var out strings.Builder
	out.WriteString("{")
	out.WriteString(a.Arg.String())
	out.WriteString(fmt.Sprintf("Default: %v, ", a.Default))
	out.WriteString(fmt.Sprintf("Value: %v, ", a.Value))
	out.WriteString("}")
	return out.String()
}

func (a *IntSliceArg) String() string {
	var out strings.Builder
	out.WriteString("{")
	out.WriteString(a.Arg.String())
	var tmpDef []string
	for _, v := range a.Default {
		tmpDef = append(tmpDef, strconv.Itoa(v))
	}
	out.WriteString(fmt.Sprintf("Default: %v, ", tmpDef))
	var tmpVal []string
	for _, v := range a.Value {
		tmpVal = append(tmpVal, strconv.Itoa(v))
	}
	out.WriteString(fmt.Sprintf("Value: %s, ", tmpVal))
	out.WriteString("}")
	return out.String()
}

func (a *StringArg) String() string {
	var out strings.Builder
	out.WriteString("{")
	out.WriteString(a.Arg.String())
	out.WriteString(fmt.Sprintf("Default: %q, ", a.Default))
	out.WriteString(fmt.Sprintf("Value: %q, ", a.Value))
	out.WriteString("}")
	return out.String()
}

func (a *StringSliceArg) String() string {
	var out strings.Builder
	out.WriteString("{")
	out.WriteString(a.Arg.String())
	var tmpDef []string
	for _, v := range a.Default {
		tmpDef = append(tmpDef, fmt.Sprintf("%q", v))
	}
	out.WriteString(fmt.Sprintf("Default: %s, ", tmpDef))
	var tmpVal []string
	for _, v := range a.Value {
		tmpVal = append(tmpVal, fmt.Sprintf("%q", *v))
	}
	out.WriteString(fmt.Sprintf("Value: %s, ", tmpVal))
	out.WriteString("}")
	return out.String()
}

func (a *BoolArg) String() string {
	var out strings.Builder
	out.WriteString("{")
	out.WriteString(a.Arg.String())
	out.WriteString(fmt.Sprintf("Default: %t, ", a.Default))
	out.WriteString(fmt.Sprintf("Value: %t, ", a.Value))
	out.WriteString("}")
	return out.String()
}
