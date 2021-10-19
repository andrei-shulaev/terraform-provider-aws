// Code generated by "internal/generate/listpages/main.go -ListOps=DescribeIpGroups -Export"; DO NOT EDIT.

package workspaces

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/workspaces"
)

func DescribeIPGroupsPages(conn *workspaces.WorkSpaces, input *workspaces.DescribeIpGroupsInput, fn func(*workspaces.DescribeIpGroupsOutput, bool) bool) error {
	return DescribeIPGroupsPagesWithContext(context.Background(), conn, input, fn)
}

func DescribeIPGroupsPagesWithContext(ctx context.Context, conn *workspaces.WorkSpaces, input *workspaces.DescribeIpGroupsInput, fn func(*workspaces.DescribeIpGroupsOutput, bool) bool) error {
	for {
		output, err := conn.DescribeIpGroupsWithContext(ctx, input)
		if err != nil {
			return err
		}

		lastPage := aws.StringValue(output.NextToken) == ""
		if !fn(output, lastPage) || lastPage {
			break
		}

		input.NextToken = output.NextToken
	}
	return nil
}
