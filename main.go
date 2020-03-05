package main

import (
    "github.com/hashicorp/terraform/plugin"
    "github.com/D4rk4/terraform-provider-croc/croc"
    )

    func main() {
     plugin.Serve(&plugin.ServeOpts{
	ProviderFunc: croc.Provider,
})
}
