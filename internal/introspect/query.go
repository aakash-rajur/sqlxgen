package introspect

type Query struct {
	Filename   string  `db:"filename" json:"filename"`
	QueryName  string  `db:"query_name" json:"query_name"`
	PascalName string  `db:"pascal_name" json:"pascal_name"`
	CamelName  string  `db:"camel_name" json:"camel_name"`
	Columns    Columns `db:"columns" json:"columns"`
	Params     Columns `db:"params" json:"params"`
	SourceDir  string  `db:"source_dir" json:"source_dir"`
}
