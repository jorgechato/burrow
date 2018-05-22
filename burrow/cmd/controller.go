package cmd

import (
	"fmt"

	"github.com/jorgechato/burrow/generator"
	"github.com/jorgechato/burrow/generator/startapp"
	"github.com/jorgechato/burrow/generator/startproject"
	"github.com/spf13/cobra"
)

func startApp(cmd *cobra.Command, args []string) {
	app, err := generator.New(args[0], dbPtr)

	if err != nil {
		fmt.Print(err)
		return
	}

	startapp.Run(app)

	fmt.Printf("Congratulations! Your application, %s, has been successfully built!\n\n", app.Name)
	fmt.Printf("You can find your new application at:\n%v\n", app.Root)
	fmt.Print("With love by adidas.")
}

func startProject(cmd *cobra.Command, args []string) {
	app, err := generator.New(args[0], dbPtr)

	if err != nil {
		fmt.Print(err)
		return
	}

	startproject.Run(app)

	fmt.Printf("Congratulations! Your project, %s, has been successfully built!\n\n", app.Name)
	fmt.Printf("You can find your new project at:\n%v\n", app.Root)
	fmt.Print("With love by adidas.")
}
