package linear

import (
	"context"
	"fmt"
	"log"

	"github.com/carlosflorencio/linear-cli/config"
	"github.com/machinebox/graphql"
)

const endpoint = "https://api.linear.app/graphql"

var (
	client = graphql.NewClient(endpoint)
)

type UserResponse struct {
	Viewer struct {
		ID    string `json:"id,omitempty"`
		Name  string `json:"name,omitempty"`
		Email string `json:"email,omitempty"`
	} `json:"viewer,omitempty"`
}

func User() {
	req := graphql.NewRequest(`
        query Me {
	  viewer {
	    id
	    name
	    email
	  }
	}`)

	var res UserResponse
	fetch(req, &res)

	fmt.Println("ID:", res.Viewer.ID)
	fmt.Printf("%+v\n", res)
}

func fetch(req *graphql.Request, output interface{}) {
	// set header fields
	req.Header.Set("Authorization", config.ApiKey())

	// define a Context for the request
	ctx := context.Background()

	if err := client.Run(ctx, req, &output); err != nil {
		log.Fatal(err)
	}
}
