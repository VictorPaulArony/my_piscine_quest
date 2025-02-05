package utils

import (
	"bufio"
	"lem-in/models"
	"log"
	"os"
	"strconv"
	"strings"
)

// reading the .txt file
// using either BFS/DFS to get the best path and all paths
// sort the best path from the search
// deploy the ants to the paths
// show/print/simulate the ants in the paths
// end


func FileReader (path string) models.Colony {
	colony := models.Colony

	var noOfAnts int
	var roomNo string
	var coordinates []int
	var links map[string]int

	file, err := os.Open(path)
	if err != nil {
		log.Println("Error opening the data file \n", err)
		return nil
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		func(){
			splited := strings.Split(text, " ")
			if len(splited) == 1{
				noOfAnts, err = strconv.Atoi(splited[0])
				if err != nil {
					log.Println("Error coneverting No of ants.\n", err)
				}
			} else if len(splited) == 3 {
				roomNo = splited[0]
				for v := range splited{
					x, err := strconv.Atoi(v)
				}
			}
		}()

		
	}

	return colony
}