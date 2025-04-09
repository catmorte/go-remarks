package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/catmorte/go-remarks/internal/templates"
	"github.com/spf13/cobra"
)

var (
	vars    = map[string]string{}
	cfgPath string

	resultFolder = ".result"
)

func assert(err error, s string, args ...any) {
	if err != nil {
		fmt.Println(fmt.Sprintf(s, args...), err)
		os.Exit(1)
	}
}

func assertOK(ok bool, s string, args ...any) {
	if !ok {
		fmt.Println(fmt.Sprintf(s, args...))
		os.Exit(1)
	}
}

var rootCmd = &cobra.Command{
	Use:   "go-remarks",
	Short: "go-remarks is a sample CLI application to generate remarks",
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println("go-remarks is a CLI application to generate remarks. use --help for detail")
		fmt.Println()
		fmt.Println("internal templates")
		for _, v := range templates.InternalTemplates() {
			fmt.Println(" - " + v.GetName())
		}
		fmt.Println()

	},
}

var templateVarsCmd = &cobra.Command{
	Use:   "template_vars",
	Short: "returns all possible template's vars",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		dts, err := templates.GetDefinedTemplates(cfgPath)
		assert(err, "failed to get defined templates")
		dt, err := dts.FindByName(args[0])
		for _, v := range dt.GetVars() {
			fmt.Println(v)
		}
	},
}

var templatesCmd = &cobra.Command{
	Use:   "templates",
	Short: "returns all available templates declared in $HOME/.config/go-remarks folder",
	Run: func(cmd *cobra.Command, args []string) {
		definedTypes, err := templates.GetDefinedTemplates(cfgPath)
		assert(err, "can't get defined types")
		for _, v := range definedTypes {
			fmt.Println(v.GetName())
		}
	},
}

var compileCmd = &cobra.Command{
	Use:   "generate",
	Short: "generate the remark",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		dts, err := templates.GetDefinedTemplates(cfgPath)
		assert(err, "failed to get defined types")
		dt, err := dts.FindByName(args[0])
		assert(err, "failed to get defined type")
		if vars == nil {
			vars = map[string]string{}
		}
		err = dt.Compile(vars)
		assert(err, "failed to compile")
	},
}

func main() {
	compileCmd.Flags().StringToStringVar(&vars, "vars", nil, "key-value parameters (e.g. --vars key1=value1 --vars key2=value2)")
	rootCmd.AddCommand(templatesCmd)
	rootCmd.AddCommand(compileCmd)
	rootCmd.AddCommand(templateVarsCmd)

	dirname, err := os.UserHomeDir()
	assert(err, "can't get user's home dir")

	cfgPath = filepath.Join(dirname, ".config", "go-remarks")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
