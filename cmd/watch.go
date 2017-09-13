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
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"image"
	"image/color"
	"image/jpeg"
	"log"
	"net/http"
	"time"
)

type Config struct {
	CameraUrl string    `mapstructure:"camera_url"`
	Api       ApiConfig `mapstructure:"api"`
}

type ApiConfig struct {
	Host     string     `mapstructure:"host"`
	BasePath string     `mapstructure:"base_path"`
	Auth     AuthConfig `mapstructure:"auth"`
}

type AuthConfig struct {
	Id  string `mapstructure:"id"`
	Key string `mapstructure:"key"`
}

var c Config

// watchCmd represents the watch command
var watchCmd = &cobra.Command{
	Use:   "watch",
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

		fmt.Println(c)

		u, _ := ultimaker.New(&ultimaker.Ultimaker3MachineConfig{
			Host:     c.Api.Host,
			BasePath: c.Api.BasePath,
			Auth: &ultimaker.Ultimaker3AuthConfig{
				Id:  c.Api.Auth.Id,
				Key: c.Api.Auth.Key,
			},
		})
		firstImage, err := takePhoto()

		if err != nil {
			log.Fatalln(err)
			return
		}

		rect := firstImage.Bounds()
		r, g, b, a := firstImage.At(rect.Max.X/2, rect.Max.Y/2).RGBA()
		u.FirstColor = &color.RGBA{uint8(r >> 8), uint8(g >> 8), uint8(b >> 8), uint8(a >> 8)}

		err = u.SetLedColor(color.RGBA{255, 255, 255, 0})

		//err = u.SetLedColor(color.RGBA{uint8(r >> 8), uint8(g >> 8), uint8(b >> 8), uint8(a >> 8)})
		log.Println(err)
		t := time.NewTicker(500 * time.Millisecond)
		defer t.Stop()
		for {
			select {
			case <-t.C:
				img, err := takePhoto()
				if err != nil {
					time.Sleep(1 * time.Second)
					panic("can't get color")
				}
				rect := img.Bounds()
				r, g, b, a := img.At(rect.Max.X/2, rect.Max.Y/2).RGBA()
				err = u.SetLedColor(color.RGBA{uint8(r >> 8), uint8(g >> 8), uint8(b >> 8), uint8(a >> 8)})
				log.Println(err)
			}
		}

	},
}

func takePhoto() (image.Image, error) {
	res, err := http.Get(c.CameraUrl)

	if err != nil {
		log.Printf("failed get camera snapshot %v %v", c.CameraUrl, err)
		return nil, err
	}

	return jpeg.Decode(res.Body)
}

func init() {
	RootCmd.AddCommand(watchCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// watchCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// watchCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
