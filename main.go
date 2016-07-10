package main

import(
	"fmt"
	"github.com/valyala/fasthttp"
	"log"
	"flag"
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

func main(){
	fmt.Println("Server started listening to port " + port)
	connectToPostgres()

	fs := &fasthttp.FS{
		Root: *dir,
		IndexNames: []string{"index.html"},
		GenerateIndexPages: *generateIndexPages,
		Compress: *compress,
		AcceptByteRange: *byteRange,
	}

	// Create request handler for serving static files.
	h := fs.NewRequestHandler()

	if err := fasthttp.ListenAndServe(port, h); err != nil {
		log.Fatalf("error in ListenAndServe: %s", err)
	}
}
