package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// load file with default filepath
	// construct graph from file
	// take input then evaluate graph

	filepath := flag.String("file", "routes.csv", "a string")
	flag.Parse()

	graph, err := createGraphFromFile(*filepath)
	if err != nil {
		fmt.Printf("cannot load %v, got err %v", filepath, err)
		os.Exit(1)
	}

	for {
		var src, dest string
		fmt.Printf("What station are you getting on the train?: ")
		fmt.Scanf("%s", &src)
		fmt.Printf("What station are you getting off the train?: ")
		fmt.Scanf("%s", &dest)

		hop, distance := graph.ShortestPath(src, dest)
		if distance >= 0 {
			fmt.Printf("Your trip from %s to %s includes %d stops and will take %d minutes", src, dest, hop-1, distance)
			fmt.Println()
		} else {
			fmt.Printf("No routes from %s to %s", src, dest)
			fmt.Println()
		}
	}
}
