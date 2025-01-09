package main

import (
	"fmt"

	serv_dto "github.com/root9464/Ton-students/module/service_module/dto"
	serv_model "github.com/root9464/Ton-students/module/service_module/model"
	"github.com/root9464/Ton-students/shared/utils"
)

func main() {
	service := &serv_dto.ServiceType{
		UserId: 12345,
		Price:  99.99,
		Infos: []serv_dto.InfosType{
			{
				Title:   "Service Title 1",
				Content: "Detailed content for service 1.",
			},
			{
				Title:   "Service Title 2",
				Content: "Detailed content for service 2.",
			},
		},
		Tags: &[]serv_dto.TagsType{
			{
				ServiceId: "service-123",
				Name:      "Tag1",
			},
			{
				ServiceId: "service-123",
				Name:      "Tag2",
			},
		},
	}

	convert, _ := utils.ConvertDtoToEntity[serv_model.Service](service)
	fmt.Printf("%+v", convert)
}
