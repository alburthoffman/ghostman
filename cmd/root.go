package cmd

// import (
// 	// "github.com/newly12/ghostman/inventory"
// 	"github.com/newly12/ghostman/request"
// 	"github.com/newly12/ghostman/response"
// 	// "log"
// 	"sync"
// )

// const (
// 	maxConcurrent = 3000
// )

// // var complete = make(chan int, 1)

// var wg sync.WaitGroup

// func Execute(hostfile string, tmpl *request.Template) {
// 	// load env

// 	// arg parse

// 	// get hosts

// 	// generate requests
// 	requests := GenRequestsWithTemplate(hostfile, tmpl)

// 	// log.Println("creating channels")
// 	reqChan := make(chan request.Request, maxConcurrent)
// 	resChan := make(chan *request.BaseResponse, maxConcurrent)

// 	complete <- 1
// 	// log.Println("fulfill request channel")
// 	go func() {
// 		for _, request := range requests {
// 			reqChan <- request
// 		}
// 		// log.Println("closing reqChan")
// 		close(reqChan)
// 	}()

// 	go func() {
// 		for baseres := range resChan {
// 			finalRes := *(response.ParseResponse(baseres))
// 			log.Printf("%v %v", finalRes.StatusCode, finalRes.Body)
// 		}
// 		// log.Println("Done results")
// 		<-complete
// 	}()

// 	// log.Println("process request channel")
// 	for req := range reqChan {
// 		wg.Add(1)
// 		go func(req request.Request) {
// 			defer wg.Done()
// 			resChan <- request.Run(&req)
// 		}(req)
// 	}

// 	wg.Wait()
// 	// log.Println("waitgroup done")
// 	close(resChan)
// }

// func GenRequests(hostfile string, tmplfile string) []request.Request {
// 	hs := request.ParseHosts(hostfile)
// 	tmpl := request.ParseTemplate(tmplfile)
// 	return request.GenRequests(tmpl, hs)
// }

// func GenRequestsWithTemplate(hostfile string, tmpl *request.Template) []request.Request {
// 	hs := request.ParseHosts(hostfile)
// 	return request.GenRequests(tmpl, hs)
// }
