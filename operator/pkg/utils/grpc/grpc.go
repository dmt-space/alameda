package grpc

import (
	"flag"
)

func GetAIServiceAddress() string {
	return flag.Lookup("ai-server").Value.(flag.Getter).Get().(string)
}

func GetServerPort() int {
	return flag.Lookup("server-port").Value.(flag.Getter).Get().(int)
}

func GetPrometheusBearerTokenFile() string {
	return flag.Lookup("prometheus-bearer-token-file").Value.(flag.Getter).Get().(string)
}
