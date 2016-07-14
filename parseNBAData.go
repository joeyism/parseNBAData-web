package main

import (
	"fmt"
	"log"
	"github.com/BurntSushi/toml"
	"database/sql"
)

type Config struct{
	Username string
	Password string
	Database string
}

type Game struct{
	game_id string
	team1 string
	team2 string
	date string
}

var postgres Postgresql

func readConfig() Config{
	var config Config
	if _, err := toml.DecodeFile("config.toml", &config); err != nil {
		log.Fatal(err)
	}
	return config
}

func createPostgresqlString(config Config) string{
	requestString := "user='"+config.Username + "' password='"+config.Password + "' dbname='"+config.Database+"' sslmode=disable"
	fmt.Println(requestString)
	return requestString
}

func connectToPostgres() {
	postgres.connect(createPostgresqlString(readConfig()))
}

func getPlayersFromTeamPostgres(teamName string){
	_ = postgres.getPlayersFrom(teamName)
}

func extractGameIdsFromRows(rows *sql.Rows) []string{
	var gameIds []string
	defer rows.Close()
	for rows.Next(){
		var matchup string
		err := rows.Scan(&matchup)
		showError(err)	
		gameIds = append(gameIds, matchup)
	}
	return gameIds
}

func getGamesFromTeams(team1 string, team2 string) []string{
	rows := postgres.getGameIdFromTeams(team1, team2)
	return extractGameIdsFromRows(rows)
}
