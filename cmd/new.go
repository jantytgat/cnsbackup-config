/*
Copyright © 2021 Jan Tytgat

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package cmd

import (
	"github.com/jantytgat/cnsbackup-config/controllers"
	"github.com/jantytgat/cnsbackup-models"
	"github.com/spf13/cobra"
	"log"
)

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Create a new configuration for cnsbackup",

	Run: func(cmd *cobra.Command, args []string) {
		runNew()
	},
}

func init() {
	rootCmd.AddCommand(newCmd)
}

func runNew() {
	var err error

	c := controllers.NewConfigurationController{
		Configuration: models.Configuration{
			Targets:  nil,
			Settings: models.Settings{},
		}}

	err = c.Run()
	if err != nil {
		log.Fatal(err)
	}

	err = updateConfiguration(c.Configuration)
	if err != nil {
		log.Fatal(err)
	}
}
