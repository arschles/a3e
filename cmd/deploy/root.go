package deploy

import (
	"github.com/a3e/a3e/cmd"
	"github.com/a3e/a3e/pkg/aci"
	"github.com/a3e/a3e/pkg/log/human"
	"github.com/a3e/ae3/pkg/cfg"
	"github.com/spf13/cobra"
)

// Root returns the command for the root of the deploy command tree
func Root() *cobra.Command {
	ret := cmd.Skeleton("deploy", "Deploy your containers")
	ret.RunE = func(cmd *cobra.Command, args []string) error {
		cl := aci.NewClient("TODO")
		return human.Check(
			cl.Deploy(
				cfg.Cfg.ID.SubscriptionID,
				cfg.Cfg.ID.ResourceGroup,
				cfg.Cfg.Name,
			),
			"Deployment Complete!",
		)
	}
	return ret
}
