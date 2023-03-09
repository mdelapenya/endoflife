package cmd

import (
	"fmt"
	"os"

	survey "github.com/AlecAivazis/survey/v2"
	"github.com/mdelapenya/endoflife/internal"
	"github.com/spf13/cobra"
	"github.com/tidwall/gjson"
)

var product string
var interactive bool

func init() {
	productCmd.Flags().StringVarP(&product, "product", "p", "", "Product to query (required)")
	productCmd.Flags().BoolVarP(&interactive, "interactive", "i", false, "Interactive mode to select the product from a list (optional)")

	rootCmd.AddCommand(productCmd)
}

var productCmd = &cobra.Command{
	Use:   "get",
	Short: "Retrieves the product information from the endoflife.date API",
	Long:  `Retrieves the product information from the endoflife.date API`,
	PreRun: func(cmd *cobra.Command, args []string) {
		if interactive {
			return
		}

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
	selProduct := product
	var err error

	if interactive {
		selProduct, err = getProductInteractive()
		if err != nil {
			fmt.Println(">> error selecting product from endoflife.date API:", err)
			os.Exit(1)
		}
	}

	result, err := internal.GetProduct(selProduct)
	if err != nil {
		fmt.Println(">> error querying endoflife.date API:", err)
		os.Exit(1)
	}

	fmt.Printf("%s versions:\n %+v\n", product, result)
}

func getProductInteractive() (string, error) {
	result, err := internal.GetProducts()
	if err != nil {
		return "", err
	}

	var products []string

	result.ForEach(func(key, value gjson.Result) bool {
		product := value.String()

		products = append(products, product)
		return true
	})

	var selectedProduct string
	survey.AskOne(&survey.Select{
		Message: "Please select the product to check its end of life:",
		Options: products,
	}, &selectedProduct, survey.WithRemoveSelectAll())

	return selectedProduct, nil
}
