/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"io"
	"log"
	"net/http"
)

// randomCmd represents the random command
var randomCmd = &cobra.Command{
	Use:   "random",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		getRandomJoke()
	},
}

func init() {
	rootCmd.AddCommand(randomCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// randomCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// randomCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

type Joke struct {
	ID     string `json:"id"`
	Joke   string `json:"joke"`
	Status int    `json:"status"`
}

func getRandomJoke() {
	url := "https://icanhazdadjoke.com/"
	data := getJokeData(url)
	joke := Joke{}
	err := json.Unmarshal(data, &joke)
	if err != nil {
		log.Printf("couldn't Unmarshal data, %v", err)
	}
	fmt.Println(joke.Joke)
}

func getJokeData(apiURL string) []byte {
	request, err := http.NewRequest(http.MethodGet, apiURL, nil)
	if err != nil {
		log.Printf("couldn't request jokeData %v", err)
	}
	request.Header.Add("Accept", "application/json")
	request.Header.Add("User-Agent", "Dadjoke CLI (https://github.com/example/dadjoke)")
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Printf("get dadjoke error. %v", err)
	}
	content, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("get resp body error. %v", err)
	}
	fmt.Println(string(content))
	return content
}
