// Copyright 2018 Yuanting Liang. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build go1.1

package main

import (
	"fmt"
	"io"
	"os"
	"runtime"
	"strings"
	"sync"
	"text/template"
	"unicode"
	"unicode/utf8"
)

// Test that go1.1 tag above is included in builds. main.go refers to this definition.
const go11tag = true

const version = "0.1"

var supportedDrivers = map[string]string{
	"mysql":    "github.com/go-sql-driver/mysql",
	"mymysql":  "github.com/ziutek/mymysql/godrv",
	"postgres": "github.com/lib/pq",
	"mssql":    "github.com/denisenkom/go-mssqldb",
}

// Commands lists the available commands and help topics.
// The order here is the order in which they are printed by 'gopm help'.
var commands = []*Command{
	CmdReverse,
	CmdShell,
	CmdDump,
	CmdDriver,
	CmdSource,
}

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	// Check length of arguments.
	args := os.Args[1:]
	if len(args) < 1 {
		usage()
		return
	}

	// Show help documentation.
	if args[0] == "help" {
		help(args[1:])
		return
	}

	// Check commands and run.
	for _, comm := range commands {
		if comm.Name() == args[0] && comm.Run != nil {
			comm.Run(comm, args[1:])
			exit()
			return
		}
	}

	fmt.Fprintf(os.Stderr, "gorm: unknown subcommand %q\nRun 'gorm help' for usage.\n", args[0])
	setExitStatus(2)
	exit()
}

var exitStatus = 0
var exitMu sync.Mutex

func setExitStatus(n int) {
	exitMu.Lock()
	if exitStatus < n {
		exitStatus = n
	}
	exitMu.Unlock()
}

var usageTemplate = `gorm is a database tool based gorm package. 

Version:

    ` + version + `

Usage:

    gorm command [arguments]

The commands are:
{{range .}}{{if .Runnable}}
    {{.Name | printf "%-11s"}} {{.Short}}{{end}}{{end}}

Use "gorm help [command]" for more information about a command.

Additional help topics:
{{range .}}{{if not .Runnable}}
    {{.Name | printf "%-11s"}} {{.Short}}{{end}}{{end}}

Use "gorm help [topic]" for more information about that topic.

`

var helpTemplate = `{{if .Runnable}}usage: gorm {{.UsageLine}}

{{end}}{{.Long | trim}}
`

// tmpl executes the given template text on data, writing the result to w.
func tmpl(w io.Writer, text string, data interface{}) {
	t := template.New("top")
	t.Funcs(template.FuncMap{"trim": strings.TrimSpace, "capitalize": capitalize})
	template.Must(t.Parse(text))
	if err := t.Execute(w, data); err != nil {
		panic(err)
	}
}

func capitalize(s string) string {
	if s == "" {
		return s
	}
	r, n := utf8.DecodeRuneInString(s)
	return string(unicode.ToTitle(r)) + s[n:]
}

func printUsage(w io.Writer) {
	tmpl(w, usageTemplate, commands)
}

func usage() {
	printUsage(os.Stderr)
	os.Exit(2)
}

// help implements the 'help' command.
func help(args []string) {
	if len(args) == 0 {
		printUsage(os.Stdout)
		// not exit 2: succeeded at 'gopm help'.
		return
	}
	if len(args) != 1 {
		fmt.Fprintf(os.Stderr, "usage: gorm help command\n\nToo many arguments given.\n")
		os.Exit(2) // failed at 'gopm help'
	}

	arg := args[0]

	for _, cmd := range commands {
		if cmd.Name() == arg {
			tmpl(os.Stdout, helpTemplate, cmd)
			// not exit 2: succeeded at 'gopm help cmd'.
			return
		}
	}

	fmt.Fprintf(os.Stderr, "Unknown help topic %#q.  Run 'gorm help'.\n", arg)
	os.Exit(2) // failed at 'gopm help cmd'
}

var atexitFuncs []func()

func atexit(f func()) {
	atexitFuncs = append(atexitFuncs, f)
}

func exit() {
	for _, f := range atexitFuncs {
		f()
	}
	os.Exit(exitStatus)
}
