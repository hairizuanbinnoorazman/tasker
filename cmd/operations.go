package cmd

import (
	"github.com/hairizuanbinnoorazman/tasker/asana"
	"github.com/hairizuanbinnoorazman/tasker/generic"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	createCmd = &cobra.Command{
		Use:   "create",
		Short: "Use this command to create tasks",
		Long:  `Not available yet`,
		Run: func(cmd *cobra.Command, args []string) {
			token := viper.Get("asana_personal_token").(string)
			generic.CreateTask(asana.Asana{Token: token}, proj, name)
		},
	}
)

func opsCmdFlags() {
	createCmd.Flags().StringVar(&name, "name", "", "No help available")
	createCmd.Flags().StringVar(&desc, "desc", "undefined", "No help available")

	createCmd.Flags().StringVar(&tool, "tool", "undefined", "No help available")
	createCmd.Flags().StringVar(&proj, "proj", "", "No help available")
}
