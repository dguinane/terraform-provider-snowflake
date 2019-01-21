package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/dguinane/terraform-provider-snowflake/snowflake"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{ProviderFunc: snowflake.Provider})
}
