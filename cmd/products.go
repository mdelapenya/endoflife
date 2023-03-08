package cmd

import (
	"fmt"
	"os"

	"github.com/mdelapenya/endoflife/internal"
	"github.com/spf13/cobra"
	"github.com/tidwall/gjson"
)

func init() {
	rootCmd.AddCommand(productsCmd)
}

var productsCmd = &cobra.Command{
	Use:   "products",
	Short: "Retrieves the list of products from endoflife.date API",
	Long:  `Retrieves the list of products from endoflife.date API`,
	Run: func(cmd *cobra.Command, args []string) {
		getProducts()
	},
}

func getProducts() {
	result, err := internal.GetProducts()
	if err != nil {
		fmt.Println(">> error querying endoflife.date API:", err)
		os.Exit(1)
	}

	var products []string

	result.ForEach(func(key, value gjson.Result) bool {
		product := value.String()

		products = append(products, product)
		return true
	})

	fmt.Printf("%s products:\n %+v\n", product, products)
}
