package cmd

import (
	"github.com/hairizuanbinnoorazman/tasker/asana"
	"github.com/hairizuanbinnoorazman/tasker/generic"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	assignLabelCmd = &cobra.Command{
		Use:   "label",
		Short: "Use this command to set a label onto a task",
		Long: `You can provide a list of task IDs that you would want to set the a label on. The 
		compulsory flags that needs to be assigned is the project as well as label flag.
		
		To effectively utilize this command, you would probably need to run a command to list the tasks first;
		after which you can then set the flags and then add the list of commands that you would want to apply 
		your label to.`,
		Run: func(cmd *cobra.Command, args []string) {
			token := viper.Get("asana_personal_token").(string)

			for _, taskID := range args {
				generic.AssignLabel(asana.Service{Token: token}, proj, taskID, label)
			}
		},
	}
)

func assignCmdFlags() {
	assignLabelCmd.Flags().StringVar(&proj, "proj", "", "No help available")
	assignLabelCmd.Flags().StringVar(&label, "label", "undefined", "No help available")
}
