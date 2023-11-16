/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

// jokeCmd represents the joke command
var jokeCmd = &cobra.Command{
	Use:   "joke",
	Short: "run the command then you will get a random dadjoke!",
	Long:  `A longer description`,
	Run:   generateJoke,
}

type Joke struct {
	Id     string `json:"id"`
	Joke   string `json:"joke"`
	Status int    `json:"status"`
}

func init() {
	rootCmd.AddCommand(jokeCmd)

	jokeCmd.Flags().IntP("test", "u", 0, "Help message for toggle")
}

func generateJoke(cmd *cobra.Command, args []string) {
	number, _ := cmd.Flags().GetInt("test")
	fmt.Println(number)
	joke := getJoke("https://icanhazdadjoke.com/")
	fmt.Println(joke.Joke)
}

func getJoke(url string) Joke {
	// 客製自己的header就必須使用NewRequest
	request, err := http.NewRequest(
		http.MethodGet,
		url,
		nil,
	)
	if err != nil {
		log.Printf("Can not access the API %s error: %s\n", url, err)
	}
	// 增加規定要加上的header
	request.Header.Add("Accept", "application/json")
	request.Header.Add("User-Agent", "Dadjoke Generator CLI (https://github.com/Cody/cobra/dadjoke)")

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Printf("Can not send the request, error: %s\n", err)
	}
	var joke Joke
	json.NewDecoder(response.Body).Decode(&joke) // 從一個流裡面直接進行解碼，代碼精幹
	if err != nil {
		log.Printf("Can not read the response body, error: %s\n", err)
	}

	return joke
}
