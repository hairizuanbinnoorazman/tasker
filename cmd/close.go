package cmd

import (
	"github.com/hairizuanbinnoorazman/tasker/asana"
	"github.com/hairizuanbinnoorazman/tasker/generic"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	closeTaskCmd = &cobra.Command{
		Use:   "tasks",
		Short: "Use this command to set tasks to complete",
		Long:  `Not available yet`,
		Run: func(cmd *cobra.Command, args []string) {
			token := viper.Get("asana_personal_token").(string)

			for _, taskID := range args {
				generic.CompleteTask(asana.Service{Token: token}, "", taskID)
			}
		},
	}
)
