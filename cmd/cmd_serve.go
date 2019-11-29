package cmd

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"

	"cw-app/hotpot-backend/configs"
	"cw-app/hotpot-backend/pkg/connection"
	"cw-app/hotpot-backend/pkg/routers"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start Server of Golang",
	Long:  `Start server of Golang`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("serve starting...")

		r := gin.Default()
		configs.InitConfig()

		pg := connection.InitPGConnection()
		defer pg.Close()

		routers.InitRouters(r)
		r.Run(":8081") // listen and serve on 0.0.0.0:8080
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
