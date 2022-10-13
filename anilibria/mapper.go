package anilibria

import (
	"pervaki/anilibria/model"
	serviceModel "pervaki/model"
)

func MapClientToService(title model.Title) serviceModel.Title {
	return serviceModel.Title{
		Code:   title.Code,
		NameRu: title.Names.Ru,
	}
}
