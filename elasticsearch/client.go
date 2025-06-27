package elasticsearch

import (
	"fmt"
	"log"
	"os"

	"github.com/elastic/go-elasticsearch/v8"
)

func NewElasticClient() (*elasticsearch.Client, error) {
	cfg := elasticsearch.Config{
		Addresses: []string{
			os.Getenv("ELASTIC_URI"),
		},
	}
	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to create Elastic client: %w", err)
	}

	// Optional: ping the cluster to verify connection
	res, err := es.Info()
	if err != nil {
		return nil, fmt.Errorf("failed to ping Elasticsearch: %w", err)
	}
	defer res.Body.Close()

	log.Println("Connected to Elasticsearch:", res.Status())
	return es, nil
}
