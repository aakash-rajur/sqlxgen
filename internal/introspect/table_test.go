package introspect

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTable_String(t *testing.T) {
	testCases := []struct {
		name  string
		table Table
		want  string
	}{
		{
			name: "empty",
			table: Table{
				SchemaName: "",
				TableName:  "",
				Columns:    []Column{},
			},
			want: "Table{SchemaName: , TableName: , Columns: []}",
		},
		{
			name: "single column",
			table: Table{
				SchemaName: "schema",
				TableName:  "table",
				Columns: []Column{
					{
						ColumnName:        "column",
						Type:              "text",
						TypeId:            "0",
						IsArray:           false,
						Nullable:          false,
						Generated:         false,
						PkName:            "pk_name",
						PkOrdinalPosition: 0,
						JsonType:          "identity",
					},
				},
			},
			want: "Table{SchemaName: schema, TableName: table, Columns: [Column{ColumnName: column, Type: text, TypeId: 0, IsArray: false, Nullable: false, Generated: false, PkName: pk_name, PkOrdinalPosition: 0, JsonType: identity}]}",
		},

		{
			name: "multiple columns",
			table: Table{
				SchemaName: "schema",
				TableName:  "table",
				Columns: []Column{
					{
						ColumnName:        "column1",
						Type:              "text",
						TypeId:            "0",
						IsArray:           false,
						Nullable:          false,
						Generated:         false,
						PkName:            "pk_name",
						PkOrdinalPosition: 0,
					},
					{
						ColumnName:        "column2",
						Type:              "text",
						TypeId:            "0",
						IsArray:           false,
						Nullable:          false,
						Generated:         false,
						PkName:            "pk_name",
						PkOrdinalPosition: 0,
					},
				},
			},
			want: "Table{SchemaName: schema, TableName: table, Columns: [Column{ColumnName: column1, Type: text, TypeId: 0, IsArray: false, Nullable: false, Generated: false, PkName: pk_name, PkOrdinalPosition: 0, JsonType: }, Column{ColumnName: column2, Type: text, TypeId: 0, IsArray: false, Nullable: false, Generated: false, PkName: pk_name, PkOrdinalPosition: 0, JsonType: }]}",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			actual := testCase.table.String()

			assert.Equal(t, testCase.want, actual)
		})
	}
}

func TestTable_PrimaryKey(t *testing.T) {
	testCases := []struct {
		name  string
		table Table
		want  Columns
	}{
		{
			name: "empty",
			table: Table{
				SchemaName: "",
				TableName:  "",
				Columns:    []Column{},
			},
			want: []Column{},
		},
		{
			name: "single column",
			table: Table{
				SchemaName: "schema",
				TableName:  "table",
				Columns: []Column{
					{
						ColumnName:        "column",
						Type:              "text",
						TypeId:            "0",
						IsArray:           false,
						Nullable:          false,
						Generated:         false,
						PkName:            "pk_name",
						PkOrdinalPosition: 1,
						JsonType:          "identity",
					},
				},
			},
			want: []Column{
				{
					ColumnName:        "column",
					Type:              "text",
					TypeId:            "0",
					IsArray:           false,
					Nullable:          false,
					Generated:         false,
					PkName:            "pk_name",
					PkOrdinalPosition: 1,
					JsonType:          "identity",
				},
			},
		},
		{
			name: "multiple columns",
			table: Table{
				SchemaName: "schema",
				TableName:  "table",
				Columns: []Column{
					{
						ColumnName:        "column1",
						Type:              "text",
						TypeId:            "0",
						IsArray:           false,
						Nullable:          false,
						Generated:         false,
						PkName:            "pk_name",
						PkOrdinalPosition: 1,
					},
					{
						ColumnName:        "column2",
						Type:              "text",
						TypeId:            "0",
						IsArray:           false,
						Nullable:          false,
						Generated:         false,
						PkName:            "pk_name",
						PkOrdinalPosition: 2,
					},
				},
			},
			want: []Column{
				{
					ColumnName:        "column1",
					Type:              "text",
					TypeId:            "0",
					IsArray:           false,
					Nullable:          false,
					Generated:         false,
					PkName:            "pk_name",
					PkOrdinalPosition: 1,
				},
				{
					ColumnName:        "column2",
					Type:              "text",
					TypeId:            "0",
					IsArray:           false,
					Nullable:          false,
					Generated:         false,
					PkName:            "pk_name",
					PkOrdinalPosition: 2,
				},
			},
		},
		{
			name: "multiple columns with no primary key",
			table: Table{
				SchemaName: "schema",
				TableName:  "table",
				Columns: []Column{
					{
						ColumnName:        "column1",
						Type:              "text",
						TypeId:            "0",
						IsArray:           false,
						Nullable:          false,
						Generated:         false,
						PkName:            "",
						PkOrdinalPosition: 0,
					},
					{
						ColumnName:        "column2",
						Type:              "text",
						TypeId:            "0",
						IsArray:           false,
						Nullable:          false,
						Generated:         false,
						PkName:            "",
						PkOrdinalPosition: 0,
					},
				},
			},
			want: []Column{
				{
					ColumnName:        "column1",
					Type:              "text",
					TypeId:            "0",
					IsArray:           false,
					Nullable:          false,
					Generated:         false,
					PkName:            "",
					PkOrdinalPosition: 0,
				},
				{
					ColumnName:        "column2",
					Type:              "text",
					TypeId:            "0",
					IsArray:           false,
					Nullable:          false,
					Generated:         false,
					PkName:            "",
					PkOrdinalPosition: 0,
				},
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			actual := testCase.table.PrimaryKey()

			assert.Equal(t, testCase.want, actual)
		})
	}
}
