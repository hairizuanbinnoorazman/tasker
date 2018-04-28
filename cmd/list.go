package cmd

import (
	"github.com/hairizuanbinnoorazman/tasker/asana"
	"github.com/hairizuanbinnoorazman/tasker/generic"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	listTaskCmd = &cobra.Command{
		Use:   "tasks",
		Short: "Use this command to list tasks",
		Long:  `Not available yet`,
		Run: func(cmd *cobra.Command, args []string) {
			token := viper.Get("asana_personal_token").(string)
			generic.ListTasks(asana.Service{Token: token}, proj)
		},
	}

	listProjectCmd = &cobra.Command{
		Use:   "projects",
		Short: "Use this command to list tasks",
		Long:  `Not available yet`,
		Run: func(cmd *cobra.Command, args []string) {
			token := viper.Get("asana_personal_token").(string)
			generic.ListProjects(asana.Service{Token: token})
		},
	}

	listUsersCmd = &cobra.Command{
		Use:   "users",
		Short: "Use this command to list tasks",
		Long:  `Not available yet`,
		Run: func(cmd *cobra.Command, args []string) {
			token := viper.Get("asana_personal_token").(string)
			generic.ListUsers(asana.Service{Token: token}, proj)
		},
	}
)

func listCmdFlags() {
	listTaskCmd.Flags().StringVar(&proj, "proj", "", "No help available")
	listUsersCmd.Flags().StringVar(&proj, "proj", "", "No help available")
}
