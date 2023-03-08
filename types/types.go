package types

type Product struct {
	Cycle             string `json:"cycle"`
	Eol               string `json:"eol"`
	Latest            string `json:"latest"`
	ReleaseDate       string `json:"releaseDate"`
	LatestReleaseDate string `json:"latestReleaseDate"`
	Lts               bool   `json:"lts"`
}
