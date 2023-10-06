package introspect

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestColumn_String(t *testing.T) {
	type fields struct {
		ColumnName        string
		Type              string
		TypeId            string
		IsArray           bool
		Nullable          bool
		Generated         bool
		PkName            string
		PkOrdinalPosition int
		JsonType          string
	}

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "column 1",
			fields: fields{
				ColumnName:        "id",
				Type:              "int",
				TypeId:            "23",
				IsArray:           false,
				Nullable:          false,
				Generated:         false,
				PkName:            "id",
				PkOrdinalPosition: 1,
				JsonType:          "identity",
			},
			want: "Column{ColumnName: id, Type: int, TypeId: 23, IsArray: false, Nullable: false, Generated: false, PkName: id, PkOrdinalPosition: 1, JsonType: identity}",
		},
		{
			name: "column 2",
			fields: fields{
				ColumnName:        "id",
				Type:              "text",
				TypeId:            "25",
				IsArray:           false,
				Nullable:          false,
				Generated:         false,
				PkName:            "id",
				PkOrdinalPosition: 1,
				JsonType:          "identity",
			},
			want: "Column{ColumnName: id, Type: text, TypeId: 25, IsArray: false, Nullable: false, Generated: false, PkName: id, PkOrdinalPosition: 1, JsonType: identity}",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			column := &Column{
				ColumnName:        tt.fields.ColumnName,
				Type:              tt.fields.Type,
				TypeId:            tt.fields.TypeId,
				IsArray:           tt.fields.IsArray,
				Nullable:          tt.fields.Nullable,
				Generated:         tt.fields.Generated,
				PkName:            tt.fields.PkName,
				PkOrdinalPosition: tt.fields.PkOrdinalPosition,
				JsonType:          tt.fields.JsonType,
			}

			assert.Equal(t, tt.want, column.String())
		})
	}
}

func TestColumn_Value(t *testing.T) {
	type fields struct {
		ColumnName        string
		Type              string
		TypeId            string
		IsArray           bool
		IsSequence        bool
		Nullable          bool
		Generated         bool
		PkName            string
		PkOrdinalPosition int
		JsonType          string
	}

	testCases := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		{
			name: "column 1",
			fields: fields{
				ColumnName:        "id",
				Type:              "int",
				TypeId:            "23",
				IsArray:           false,
				IsSequence:        true,
				Nullable:          false,
				Generated:         false,
				PkName:            "id",
				PkOrdinalPosition: 1,
				JsonType:          "identity",
			},
			want:    `{"column_name":"id","type":"int","type_id":"23","is_array":false,"is_sequence":true,"nullable":false,"generated":false,"pk_name":"id","pk_ordinal_position":1,"json_type":"identity"}`,
			wantErr: false,
		},
		{
			name: "column 2",
			fields: fields{
				ColumnName:        "id",
				Type:              "text",
				TypeId:            "25",
				IsArray:           false,
				IsSequence:        false,
				Nullable:          false,
				Generated:         false,
				PkName:            "id",
				PkOrdinalPosition: 1,
				JsonType:          "identity",
			},
			want:    `{"column_name":"id","type":"text","type_id":"25","is_array":false,"is_sequence":false,"nullable":false,"generated":false,"pk_name":"id","pk_ordinal_position":1,"json_type":"identity"}`,
			wantErr: false,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			column := &Column{
				ColumnName:        testCase.fields.ColumnName,
				Type:              testCase.fields.Type,
				TypeId:            testCase.fields.TypeId,
				IsArray:           testCase.fields.IsArray,
				IsSequence:        testCase.fields.IsSequence,
				Nullable:          testCase.fields.Nullable,
				Generated:         testCase.fields.Generated,
				PkName:            testCase.fields.PkName,
				PkOrdinalPosition: testCase.fields.PkOrdinalPosition,
				JsonType:          testCase.fields.JsonType,
			}

			got, err := column.Value()

			assert.Nil(t, err)

			assert.Equal(t, testCase.want, string(got.([]byte)))
		})
	}
}
