package main

import (
	sdtd "github.com/digitalfiz/sdtd-stats-go"
	"log"
)

func main() {
	client := sdtd.SdtdStats{}
	client.SetConnection("0.0.0.0", 25004)
	client.SetLogin("admin", "sometoken")

	players := client.GetPlayersLocation()
	log.Println(players)

	// players := client.GetPlayersOnline()
	// log.Println(players)

	// inventory := client.GetPlayerInventory("someplayerid")
	// log.Println(inventory)

	// updates := client.GetWebUIUpdates()
	// log.Println(updates)

	// stats := client.GetStats()
	// log.Println(stats)

	// claims := client.GetLandClaims()
	// log.Println(claims)

	// entities := client.GetHostileLocation()
	// log.Println(entities)

	// 	entities := client.GetAnimalsLocation()
	// 	log.Println(entities)
}
