/**
	爬虫引擎
 */

package engine

type ConcurrentEngine struct {
	Scheduler Scheduler
	WorkerCount int
	ItemChan chan Item
}

type Scheduler interface {
	ReadyNotifier
	Submit(Request)
	WorkerChan() chan Request
	Run()
}

type ReadyNotifier interface {
	WorkerReady(chan Request)
}


func (e *ConcurrentEngine) Run(seeds ...Request){

	out := make(chan ParserResult)
	e.Scheduler.Run()

	for i := 0; i < e.WorkerCount; i++ {
		createWorker(e.Scheduler.WorkerChan(),out,e.Scheduler)
	}

	for _, r := range seeds {
		e.Scheduler.Submit(r)
	}
	for {
		result := <- out
		for _, item := range result.Items {
			go func() {e.ItemChan <- item}()
		}

		for _, request := range result.Requests {
			if isDuplicate(request.Url) {
				continue
			}
			e.Scheduler.Submit(request)
		}
	}
}


func createWorker(in chan Request,out chan ParserResult,ready ReadyNotifier){
	go func() {
		for {
			ready.WorkerReady(in)
			request := <- in
			result, err := Worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}

var visitedUrls = make(map[string]bool)

func isDuplicate(url string) bool {
	if visitedUrls[url] {
		return true
	}
	visitedUrls[url] = true
	return false
}