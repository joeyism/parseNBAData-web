package main

import(
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

type Postgresql struct{
	db *sql.DB
}

func (postgresl *Postgresql) connect(config string) {
	var err error
	postgresl.db, err = sql.Open("postgres", config)
	//db, err = sql.Open("postgres",createPostgresqlString(readConfig()))
	if err != nil {
		log.Fatal(err)
	}
}

func (postgresql *Postgresql) getPlayersFrom(short string) *sql.Rows{
	rows, err := postgresql.db.Query("SELECT game_id, play_no, time_left, team, first_player, first_action, first_stat, second_player, second_action, second_stat FROM plays WHERE team='" + short + "'");
	if err != nil{
		log.Fatal(err)
	}
	return rows
}

func (postgresql *Postgresql) getGameIdFromTeams(team1 string, team2 string) *sql.Rows{
	rows, err := postgresql.db.Query("SELECT DISTINCT game_id FROM plays WHERE game_id similar to '%(" + team1+team2 + "|" + team2+team1 + ")%'")
	if err != nil{
		log.Fatal(err)
	}
	return rows
}
