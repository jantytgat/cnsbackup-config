/*
Copyright 2021 Jan Tytgat

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
	"bytes"
	"fmt"
	"github.com/jantytgat/cnsbackup-models"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
	"log"
	"os"
)

var configFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "cnsbackup-config",
	Short: "Configuration tool for cnsbackup",
	Long:  `Manage configuration files used by cnsbackup.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&configFile, "file", "f", "cnsbackup.yaml", "config file")
	cobra.OnInitialize(configureViper)
}

func configureViper() {
	viper.SetConfigFile(configFile)
	viper.SetConfigType("yaml")
}

func canLoadConfigurationFromViper() error {
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Could not load configuration file, create a new configuration by running \"cnsbackup new -f %s\"\n", configFile)
		return err
	}
	return nil
}

func getConfiguration() (models.Configuration, error) {
	var config models.Configuration
	err := viper.Unmarshal(&config)

	return config, err
}

func updateConfiguration(configuration models.Configuration) error {
	var marshalledConfiguration, err = yaml.Marshal(&configuration)
	if err != nil {
		log.Fatal(err)
	}

	err = viper.ReadConfig(bytes.NewReader(marshalledConfiguration))
	if err != nil {
		log.Fatal(err)
	}

	err = viper.SafeWriteConfigAs(configFile)
	return err
}
