package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/Khan/genqlient/graphql"
	"github.com/jccr/steampipe-plugin-linear/gql"
)

type authedTransport struct {
	key     string
	wrapped http.RoundTripper
}

func (t *authedTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Set("Authorization", "Bearer "+t.key)
	return t.wrapped.RoundTrip(req)
}

func main() {
	var err error
	defer func() {
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}()

	key := os.Getenv("LINEAR_TOKEN")
	if key == "" {
		err = fmt.Errorf("must set LINEAR_TOKEN=<access token>")
		return
	}

	httpClient := http.Client{
		Transport: &authedTransport{
			key:     key,
			wrapped: http.DefaultTransport,
		},
	}
	graphqlClient := graphql.NewClient("https://api.linear.app/graphql", &httpClient)

	var listIssueResponse *gql.ListIssueResponse
	listIssueResponse, err = gql.ListIssue(context.Background(), graphqlClient, 10, "")
	if err != nil {
		return
	}
	fmt.Printf("listIssueResponse: %v\n", listIssueResponse.Issues.Nodes[0])
}

//go:generate go run github.com/Khan/genqlient ../genqlient.yaml
