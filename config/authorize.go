package config

type Whitelist struct {
	Url string
	Act string
}

// 白名单
var Whitelists = []Whitelist{
	{Url: "/login", Act: "POST"},
	{Url: "/register", Act: "POST"},
	{Url: "/init", Act: "GET"},
	{Url: "/refresh-permissions", Act: "GET"},
}
