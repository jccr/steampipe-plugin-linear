package gql

import (
	"context"

	"github.com/Khan/genqlient/graphql"
)

type ListIssueResponse = listIssueResponse

func ListIssue(
	ctx context.Context,
	client graphql.Client,
	first int,
	after string,
) (*ListIssueResponse, error) {
	return listIssue(ctx, client, first, after)
}
