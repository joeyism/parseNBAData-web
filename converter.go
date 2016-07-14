package main

import(
	"encoding/json"
)

type gameDate struct{
	Year string
	Month string
	Day string
	Away string
	Home string
	GameId string
}

func convertArrayToString(arr []string) string{
	ret :="["
	for i := 0; i < len(arr); i++ {
		gameDate := convertGameIdtoGameDate(arr[i])
		b, err := json.Marshal(gameDate)
		showError(err)
		ret += string(b)
		if (i < len(arr) -1){
			ret += ", "
		}
	}
	ret += "]"
	return ret
}

func convertGameIdtoGameDate(gameId string) *gameDate{
	year := gameId[0:4]	
	month := gameId[4:6]
	day := gameId[6:8]
	away := gameId[8:11]
	home := gameId[11:14]
	return &gameDate{Year: year, Month: month, Day: day, GameId: gameId, Away: away, Home: home}
} 

