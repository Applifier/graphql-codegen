package cmd

import (
	"io/ioutil"
	"path"

	"github.com/Applifier/graphql-codegen/codegen"
	"github.com/Applifier/graphql-codegen/config"
	"github.com/spf13/cobra"
)

func init() {
	var schemaFile string
	var configFile string
	var packageName string
	var outputDir string

	var generateCmd = &cobra.Command{
		Use:   "generate",
		Short: "Generate go code for a graphql schema",
		Long:  "Generate go code for a graphql schema",
		Run: func(cmd *cobra.Command, args []string) {
			var conf config.Config
			conf.Package = packageName

			if configFile != "" {
				fileBytes, err := ioutil.ReadFile(configFile)
				if err != nil {
					panic(err)
				}
				conf, err = config.Parse(string(fileBytes))
				if err != nil {
					panic(err)
				}
			}

			schemaBytes, err := ioutil.ReadFile(schemaFile)
			if err != nil {
				panic(err)
			}

			cg := codegen.NewCodeGen(string(schemaBytes), conf)
			fileMap, err := cg.Generate()
			if err != nil {
				panic(err)
			}

			for filename, fileContent := range fileMap {
				err := ioutil.WriteFile(path.Join(outputDir, filename), []byte(fileContent), 0644)
				if err != nil {
					panic(err)
				}
			}
		},
	}

	RootCmd.AddCommand(generateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	generateCmd.PersistentFlags().StringVarP(&schemaFile, "schema", "s", "graphql.schema", "graphql.schema file")
	generateCmd.PersistentFlags().StringVarP(&configFile, "config", "c", "", "Optional configuration file. Default options are used if path is not defined")
	generateCmd.PersistentFlags().StringVarP(&packageName, "package", "p", "main", "Package name for generated files")
	generateCmd.PersistentFlags().StringVarP(&outputDir, "output", "o", ".", "Output directory. Defaults to current working directory")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// generateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
