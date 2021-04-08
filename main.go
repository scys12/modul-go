package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Response struct {
	Name           string  `json:"name"`
	BaseExperience int     `json:"base_experience"`
	Weight         int     `json:"weight"`
	Stats          []Stats `json:"stats"`
}

type Stats struct {
	BaseStat int         `json:"base_stat"`
	Stat     PokemonStat `json:"stat"`
}

type PokemonStat struct {
	StatName string `json:"name"`
}

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Halo Keluarga RPL!")
	fmt.Println("-------------------")
	fmt.Println("Pilih pokemon:")
	fmt.Println("1. Bulbasaur")
	fmt.Println("2. Pikachu")
	fmt.Printf("\nMasukkan pilihan: ")

	char, _, err := reader.ReadRune()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	url := "http://pokeapi.co/api/v2/pokemon/"

	switch char {
	case '1':
		url += "bulbasaur"
	case '2':
		url += "pikachu"
	}

	response, err := http.Get(url)

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var responseObject Response
	json.Unmarshal(responseData, &responseObject)

	fmt.Printf("Name: %v\n", responseObject.Name)
	fmt.Printf("Base Experience: %v\n", responseObject.BaseExperience)
	fmt.Printf("Weight: %v\n", responseObject.Weight)

	for i := 0; i < len(responseObject.Stats); i++ {
		fmt.Printf("%v: %v\n", responseObject.Stats[i].Stat.StatName, responseObject.Stats[i].BaseStat)
	}

}
