package main

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// Set by GoReleaser for tagged releases.
var version = "dev"

type doctestProvider struct{}

func (p *doctestProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "doctest"
	resp.Version = version
}

func (p *doctestProvider) Schema(_ context.Context, _ provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "An empty provider for testing Terraform Registry documentation publishing.",
	}
}

func (p *doctestProvider) Configure(context.Context, provider.ConfigureRequest, *provider.ConfigureResponse) {
}

func (p *doctestProvider) Resources(context.Context) []func() resource.Resource {
	return nil
}

func (p *doctestProvider) DataSources(context.Context) []func() datasource.DataSource {
	return nil
}

func main() {
	err := providerserver.Serve(context.Background(), func() provider.Provider {
		return &doctestProvider{}
	}, providerserver.ServeOpts{
		Address: "registry.terraform.io/gerrytan/doctest",
	})
	if err != nil {
		log.Fatal(err)
	}
}
