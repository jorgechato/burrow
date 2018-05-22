package cmd

import "github.com/spf13/cobra"

var (
	startAppCmd = &cobra.Command{
		Use:   "startapp [APPNAME]",
		Short: "Application module generator",
		Long: `Module generator.
This directory structure will house the module application.
You can build it without a project and plug in into any project.`,
		Args: cobra.ExactArgs(1),
		Run:  startApp,
	}
	startProjectCmd = &cobra.Command{
		Use:   "startproject [PROJECTNAME]",
		Short: "Project generator",
		Long: `Project generator.
This directory structure will house the main project.
You can add as many applications as your project require.`,
		Args: cobra.ExactArgs(1),
		Run:  startProject,
	}
	dbPtr bool

	rootCmd = &cobra.Command{Use: "burrow"}
)

func init() {
	startAppCmd.Flags().BoolVarP(
		&dbPtr,
		"database",
		"d",
		false,
		"Add a database support for your app.",
	)
	startProjectCmd.Flags().BoolVarP(
		&dbPtr,
		"database",
		"d",
		false,
		"Add a database support for your app.",
	)
}
