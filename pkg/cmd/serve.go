package cmd

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/spf13/cobra"
)

func serve(args []string) {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Echo: "+c.Request().URL.Path)
	})

	fmt.Println("Server started on port 8080...")
	e.Start(":8080")
}

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Starts an echo server",
	Run: func(cmd *cobra.Command, args []string) {
		serve(args)
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
