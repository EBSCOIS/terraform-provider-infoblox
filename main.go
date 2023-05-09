package main

import (
	"github.com/ebscois/platform.infrastructure.terraform-provider-infoblox/infoblox"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: infoblox.Provider})
}
