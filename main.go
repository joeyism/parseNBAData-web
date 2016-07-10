package main

import(
	"github.com/kataras/iris"
	"fmt"
	"strings"
)

const port = "8080"

func getPlayersFromTeam(ctx *iris.Context){
	team := strings.Split(ctx.PathString(),"/")[3]

	ctx.JSON(iris.StatusOK, map[string]string{"team": team})
}

func main(){
	fmt.Println("Server started listening to port " + port)
	connectToPostgres()

	static := iris.New()
	static.StaticWeb("/", "./public", 0)

	api := iris.New()
	apiParty := api.Party("/api")
	{
		apiParty.Get("/team/:team", getPlayersFromTeam)
	}


	go static.Listen(":8081")
	api.Listen(":8080")
	
}
