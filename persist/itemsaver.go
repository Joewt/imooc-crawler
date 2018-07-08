/**
	保存数据到elastic中
 */

package persist

import (
	"log"
	"gopkg.in/olivere/elastic.v5"
	"context"
	"joewt.com/joe/learngo/crawler/engine"
	"github.com/pkg/errors"
)


func ItemSaver() (chan engine.Item,error) {
	client, err := elastic.NewClient(elastic.SetSniff(false))

	if err != nil {
		return nil,err
	}
	out := make(chan engine.Item)
	go func() {
		itemCount := 0
		for {
			item := <- out
			log.Printf("Item saver: #%d: %v",itemCount,item)
			itemCount++
			err := save(client,item)
			if err != nil {
				log.Printf("Item saver error %v : %v\n",item,err)
			}
		}
	}()
	return out,nil
}


func save(client *elastic.Client,item engine.Item) error{


	if item.Type == "" {
		return errors.New("must supply type")
	}

	indexService := client.Index().
		Index("dating_profile").
		Type(item.Type).
		Id(item.Id).
		BodyJson(item)

	if item.Id != "" {
		indexService.Id(item.Id)
	}

	_, err := indexService.
		Do(context.Background())

	if err != nil {
		return err
	}

	return nil

}