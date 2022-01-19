package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

func Generate() *cli.App {
	app := cli.NewApp()
	app.Name = "Aplicação de linha de comando"
	app.Usage = "Busca de IPS e nome de Servidor na internet"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "google.com",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Busca de enderecos na net",
			Flags:  flags,
			Action: searchIps,
		},
		{
			Name:   "servidores",
			Usage:  "Busca o servidores na net",
			Flags:  flags,
			Action: searchServer,
		},
	}
	return app
}

func searchIps(c *cli.Context) {
	host := c.String("host")
	ips, err := net.LookupIP(host)

	if err != nil {
		log.Fatal(err)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func searchServer(c *cli.Context) {
	host := c.String("host")
	servers, err := net.LookupNS(host)

	if err != nil {
		log.Fatal(err)
	}
	for _, server := range servers {
		fmt.Println(server)
	}
}
