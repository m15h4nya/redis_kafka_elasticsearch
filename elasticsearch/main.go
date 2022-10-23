package main

import (
	"context"
	"fmt"
	elasticlib "github.com/olivere/elastic/v7"
	"github.com/segmentio/kafka-go"
	"time"
)

func main() {
	c := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{"kafka:9092"}, // kafka_broker: from docker-compose network
		Topic:     "test_topic",
		GroupID:   "AnyGroup",
		Partition: 0,
		MinBytes:  0,    // 10KB
		MaxBytes:  10e6, // 10MB
	})

	m, err := c.FetchMessage(context.Background())
	if err != nil {
		fmt.Println(err)
	}
	err = c.CommitMessages(context.Background(), m)
	if err != nil {
		fmt.Println("ERROR: ", err)
	}
	fmt.Printf("message at offset %d: %s = %s\n", m.Offset, string(m.Key), string(m.Value))

	conn, err := elasticlib.NewClient(
		elasticlib.SetSniff(false),
		elasticlib.SetURL("http://elasticsearch:9200"),
		elasticlib.SetHealthcheckInterval(10*time.Second),
	)

	if err != nil {
		fmt.Println(err)
	}

	mapping := `{
		"mappings": {
			"settings": {
				"number_of_replicas": 0
			}
			"properties": {
				"name": {"type": "keyword"},
				"surname": {"type": "keyword"},
				"second_name": {"type": "keyword"},
				"married": {"type": "boolean"},
				"age": {"type": "integer"}
			}
		}
	}`

	if _, err = conn.CreateIndex("test_index").Body(mapping).Do(context.Background()); err != nil {
		fmt.Println(err)
	}

	query := `{"query":
					{"bool":
						{"must":
							[
								{"term": {"name": "k0tletka"}},
								{"term": {"surname": "po"}},
								{"term": {"age": 21}}
							] }}}`

	res, err := conn.Search("test_index").Source(query).Do(context.Background())
	if err != nil {
		fmt.Println(err)
	}

	if len(res.Hits.Hits) == 0 {
		if _, err := conn.Index().Index("test_index").BodyString(string(m.Value)).Do(context.Background()); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Успешно")
		}
	} else {
		fmt.Println("Запись уже существует")
	}

}
