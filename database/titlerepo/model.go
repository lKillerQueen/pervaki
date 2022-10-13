package titlerepo

type Title struct {
	Code   string `db:"code"`
	NameRu string `db:"name_ru"`
}
