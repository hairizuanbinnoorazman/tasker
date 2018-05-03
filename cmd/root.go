package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/user"
	"path/filepath"

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

	initCmd = &cobra.Command{
		Use:   "init",
		Short: "Use this command to initialize configurations in the default .tasker folder",
		Long: `The init command would create a config.json file in the .tasker folder of your home directory.
		This would allow the configuration to be used everywhere in your computer. In order to fully utilize the 
		cli, you would also need to add tasker into the path arguments as well.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Initializing configuration")

			usr, err := user.Current()
			if err != nil {
				fmt.Println("Unable to retrive current user")
				return
			}

			folderPath := filepath.Join(usr.HomeDir, ".tasker")
			filePath := filepath.Join(folderPath, "config.json")

			_, err = os.Stat(filePath)
			if err == nil {
				fmt.Println("File exists. Will not re-initialize")
				return
			}

			err = os.MkdirAll(folderPath, os.ModePerm)
			if err != nil {
				fmt.Println(err.Error())
				fmt.Println("Folder unable to be created")
				return
			}

			initialConfiguration := Config{
				PrimaryTool: ToolConfig{
					Name:             "asana",
					Token:            "example-token",
					DefaultProjectID: "example-default-project",
				},
				SecondaryTools: []ToolConfig{
					ToolConfig{
						Name:             "github",
						Token:            "example-token",
						DefaultProjectID: "example-default-project",
					},
				},
			}

			initialConfigBytes, errByte := json.MarshalIndent(initialConfiguration, "", "\t")
			if errByte != nil {
				fmt.Println("Unable to retrive bytestrem of configuration")
				return
			}
			fileWriteErr := ioutil.WriteFile(filePath, initialConfigBytes, 0644)
			if fileWriteErr != nil {
				fmt.Println(fileWriteErr.Error())
				return
			}
		},
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

	rootCmd.AddCommand(initCmd)

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
	usr, err := user.Current()
	if err != nil {
		panic("Error in trying to get current user.")
	}

	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		viper.SetConfigFile(fmt.Sprintf("%v/.tasker/config.json", usr.HomeDir))
	}

	errRead := viper.ReadInConfig()
	if errRead == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
		fmt.Println()
	} else {
		fmt.Println(errRead.Error())
	}
}
