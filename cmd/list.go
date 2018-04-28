package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	listTaskCmd = &cobra.Command{
		Use:   "tasks",
		Short: "Use this command to list tasks",
		Long:  `Not available yet`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Prepping")
		},
	}

	listProjectCmd = &cobra.Command{
		Use:   "projects",
		Short: "Use this command to list tasks",
		Long:  `Not available yet`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Prepping")
		},
	}

	listUsersCmd = &cobra.Command{
		Use:   "users",
		Short: "Use this command to list tasks",
		Long:  `Not available yet`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Prepping")
		},
	}
)
