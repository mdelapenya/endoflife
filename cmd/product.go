package cmd

import (
	"fmt"
	"os"

	"github.com/mdelapenya/endoflife/internal"
	"github.com/mdelapenya/endoflife/types"
	"github.com/spf13/cobra"
	"github.com/tidwall/gjson"
)

var product string

func init() {
	productCmd.Flags().StringVarP(&product, "product", "p", "", "Product to query (required)")

	rootCmd.AddCommand(productCmd)
}

var productCmd = &cobra.Command{
	Use:   "get",
	Short: "Retrieves the product information from the endoflife.date API",
	Long:  `Retrieves the product information from the endoflife.date API`,
	PreRun: func(cmd *cobra.Command, args []string) {
		if product == "" {
			fmt.Println(">> product is required")
			os.Exit(1)
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		getProduct(product)
	},
}

func getProduct(product string) {
	result, err := internal.GetProduct(product)
	if err != nil {
		fmt.Println(">> error querying endoflife.date API:", err)
		os.Exit(1)
	}

	var products []types.Product

	result.ForEach(func(key, value gjson.Result) bool {
		cycle := value.Get("cycle")
		latest := value.Get("latest")
		eol := value.Get("eol")
		releaseDate := value.Get("releaseDate")
		latestReleaseDate := value.Get("latestReleaseDate")
		lts := value.Get("lts")

		product := types.Product{
			Cycle:             cycle.String(),
			Latest:            latest.String(),
			Eol:               eol.String(),
			ReleaseDate:       releaseDate.String(),
			LatestReleaseDate: latestReleaseDate.String(),
			Lts:               lts.Bool(),
		}

		products = append(products, product)
		return true
	})

	fmt.Printf("%s versions:\n %+v\n", product, products)
}
