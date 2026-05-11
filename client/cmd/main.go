package main

import (
	"os"

	httpclient "yadroTestAssignment/client/internal/infrastructure/http"
	"yadroTestAssignment/client/internal/presentation/commands"
	"yadroTestAssignment/client/internal/service"
)

func main() {
	host := os.Getenv("DNS_SERVER_HOST")
	if host == "" {
		host = "http://localhost:8000"
	}

	client := httpclient.NewClient(host)
	svc := service.NewDNS(client)

	commands.Execute(svc)
}
