/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"io/ioutil"
	"log"
	"net/http"
	"bytes"
	//"encoding/json"
	"os"
	"bufio"
	"strings"

	"github.com/spf13/cobra"
)

// putCmd represents the put command
var putCmd = &cobra.Command{
	Use:   "put",
	Short: "Make PUT request to the given URL",
	Run: func(cmd *cobra.Command, args []string) {
		client := &http.Client{}
		//Recieve url from user
		url := args[0]
		//Recieve input data from user
		scanner := bufio.NewScanner(os.Stdin)

		putBody := bytes.NewBuffer(nil)
		for scanner.Scan() {
			putBody.WriteString(scanner.Text())
			// Check if the user has typed in the `}` character.
			if strings.Contains(scanner.Text(),"}"){
				break
			}

			
		}
		//Make POST request to the given URL
		req, err := http.NewRequest("PUT", url, putBody)
		if err != nil {
			log.Fatalln(err)
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
	rootCmd.AddCommand(putCmd)
}
