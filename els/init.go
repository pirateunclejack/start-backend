package els

import (
	"context"
	"log"
	"start-backend/helper"
)

var elasticsearch_url string
var Elasticsearch_index string


func init() {
    vip := helper.GetConfig()
    elasticsearch_url = vip.GetString("ELASTICSEARCH_URL")
    Elasticsearch_index = vip.GetString("ELASTICSEARCH_INDEX")

    client, err := NewElasticsearchClient()
    if err != nil {
        log.Fatalln("failed to create elasticsearch client: ", err)
    }

    exists, err  := client.Indices.Exists(Elasticsearch_index).Do(context.TODO())
    if err != nil {
        log.Fatalln("failed to check index exists: ", err)
    }
    if !exists {
        res, err := client.Indices.Create(
            Elasticsearch_index,
        ).Do(context.TODO())
        if err != nil {
            log.Fatalln("failed to create elasticsearch index: ", err)
        }
        log.Println("create elasticsearch index success: ", res.Index)
    }
}
