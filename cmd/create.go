package cmd

import (
	"github.com/lemuelZara/to-cobra/internal/database"
	"github.com/spf13/cobra"
)

type CreateCategoryParams struct {
	Name        string
	Description string
}

var params CreateCategoryParams

func newCreateCmd(db database.Category) cobra.Command {
	return cobra.Command{
		Use:   "create",
		Short: "",
		Long:  "",
		RunE:  createCategory(db),
	}
}

func createCategory(categoryDB database.Category) RunEFunc {
	return func(cmd *cobra.Command, args []string) error {
		_, err := categoryDB.Save(params.Name, params.Description)
		if err != nil {
			return err
		}

		return nil
	}
}

func init() {
	createCmd := newCreateCmd(GetCategoryDB(GetDB()))
	categoryCmd.AddCommand(&createCmd)

	createCmd.Flags().StringVarP(&params.Name, "name", "n", "", "Name of the category")
	createCmd.Flags().StringVarP(&params.Description, "description", "d", "", "Description of the category")
	createCmd.MarkFlagsRequiredTogether("name", "description")
}
