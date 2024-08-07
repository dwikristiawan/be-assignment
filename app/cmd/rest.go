package cmd

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

var (
	restCmd = &cobra.Command{
		Use:   "rest",
		Short: "be-assigment",
		Run:   restServer,
	}
)

func init() {
	rootCmd.AddCommand(restCmd)
}
func restServer(cmd *cobra.Command, args []string) {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"mesage": "pong"})
	})

	err := r.Run(fmt.Sprintf("%s:%s", rootConfig.Server.HostServer, rootConfig.Server.PortServer))
	if err != nil {
		panic(err)
	}
}
