package cmd

import (
	"github.com/newly12/ghostman/config"
	"github.com/newly12/ghostman/execute"
	"github.com/newly12/ghostman/inventory"
	"github.com/newly12/ghostman/logging"
	"github.com/newly12/ghostman/output"
	"github.com/newly12/ghostman/request"
	"github.com/newly12/ghostman/response"
	"gopkg.in/alecthomas/kingpin.v2"
	"os"
	"sync"
)

var defaults *config.Defaults

var (
	maxConcurrent = 2000
	log           = logging.GetLogger()
	complete      = make(chan int)
)

func init() {
	// load env
	defaults = config.GetDefaults()
}

func Entrypoint() {
	// arg parse, along with template parse
	cmds := *ListCommands(defaults.TmplDir)
	switch kingpin.MustParse(app.Parse(os.Args[1:])) {
	case list.FullCommand():
		println("available commands:")

	case runCmd.FullCommand():
		run(cmds)
	}
}

func run(cmds FullCommands) {
	var wg sync.WaitGroup

	// get hosts
	log.Debug("load host file")
	hs := inventory.LoadHostFile(*hostfile)

	tmpl := cmds[*cmd][*subCmd]

	// create channels
	log.Debug("creating channels")
	reqChan := make(chan *request.Request, defaults.MaxConcur)
	resChan := make(chan *response.Response, defaults.MaxConcur)

	// generate requests
	log.Debug("generate requests")
	requests := request.GetRequests(hs, tmpl)

	// execute
	go func() {
		for _, request := range requests {
			reqChan <- &request
		}
		close(reqChan)
	}()

	go func() {
		for {
			res, more := <-resChan
			if more {
				output.Println(res)
			} else {
				complete <- 1
				return
			}
		}
	}()

	log.Debug("process request channel")
	for req := range reqChan {
		wg.Add(1)
		go func(req *request.Request) {
			defer wg.Done()
			resChan <- execute.Execute(req)
		}(req)
	}

	wg.Wait()
	close(resChan)
	<-complete
}
