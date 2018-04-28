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
		Short: "Use this command to list tasks",
		Long:  `Not available yet`,
		// Run: func(cmd *cobra.Command, args []string) {
		// 	token := viper.Get("asana_personal_token").(string)
		// 	for a, val := range args {
		// 		fmt.Println(a)
		// 		fmt.Println(val)
		// 	}
		// 	fmt.Println(token)
		// 	// if proj == "" {
		// 	// 	generic.ListProjects(asana.Asana{Token: token})
		// 	// } else {
		// 	// 	// generic.ListTasks(asana.Asana{Token: token}, proj)
		// 	// 	generic.ListUsers(asana.Asana{Token: token}, proj)
		// 	// }
		// },
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
