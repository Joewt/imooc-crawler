package main

import (
	"joewt.com/joe/learngo/crawler/engine"
	"joewt.com/joe/learngo/crawler/zhenai/parser"
	"joewt.com/joe/learngo/crawler/Scheduler"
	"joewt.com/joe/learngo/crawler/persist"
)

func main() {
	//使用并发版并开启了100个协程处理
	e := engine.ConcurrentEngine{
		Scheduler:   &Scheduler.QueuedScheduler{},
		WorkerCount: 100,
		ItemChan: persist.ItemSaver(),
	}

	e.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParserCityList,
	})

}
