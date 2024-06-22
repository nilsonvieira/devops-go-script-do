package main

import (
	"context"
	"fmt"
	"github.com/digitalocean/godo"
	"github.com/olekukonko/tablewriter"
	"golang.org/x/oauth2"
	"os"
)

func main() {
	token := os.Getenv("DIGITALOCEAN_ACCESS_TOKEN")
	if token == "" {
		fmt.Println("DIGITALOCEAN_ACCESS_TOKEN Variável de Ambiente não Definida!")
		return
	}

	oauthClient := oauth2.NewClient(context.Background(), oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	))

	client := godo.NewClient(oauthClient)

	droplets, _, err := client.Droplets.List(context.Background(), &godo.ListOptions{})
	if err != nil {
		fmt.Printf("Erro ao Buscar Droplets: %s\n", err)
		return
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Name", "ID", "Internal IP", "Public IP", "Volumes", "Cost/Month ($)"})

	for _, droplet := range droplets {
		internalIP := ""
		publicIP := ""
		volumes := ""

		// Get internal and public IPs
		for _, network := range droplet.Networks.V4 {
			if network.Type == "private" {
				internalIP = network.IPAddress
			} else if network.Type == "public" {
				publicIP = network.IPAddress
			}
		}

		// Get volumes
		dropletVolumes, _, err := client.Storage.ListVolumes(context.Background(), &godo.ListVolumeParams{})
		if err != nil {
			fmt.Printf("Erro ao Buscar Volumes: %s\n", err)
			return
		}
		for _, volume := range dropletVolumes {
			for _, attachment := range volume.DropletIDs {
				if attachment == droplet.ID {
					volumes += fmt.Sprintf("ID: %s, Name: %s, Size: %dGB\n", volume.ID, volume.Name, volume.SizeGigaBytes)
				}
			}
		}

		// Calculate cost
		cost := float64(droplet.Size.PriceHourly * 24 * 30)

		table.Append([]string{
			droplet.Name,
			fmt.Sprintf("%d", droplet.ID),
			internalIP,
			publicIP,
			volumes,
			fmt.Sprintf("%.2f", cost),
		})
	}

	table.Render()
}
