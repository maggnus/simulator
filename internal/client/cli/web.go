package cli

import (
	"github.com/spf13/cobra"
	"log"
	"simulator/internal/client/web"
)

var webCmd = &cobra.Command{
	Use:   "web",
	Short: "Run web API client",
	Run: func(cmd *cobra.Command, args []string) {
		log.Printf("start web server on http://localhost:8080")
		web.Server()
	},
}
