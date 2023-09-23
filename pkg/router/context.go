package router

type Context interface {
	Bind(interface{}) error
	BindQuery(interface{}) error
	JSON(int, interface{})
	Query(string) string
	Param(string) string
	GetHeader(string) string
	SetHeader(string, string)
	Next()
	GetClientIP() string
	GetRemoteIP() string
	GetMethod() string
	GetPath() string
}