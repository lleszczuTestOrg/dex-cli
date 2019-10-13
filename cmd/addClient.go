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
	"log"

	"github.com/spf13/cobra"
	"github.com/google/uuid"
)

// addClientCmd represents the addClient command
var addClientCmd = &cobra.Command{
	Use:   "add-client",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		dexClient, err := newDexClient()
		if err != nil {
			fmt.Printf("Failed to connect to Dex: %v", err)
		}

		clientSecret := uuid.New().String()
		clientID, _ := cmd.Flags().GetString("client-id")
		uris, _ := cmd.Flags().GetStringSlice("redirect-uris")
		public, _ := cmd.Flags().GetBool("public")

		oidcClient := api.Client{
			Id:           clientID,
			Secret:       clientSecret,
			RedirectUris: uris,
			TrustedPeers: nil,
			Public:       public,
			Name:         clientID,
			LogoUrl:      "",
		}
		createClientR := &api.CreateClientReq{
			Client: &oidcClient,
		}

		resp, err := dexClient.CreateClient(context.TODO(), createClientR)
		if err != nil {
			fmt.Printf("failed to create client: %v", err)
		}
		if resp.AlreadyExists == true {
			fmt.Printf("Client with given name already exists.\n")
			return
		}
		log.Printf("created  client: %v", resp)
	},
}

func init() {
	rootCmd.AddCommand(addClientCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addClientCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	//addClientCmd.Flags().BoolP("client-id", "t", false, "Help message for toggle")
	addClientCmd.Flags().String("client-id", "", "Client id")
	addClientCmd.Flags().StringSlice("redirect-uris", []string{}, "List of redirect uris")
	addClientCmd.MarkFlagRequired("client-id")
	addClientCmd.MarkFlagRequired("redirect-uris")


}
