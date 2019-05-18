package main

import (
	"fmt"
	"net/url"

	"github.com/spf13/cobra"
	"github.com/stellar/go/clients/horizonclient"
)

func createCmdAssets() *cobra.Command {
	return &cobra.Command{
		Use:   "assets",
		Short: "List assets.",
		Run: func(cmd *cobra.Command, args []string) {
			count := 0
			lastCursor := ""
			cursor := ""
			for {
				assets, err := client.Assets(horizonclient.AssetRequest{Cursor: cursor})
				if err != nil {
					fmt.Println(err)
					return
				}
				for _, a := range assets.Embedded.Records {
					count++
					fmt.Printf("%6d: %-12s %s %25s %5d\n", count, a.Code, a.Issuer, a.Amount, a.NumAccounts)
				}
				nextURL, err := url.Parse(assets.Links.Next.Href)
				if err != nil {
					fmt.Println(err)
					return
				}
				cursor = nextURL.Query().Get("cursor")
				if cursor == lastCursor {
					break
				}
			}
		},
	}
}
