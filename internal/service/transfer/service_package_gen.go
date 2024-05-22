// Code generated by internal/generate/servicepackages/main.go; DO NOT EDIT.

package transfer

import (
	"context"

	aws_sdkv2 "github.com/aws/aws-sdk-go-v2/aws"
	transfer_sdkv2 "github.com/aws/aws-sdk-go-v2/service/transfer"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/types"
	"github.com/hashicorp/terraform-provider-aws/names"
)

type servicePackage struct{}

func (p *servicePackage) FrameworkDataSources(ctx context.Context) []*types.ServicePackageFrameworkDataSource {
	return []*types.ServicePackageFrameworkDataSource{}
}

func (p *servicePackage) FrameworkResources(ctx context.Context) []*types.ServicePackageFrameworkResource {
	return []*types.ServicePackageFrameworkResource{}
}

func (p *servicePackage) SDKDataSources(ctx context.Context) []*types.ServicePackageSDKDataSource {
	return []*types.ServicePackageSDKDataSource{
		{
			Factory:  DataSourceServer,
			TypeName: "aws_transfer_server",
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  resourceAccess,
			TypeName: "aws_transfer_access",
			Name:     "Access",
		},
		{
			Factory:  resourceAgreement,
			TypeName: "aws_transfer_agreement",
			Name:     "Agreement",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceCertificate,
			TypeName: "aws_transfer_certificate",
			Name:     "Certificate",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceConnector,
			TypeName: "aws_transfer_connector",
			Name:     "Connector",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceProfile,
			TypeName: "aws_transfer_profile",
			Name:     "Profile",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceServer,
			TypeName: "aws_transfer_server",
			Name:     "Server",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  ResourceSSHKey,
			TypeName: "aws_transfer_ssh_key",
		},
		{
			Factory:  resourceTag,
			TypeName: "aws_transfer_tag",
			Name:     "Transfer Resource Tag",
		},
		{
			Factory:  ResourceUser,
			TypeName: "aws_transfer_user",
			Name:     "User",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  ResourceWorkflow,
			TypeName: "aws_transfer_workflow",
			Name:     "Workflow",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.Transfer
}

// NewClient returns a new AWS SDK for Go v2 client for this service package's AWS API.
func (p *servicePackage) NewClient(ctx context.Context, config map[string]any) (*transfer_sdkv2.Client, error) {
	cfg := *(config["aws_sdkv2_config"].(*aws_sdkv2.Config))

	return transfer_sdkv2.NewFromConfig(cfg, func(o *transfer_sdkv2.Options) {
		if endpoint := config[names.AttrEndpoint].(string); endpoint != "" {
			o.BaseEndpoint = aws_sdkv2.String(endpoint)
		}
	}), nil
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}
