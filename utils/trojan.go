package utils

import (
	"regexp"

	"github.com/micro/go-micro/registry"
)

// TrojanInfoHandler 处理Torjan微服务并返回信息
func TrojanInfoHandler(value *registry.Service) (string, string, int) {
	// fmt.Printf("%#v\n", service)
	// fmt.Printf("%v - %v - %v - %v\n", service.Name(), service.Options(), service.String(), service.Server().Options().Metadata)
	// regexp
	re, _ := regexp.Compile("^\\w{24}")
	if !re.MatchString(value.Name) {
		return value.Name, "unkwon_service", 1
	}

	return value.Name, "todo", 0 // 目前仅需返回ID即可
}
