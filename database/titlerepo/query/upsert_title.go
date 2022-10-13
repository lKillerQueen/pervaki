package query

const UpsertTitleSql = `
	insert into title (code, name_ru) 
	values ($1, $2)
	on conflict (code) do update set name_ru = excluded.name_ru;
`
