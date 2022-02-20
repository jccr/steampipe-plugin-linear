package main

import (
	"github.com/jccr/steampipe-plugin-linear/linear"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{PluginFunc: linear.Plugin})
}
