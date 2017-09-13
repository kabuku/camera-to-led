// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"

	"github.com/kabuku/camera-to-led/ultimaker"
	"github.com/kabuku/camera-to-led/ultimaker/api/client/print_job"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"os"
)

// printCmd represents the print command
var printCmd = &cobra.Command{
	Use:   "print",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := viper.Unmarshal(&c); err != nil {
			log.Fatalf("unable to decode into struct, %v", err)
		}

		fmt.Println(&c)

		u, _ := ultimaker.New(&ultimaker.Ultimaker3MachineConfig{
			Host:     c.Api.Host,
			BasePath: c.Api.BasePath,
			Auth: &ultimaker.Ultimaker3AuthConfig{
				Id:  c.Api.Auth.Id,
				Key: c.Api.Auth.Key,
			},
		})

		f, err := os.Open("/Users/keisuke/duplo-set.gcode")
		if err != nil {
			log.Println(err)
			return
		}
		defer f.Close()

		r, err := u.Client().PrintJob.PostPrintJob(print_job.NewPostPrintJobParams().WithFile(*f).WithJobname("duplo-set.gcode"))

		log.Println(err)
		log.Println(r)


	},
}

func init() {
	RootCmd.AddCommand(printCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// printCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// printCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
