package cmd

import (
	"fmt"

	"github.com/lemuelZara/to-cobra/internal/database"
	"github.com/spf13/cobra"
)

func newListCmd(db database.Category) cobra.Command {
	return cobra.Command{
		Use:   "list",
		Short: "",
		Long:  "",
		RunE:  listCategories(db),
	}
}

func listCategories(categoryDB database.Category) RunEFunc {
	return func(cmd *cobra.Command, args []string) error {
		cs, err := categoryDB.FindAll()
		if err != nil {
			return err
		}

		for _, c := range cs {
			fmt.Printf("[ID: %s] Name: %s, Description: %s\n", c.ID, c.Name, c.Description)
		}

		return nil
	}
}

func init() {
	listCmd := newListCmd(GetCategoryDB(GetDB()))
	categoryCmd.AddCommand(&listCmd)
}
