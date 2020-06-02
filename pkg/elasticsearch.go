package elasticsearch

import (
	"context"
	"fmt"
	"github.com/olivere/elastic"
)

type Es struct {
	client *elastic.Client
}

func NewESClient(esUrl string) (IElasticSearchClient, error) {
	instance, err := elastic.NewSimpleClient(elastic.SetURL(esUrl))
	if err != nil {
		return nil, err
	}
	return &Es{
		client: instance,
	}, nil
}

type IElasticSearchClient interface {
	GetClient() *elastic.Client
	IndexExists(index string) (bool, error)
	CreateIndex(index string, mapping string) (*elastic.IndicesCreateResult, error)
	DeleteIndex(index string) (*elastic.IndicesDeleteResponse, error)
	Insert(index string, byteSlice []byte) (*elastic.IndexResponse, error)
}

func (es *Es) GetClient() *elastic.Client {
	return es.client
}

func (es *Es) IndexExists(index string) (bool, error) {
	return es.GetClient().IndexExists(index).Do(context.Background())
}

func (es *Es) CreateIndex(index string, mapping string) (*elastic.IndicesCreateResult, error) {
	if mapping != "" {
		return es.GetClient().CreateIndex(index).Body(mapping).Do(context.Background())
	} else {
		return es.GetClient().CreateIndex(index).Do(context.Background())
	}
}

func (es *Es) DeleteIndex(index string) (*elastic.IndicesDeleteResponse, error) {
	return es.GetClient().DeleteIndex(index).Do(context.Background())
}

func (es *Es) Insert(index string, byteSlice []byte) (*elastic.IndexResponse, error) {
	data := string(byteSlice)
	fmt.Println(data)
	return es.GetClient().Index().Index(index).Type("entity").Id("1").BodyJson(data).Refresh("true").Do(context.Background())
}