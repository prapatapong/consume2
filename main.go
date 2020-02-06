package main

import (
	"consume-repair/repository"
	"consume-repair/service"
)

func main() {
	service.ReadConfig()
	service.SetupConfigToParam()
	repository.SetupConfigToParam()
	service.ConsumeData()

}
