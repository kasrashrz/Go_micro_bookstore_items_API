package elasticsearch

import (
	"github.com/federicoleon/bookstore_utils-go/logger"
	"os"
	"time"
)

import (
	"github.com/olivere/elastic"
)

const (
	envEsHosts = "ELASTIC_HOSTS"
)

var (
	Client esClientInterface = &esClient{}
)

type esClientInterface interface {
	setClient(*elastic.Client)
	//Index(string, string, interface{}) (*elastic.IndexResponse, error)
	//Get(string, string, string) (*elastic.GetResult, error)
	//Search(string, elastic.Query) (*elastic.SearchResult, error)
}

type esClient struct {
	client *elastic.Client
}

func Init() {
	log := logger.GetLogger()

	client, err := elastic.NewClient(
		elastic.SetURL(os.Getenv(envEsHosts)),
		elastic.SetHealthcheckInterval(10*time.Second),
		elastic.SetErrorLog(log),
		elastic.SetInfoLog(log),
	)
	if err != nil {
		panic(err)
	}
	Client.setClient(client)

}

func (c *esClient) setClient(client *elastic.Client) {
	c.client = client
}
