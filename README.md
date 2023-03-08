# eol

`eol` stands for "End of Life".

It is a CLI tool to query the [endoflife.date API][endoflife_api_url], which provides information about the end of life of various software.

## Commands and subcommands

### `products`

The `products` command is used to query the [endoflife.date API][endoflife_api_url] and retrieve the list of all the products it supports. The list is obtained from the endoflife.date API `All Products` endpoint.

```bash
eol products
```

### `get`

The `get` command is used to query the [endoflife.date API][endoflife_api_url] and retrieve the end of life for all the releases for a specific software, identified by its product name. The product name is passed in the `--product` flag, and it's obtained from the endoflife.date API `All Products` endpoint.

```bash
eol get --product "go"
```

If you do not know the name of the product you are looking for, you can use an interactive mode that will retrieve all the products from the [endoflife.date API][endoflife_api_url], presenting a prompt to select one of them.

```bash
eol get --interactive
```

[endoflife_api_url]: https://endoflife.date/docs/api
