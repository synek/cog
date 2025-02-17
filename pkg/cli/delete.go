package cli

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/replicate/cog/pkg/client"
)

func newDeleteCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "delete",
		Short:   "Delete a version",
		RunE:    deleteModel,
		Args:    cobra.MinimumNArgs(1),
		Aliases: []string{"rm"},
	}
	return cmd
}

func deleteModel(cmd *cobra.Command, args []string) error {
	model, err := getModel()
	if err != nil {
		return err
	}
	cli := client.NewClient()
	for _, id := range args {
		if err := cli.DeleteVersion(model, id); err != nil {
			return err
		}
		fmt.Printf("Deleted version %s:%s\n", model, id)
	}
	return nil
}
