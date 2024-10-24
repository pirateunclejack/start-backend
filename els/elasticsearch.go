package els

import (
	"github.com/elastic/go-elasticsearch/v8"
)

func NewElasticsearchClient()(*elasticsearch.TypedClient, error)  {
    cfg := elasticsearch.Config{
        Addresses: []string{elasticsearch_url},
    }

    return elasticsearch.NewTypedClient(cfg)
}
