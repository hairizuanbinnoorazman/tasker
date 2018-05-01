package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string
	name    string
	desc    string
	label   string
	tool    string
	task    string
	proj    string

	rootCmd = &cobra.Command{
		Use:   "tasker",
		Short: "A tool to create and manage tasks on the CLI",
		Long:  `tasker is a tool that provides a CLI interface to be used for managing tasks list between the various task management platforms.`,
	}

	versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Print the version number of tasker",
		Long:  `Print the version number of tasker`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Tasker v0.0.1")
		},
	}

	listCmd = &cobra.Command{
		Use:   "list",
		Short: "Use this command to list entities within a task management system",
		Long: `The list command allows one to list entities such as projects, tasks as well as users 
		that are part of a project. After listing out the entities, one would then be able to utilize other 
		aspects of the tasker cli tool to be able to manipulat the task lists accordingly.`,
	}

	closeCmd = &cobra.Command{
		Use:   "close",
		Short: "Use this command to set tasks to complete or to close tasks which are no longer relevant",
		Long: `This command allows you close tasks; however, a caveat here is that it kind of requires you 
		to run a list task command first to get the task id. With that, we then have th task ID which we 
		can then utilize to close issues/tasks.`,
	}

	assignCmd = &cobra.Command{
		Use:   "assign",
		Short: "Use this command to be able to assign tags or users to a task",
		Long:  `This command allows us to assign tags to users.`,
	}
)

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.AddCommand(createCmd)

	rootCmd.AddCommand(listCmd)
	listCmd.AddCommand(listTaskCmd)
	listCmd.AddCommand(listProjectCmd)
	listCmd.AddCommand(listUsersCmd)
	opsCmdFlags()
	listCmdFlags()

	rootCmd.AddCommand(versionCmd)

	rootCmd.AddCommand(closeCmd)
	closeCmd.AddCommand(closeTaskCmd)

	rootCmd.AddCommand(assignCmd)
	assignCmd.AddCommand(assignLabelCmd)
	assignCmdFlags()
}

// Execute cli commands
func Execute() {
	rootCmd.Execute()
}

func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		viper.SetConfigFile("./config.json")
	}

	err := viper.ReadInConfig()
	if err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
		fmt.Println()
	} else {
		fmt.Println(err.Error())
	}
}
