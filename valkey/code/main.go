package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
	"github.com/valkey-io/valkey-go"
)

var output = false

type Device struct {
	ID    string `json:"id"`
	Value string `json:"value"`
	UUID  string `json:"uuid"`
}

func main() {
	client, err := valkey.NewClient(valkey.ClientOption{InitAddress: []string{"valkey:6379"}})
	if err != nil {
		panic(err)
	}
	defer client.Close()

	var keyList [150_000_00]string
	ctx := context.Background()

	for i := range keyList {
		key, value := device_generator()

		keyList[i] = key

		if output {
			fmt.Printf("[%s] => (%s)\n", key, json_string(value))
		}

		err = client.Do(ctx, client.B().Set().Key(key).Value(json_string(value)).Nx().Build()).Error()
		if err != nil {
			fmt.Println(err)
			panic(err)
		}
	}

	if output {
		fmt.Print("\nResulting Map\n")

		for _, k := range keyList {
			v, err := client.Do(ctx, client.B().Get().Key(k).Build()).ToString()
			if err != nil {
				fmt.Println(err)
				panic(err)
			}
			fmt.Printf("[%s] => (%s)\n", k, v)
		}
	}

	print(len(keyList))
}

func json_string(device Device) string {
	out, err := json.Marshal(device)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	return string(out)
}

func device_generator() (string, Device) {
	return generate_key(), Device{
		ID:    generate_key(),
		Value: generate_key(),
		UUID:  generate_key(),
	}
}

func generate_key() string {
	key, err := uuid.NewRandom()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	return key.String()
}
