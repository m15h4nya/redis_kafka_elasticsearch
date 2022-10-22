package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v9"
)

type person struct {
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	SecondName string `json:"second_name"`
	Married    bool   `json:"married"`
	Age        int    `json:"age"`
}

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})

	misha := person{
		Name:       "k0tletka",
		Surname:    "po",
		SecondName: "kievski",
		Married:    false,
		Age:        21,
	}

	value, err := json.Marshal(&misha)
	if err != nil {
		fmt.Println(err)
	}

	if err := rdb.Set(context.Background(), "first_key", value, 0); err != nil {
		fmt.Println(err)
	}

}
