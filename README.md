# eol

`eol` stands for "End of Life". It is a CLI tool to query the endoflife.date API, which provides information about the end of life of various software.

## Commands and subcommands

### `products`

The `products` command is used to query the endoflife.date API and retrieve the list of all the products it supports. The list is obtained from the endoflife.date API `All Products` endpoint.

```bash
eol products
```

### `get`

The `get` command is used to query the endoflife.date API and retrieve the end of life for all the releases for a specific software, identified by its product name. The product name is passed in the `--product` flag, and it's obtained from the endoflife.date API `All Products` endpoint.

```bash
eol get --product "go"
```
