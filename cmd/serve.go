/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

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
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"gitlab.cloudwalk.work/product-center/go_starter.git/configs"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start Server of Golang",
	Long:  `Start server of Golang`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("serve starting...")
		r := gin.Default()
		configs.InitApplication()

		r.Run() // listen and serve on 0.0.0.0:8080
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
