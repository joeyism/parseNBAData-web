package main

import(
	"fmt"
	"github.com/valyala/fasthttp"
    "github.com/buaazp/fasthttprouter"
	"log"
	"flag"
	"strings"
)

const port = ":8080"

func getPlayersFromTeam(){
	// ctx.JSON(iris.StatusOK, map[string]string{"team": team})
}

var (
	compress           = flag.Bool("compress", false, "Enables transparent response compression if set to true")
	byteRange          = flag.Bool("byteRange", false, "Enables byte range requests if set to true")
	dir                = flag.String("dir", "./public", "Directory to serve static files from")
	generateIndexPages = flag.Bool("generateIndexPages", true, "Whether to generate directory index pages")
)

func Test(ctx *fasthttp.RequestCtx, _ fasthttprouter.Params) {
	fmt.Fprintf(ctx, "Hello, world!\n\n")

	fmt.Fprintf(ctx, "Request method is %q\n", ctx.Method())
	fmt.Fprintf(ctx, "RequestURI is %q\n", ctx.RequestURI())
	fmt.Fprintf(ctx, "Requested path is %q\n", ctx.Path())
	fmt.Fprintf(ctx, "Host is %q\n", ctx.Host())
	fmt.Fprintf(ctx, "Query string is %q\n", ctx.QueryArgs())
	fmt.Fprintf(ctx, "User-Agent is %q\n", ctx.UserAgent())
	fmt.Fprintf(ctx, "Connection has been established at %s\n", ctx.ConnTime())
	fmt.Fprintf(ctx, "Request has been started at %s\n", ctx.Time())
	fmt.Fprintf(ctx, "Serial request number for the current connection is %d\n", ctx.ConnRequestNum())
	fmt.Fprintf(ctx, "Your ip is %q\n\n", ctx.RemoteIP())

	fmt.Fprintf(ctx, "Raw request is:\n---CUT---\n%s\n---CUT---", &ctx.Request)

	ctx.SetContentType("text/plain; charset=utf8")

	// Set arbitrary headers
	ctx.Response.Header.Set("X-My-Header", "my-header-value")

	// Set cookies
	var c fasthttp.Cookie
	c.SetKey("cookie-name")
	c.SetValue("cookie-value")
	ctx.Response.Header.SetCookie(&c)
}

func GetPlayersFromTeamHandler(ctx *fasthttp.RequestCtx, params fasthttprouter.Params){
	path := string(ctx.Path())
	team := path[len("/api/team/"):len(path)]
	log.Print(team)
}

func GetMatchupsBetweenTeams(ctx *fasthttp.RequestCtx, params fasthttprouter.Params){
	url := strings.Split(string(ctx.Path()), "/")
	team1 := url[2]
	team2 := url[4]
	gameIds := getGamesFromTeams(team1, team2)
	ctx.SetContentType("application/json")
	fmt.Fprintf(ctx, convertArrayToString(gameIds))
}

func main(){
	fmt.Println("Server started listening to port " + port)
	connectToPostgres()

	router := fasthttprouter.New()
    router.GET("/test", Test)
	router.GET("/api/:team1/vs/:team2", GetMatchupsBetweenTeams)
	router.NotFound = fasthttp.FSHandler("./public", 0)

    log.Fatal(fasthttp.ListenAndServe(":8080", router.Handler))
}
