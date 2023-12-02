package constants

type SiteInfo struct {
	DDVersion  string
	DDK        string
	DDEndpoint string
}

var SupportedSites = map[string]SiteInfo{
	"footlocker": {
		DDVersion:  "4.19.0",
		DDK:        "A55FBF4311ED6F1BF9911EB71931D5",
		DDEndpoint: "https://api-js.datadome.co/js/",
	},
	"tickets.mancity": {
		DDVersion:  "4.19.0",
		DDK:        "60D428DD4BC75DF55D205B3DBE4AFF",
		DDEndpoint: "https://api-js.datadome.co/js/",
	},
}
