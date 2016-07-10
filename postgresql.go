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

func (postgresl *Postgresql) getPlayersFrom(short string){
	_, _= postgresl.db.Query("SELECT game_id, play_no, time_left, team, first_player, first_action, first_stat, second_player, second_action, second_stat FROM plays WHERE team='" + short + "'");
}
