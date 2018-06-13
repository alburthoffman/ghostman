package cmd

import (
	"github.com/newly12/ghostman/template"
	"gopkg.in/alecthomas/kingpin.v2"
	"io/ioutil"
	"strings"
)

const (
	tmplSuffix = ".yml"
)

type SubComands map[string]*template.Template
type FullCommands map[string]SubComands

var (
	// args vars
	app = kingpin.New("ghostman", "ghostman")

	// runtime
	debug          = app.Flag("debug", "debug mode").Bool()
	_maxConcurrent = app.Flag("maxconcurrent", "max concurrent").Short('g').Int()

	// list
	list = app.Command("list", "list available commands")

	// run
	runCmd   = app.Command("run", "run template based commands")
	cmd      = runCmd.Arg("cmdSet", "command set").String()
	subCmd   = runCmd.Arg("subCmd", "sub command").String()
	hostfile = app.Flag("hostfile", "host file").Short('f').String()
	dryrun   = app.Flag("dry", "dryrun").Bool()
)

func ListCommands(templateDir string) *FullCommands {
	cmds := make(FullCommands)
	var cmd string
	var subCmd string
	var tmplPath string

	cmdDirs, err := ioutil.ReadDir(templateDir)
	if err != nil {
		log.Fatal("no templates dir!")
	}

	for _, _cmd := range cmdDirs {
		cmd = _cmd.Name()
		if !_cmd.IsDir() {
			continue
		}

		subCmds, err := ioutil.ReadDir(templateDir + "/" + cmd)
		if err != nil {
			log.Fatal("invalid templates dir: ", cmd)
		}

		cmds[cmd] = SubComands{}
		for _, _subCmd := range subCmds {
			subCmd = strings.TrimSuffix(_subCmd.Name(), tmplSuffix)
			tmplPath = templateDir + "/" + cmd + "/" + _subCmd.Name()
			cmds[cmd][subCmd] = template.ParseTemplate(tmplPath)
		}
	}
	return &cmds
}
