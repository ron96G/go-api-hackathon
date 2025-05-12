package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"os/signal"
	"syscall"

	"github.com/ron96G/go-api-hackathon/internal/api"
)

var (
	rootCtx = context.Background()

	id string
	op string
)

func init() {
	flag.StringVar(&id, "id", "", "ID of the pet")
	flag.StringVar(&op, "op", "", "Operation to perform (add, delete, list, get)")
}

func main() {
	flag.Parse()

	ctx, cancel := signal.NotifyContext(rootCtx, syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	client, err := api.NewClientWithResponses("http://localhost:8080")
	if err != nil {
		panic(err)
	}

	if op == "add" {
		res, err := client.AddPetWithResponse(ctx, api.AddPetJSONRequestBody{
			Name: "Dog",
			Category: api.Category{
				Name: "Dogs",
			},
			Tags: []api.Tag{
				{
					Name: "something",
				},
			},
		})
		if err != nil {
			handleErr(err, res.Body)
		}

		prettyPrint(res.JSON200)
	}

	if op == "list" {
		res, err := client.ListPetsWithResponse(ctx)
		if err != nil {
			handleErr(err, res.Body)
		}

		prettyPrint(res.JSON200)
	}
	if op == "delete" {
		if id == "" {
			fmt.Println("ID is required for get operation")
			return
		}
		res, err := client.DeletePetWithResponse(ctx, id)
		if err != nil {
			handleErr(err, res.Body)
		}

		if res.StatusCode() == 200 {
			fmt.Println("Pet deleted successfully")
		} else {
			fmt.Printf("Failed to delete pet: %s\n", res.Body)
		}
	}

	if op == "get" {
		if id == "" {
			fmt.Println("ID is required for get operation")
			return
		}
		res, err := client.GetPetByIdWithResponse(ctx, id)
		if err != nil {
			handleErr(err, res.Body)
		}

		if res.StatusCode() == 200 {
			prettyPrint(res.JSON200)
		} else {
			fmt.Printf("Failed to get pet: %s\n", res.Body)
		}
	}
}

func handleErr(err error, body []byte) {
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	if body != nil {
		var errResponse api.ErrorResponse
		if err := json.Unmarshal(body, &errResponse); err != nil {
			fmt.Printf("Failed to unmarshal error response: %v\n", err)
		} else {
			fmt.Printf("Error response: %+v\n", errResponse)
		}
	}
}

func prettyPrint(v any) {
	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		fmt.Printf("Failed to marshal: %v\n", err)
		return
	}
	fmt.Println(string(b))
}
