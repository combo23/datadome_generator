package datadome

var possibleDvm = []int{4, 8, 16, 32}
var possiblePixelRatios = []int{1, 2}
var possibleHc = []int{4, 6, 8, 10, 12, 16, 32, 48}
var possibleGlvd = []string{
	"Google Inc. (AMD)",
	"Google Inc. (NVIDIA)",
	"Google Inc. (Apple)",
}

var possibleOrientationTypes = []string{"portrait-primary", "landscape-primary", "portrait-secondary", "landscape-secondary"}

var possibleRs_cd = []int{1, 4, 8, 15, 16, 24, 32, 48}
var possibleGlrd = map[string][]string{
	"Google Inc. (AMD)": {
		"ANGLE (AMD, AMD Radeon RX 6500 XT Direct3D11 vs_5_0 ps_5_0, D3D11)",
		"ANGLE (AMD, AMD Radeon RX 6600 XT Direct3D11 vs_5_0 ps_5_0, D3D11)",
		"ANGLE (AMD, AMD Radeon RX 6700 XT Direct3D11 vs_5_0 ps_5_0, D3D11)",
		"ANGLE (AMD, AMD Radeon RX 6800 XT Direct3D11 vs_5_0 ps_5_0, D3D11)",
		"ANGLE (AMD, AMD Radeon RX 6900 XT Direct3D11 vs_5_0 ps_5_0, D3D11)",
		"ANGLE (AMD, AMD Radeon(TM) RX Vega 11 Graphics Direct3D11 vs_5_0 ps_5_0, D3D11)",
		"ANGLE (AMD, AMD Radeon(TM) Graphics Direct3D11 vs_5_0 ps_5_0, D3D11)",
		"ANGLE (AMD, AMD Radeon RX 6700 Pro Direct3D11 vs_5_0 ps_5_0, D3D11)",
		"ANGLE (AMD, AMD Radeon RX 6600 Direct3D11 vs_5_0 ps_5_0, D3D11)",
	},
	"Google Inc. (NVIDIA)": {
		"ANGLE (NVIDIA, NVIDIA GeForce RTX 3090 Direct3D11 vs_5_0 ps_5_0, D3D11)",
		"ANGLE (NVIDIA, NVIDIA GeForce RTX 3080 Direct3D11 vs_5_0 ps_5_0, D3D11)",
		"ANGLE (NVIDIA, NVIDIA GeForce RTX 3070 Direct3D11 vs_5_0 ps_5_0, D3D11)",
		"ANGLE (NVIDIA, NVIDIA GeForce RTX 3060 Ti Direct3D11 vs_5_0 ps_5_0, D3D11)",
		"ANGLE (NVIDIA, NVIDIA GeForce RTX 3060 Direct3D11 vs_5_0 ps_5_0, D3D11)",
		"ANGLE (NVIDIA, NVIDIA GeForce GTX 1660 SUPER Direct3D11 vs_5_0 ps_5_0, D3D11)",
		"ANGLE (NVIDIA, NVIDIA GeForce GTX 1060 6GB Direct3D11 vs_5_0 ps_5_0, D3D11-30.0.14.9613)",
		"ANGLE (NVIDIA, NVIDIA GeForce GTX 1050 Ti Direct3D11 vs_5_0 ps_5_0, D3D11)",
		"ANGLE (NVIDIA, NVIDIA GeForce GTX 1050 Direct3D11 vs_5_0 ps_5_0, D3D11-27.21.14.5671)",
		"ANGLE (NVIDIA, NVIDIA GeForce GTX 1660 Direct3D11 vs_5_0 ps_5_0, D3D11)",
		"ANGLE (NVIDIA, NVIDIA GeForce GTX 1650 Super Direct3D11 vs_5_0 ps_5_0, D3D11)",
	},
	"Google Inc. (Apple)": {
		"ANGLE (Apple, Apple M2, OpenGL 4.1)",
		"ANGLE (Apple, Apple M1 Pro, OpenGL 4.1)",
		"ANGLE (Apple, Apple M1, OpenGL 4.1)",
	},
}

var resolutions = []Resolution{
	{
		Width:  2560,
		Height: 1440,
	},
	{
		Width:  1920,
		Height: 1080,
	},
	{
		Width:  1280,
		Height: 720,
	},
	{
		Width:  1024,
		Height: 768,
	},
	{
		Width:  1280,
		Height: 960,
	},
	{
		Width:  1280,
		Height: 1024,
	},
	{
		Width:  1400,
		Height: 1050,
	},
	{
		Width:  1600,
		Height: 1200,
	},
	{
		Width:  4096,
		Height: 2160,
	},
	{
		Width:  1366,
		Height: 768,
	},
	{
		Width:  3200,
		Height: 1800,
	},
	{
		Width:  3840,
		Height: 2160,
	},
	{
		Width:  1600,
		Height: 900,
	},
}
