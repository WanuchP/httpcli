/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Make GET request to the given URL",
	Run: func(cmd *cobra.Command, args []string) {
		client := &http.Client{}
		//Recieve url from user
		url := args[0]
		//Receive query flag
		if cmd.Flags().Changed("query") {
			// Get the values of the -q flag.
			queries, err := cmd.Flags().GetStringSlice("query")
			if err != nil {
				log.Fatalln(err)
			}

			// Concatenate the query strings to the URL.
			url = url + "?" + strings.Join(queries, "&")
		}
		//Make GET request to the given URL
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			log.Fatalln(err)
		}
		//Receive header flag
		if cmd.Flags().Changed("header") {
			// Get the values of the -H flag.
			headers, err := cmd.Flags().GetStringSlice("header")
			if err != nil {
				log.Fatalln(err)
			}
			for _, header := range headers {
				key := strings.Split(header, "=")
				if err != nil {
					log.Fatalln(err)
				}
				req.Header.Add(key[0], key[1])
			}
		}
		//Send the request
		resp, err := client.Do(req)
		if err != nil {
			log.Fatalln(err)
		}
		defer resp.Body.Close()
		//Read response body
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}
		//Convert body to string
		sb := string(body)
		log.Printf(sb)
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}
