package servers

type IModuleFactory interface {
	MonitorModule()
}

type moduleFactory struct {
	server *server
}

// func NewModuleFactory(server *server) IModuleFactory {
// 	return &moduleFactory{
// 		server: server,
// 	}
// }

