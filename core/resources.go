package core

type JsonRoot struct {
	Router Router
}

type Router struct {
	Port     string
	Services []Services
}

type Services struct {
	ServicePrefix string
	TargetPath    TargetPath
}

type TargetPath struct {
	Path string
	Auth bool
}
