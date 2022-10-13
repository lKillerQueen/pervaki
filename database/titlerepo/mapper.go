package titlerepo

import "pervaki/model"

func MapServiceToDb(title model.Title) Title {
	return Title{
		Code:   title.Code,
		NameRu: title.NameRu,
	}
}
