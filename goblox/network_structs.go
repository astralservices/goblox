package goblox

// HTTPRequestType : Type of HTTP Request to send
type HTTPRequestType int

// HTTPContentType : HTTP Content Types
type HTTPContentType string

// HTTPRequestType
const (
	GET  HTTPRequestType = 0
	POST HTTPRequestType = 1
)

// HTTPContentType
const (
	APPJSON HTTPContentType = "application/json"
	TXTJSON HTTPContentType = "text/json"
)
