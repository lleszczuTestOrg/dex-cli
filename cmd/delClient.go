/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"context"
	"fmt"
	"github.com/dexidp/dex/api"
	"github.com/spf13/cobra"
)

// delClientCmd represents the delClient command
var delClientCmd = &cobra.Command{
	Use:   "del-client",
	Short: "Deletes a client",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("delClient called")
		dexClient, err := newDexClient()
		if err != nil {
			fmt.Printf("Failed to connect to Dex: %v", err)
		}

		clientID, _ := cmd.Flags().GetString("client-id")

		createClientR := &api.DeleteClientReq{
			Id:                   clientID,
		}

		resp, err := dexClient.DeleteClient(context.TODO(), createClientR)
		if err != nil {
			fmt.Printf("failed to create client: %v", err)
		}
		if resp.NotFound == true {
			fmt.Printf("Client with given name doesn't exist.\n")
			return
		}
		fmt.Printf("client %s deleted\n", clientID)
	},
}

func init() {
	rootCmd.AddCommand(delClientCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// delClientCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// delClientCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	delClientCmd.Flags().String("client-id", "", "Client id")
	delClientCmd.MarkFlagRequired("client-id")
}
