package main

import (
	"os/exec"
	"fmt"
	"strings"
	//"strconv"
)

func getData(playerName string)(string, error){
	cmd := exec.Command("./python/parseNBAData.py", "playbyplay20092010reg20100418", "20091027BOSCLE", playerName)
	out, err := cmd.Output()

	if err != nil {
		println("error!")
		println(err.Error())
		return "", err
	}

	return string(out), nil;
}

func indexOf(s []string, e string) int{
	for i, a := range s {
		if a == e {
			return i
		}
	}
	return -1 
}

func printSlice(sl []string){
	for _, s := range sl{
		fmt.Printf(s + " ")
	}
}

func newline(){
	fmt.Printf("\n")
}

func printSliceln(sl []string){
	printSlice(sl)
	newline()
}

func getBracketInformation(playLine string, playerName string){
	fmt.Printf(playLine + "|| ")
	nameIndex := strings.Index(playLine, playerName)
	restOfLine := playLine[nameIndex:len(playLine)]
	indexOpenB := strings.Index(restOfLine, "(")
	indexCloseB := strings.Index(restOfLine, ")")
	fmt.Printf(restOfLine + "||")
	if indexOpenB != -1 && indexCloseB != -1 {
		withinBracket := restOfLine[indexOpenB + 1:indexCloseB]
		fmt.Printf(withinBracket + " ")
	}
	newline()
}

func getPlayInformation(playLine []string, playerName string){
	nameIndex := indexOf(playLine, playerName)
	//ret := []string{}
	if nameIndex == -1 {
		printSliceln(playLine)
	} else if nameIndex != len(playLine)-1 {
		getBracketInformation(strings.Join(playLine, " "), playerName)
	//	ret = []string{ playLine[nameIndex], playLine[nameIndex+1] }
	} else {
	//	ret = []string{ playLine[nameIndex - 1], playLine[nameIndex]}
	}
}

func parseData(data string, playerName string){
	totalPlays := [][] string{}
	for _, playLine := range strings.Split(data, "\n"){
		playLineSlice := strings.Split(playLine, " ")
		time := playLineSlice[0]
		getPlayInformation(playLineSlice[1:len(playLineSlice)], playerName)
		playLineSplice := []string{time}
		totalPlays = append(totalPlays, playLineSplice)
	}
}

func main() {
	playerName := "Garnett"
	result, _:= getData(playerName)
	parseData(result, playerName)
	//fmt.Printf(result)
}
