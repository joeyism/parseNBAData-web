package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct{
	username string
	password string
	dbname string
}

type Game struct{
	game_id string
	team1 string
	team2 string
	date string
}

var postgres Postgresql

func readConfig() Config{
	file, _ := os.Open("config.json")
	decoder := json.NewDecoder(file)
	config := Config{}
	err := decoder.Decode(&config)
	if err != nil {
  		fmt.Println("error:", err)
	}
	return config
}

func createPostgresqlString(config Config) string{
	return "user="+config.username + " password="+config.password + " dbname="+config.dbname+" sslmode=verify-full"
}

func connectToPostgres() {
	postgres.connect(createPostgresqlString(readConfig()))
}

func getPlayersFromTeamPostgres(teamName string){
	_ = postgres.getPlayersFrom(teamName)
}

func getGamesFromTeams(team1 string, team2 string) []*Game{
	rows = postgres.getGameIdFromTeams(team1, team2)
}
