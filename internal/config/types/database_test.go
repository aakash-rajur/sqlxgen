package types

import (
	"testing"

	"github.com/aakash-rajur/sqlxgen/internal/utils"
	"github.com/joomcode/errorx"
	"github.com/stretchr/testify/assert"
)

func TestDatabase_Merge(t *testing.T) {
	type fields struct {
		Host     *string
		Port     *string
		Db       *string
		User     *string
		Password *string
		SslMode  *string
		Url      *string
	}

	type args struct {
		other *Database
	}

	testCases := []struct {
		name   string
		fields fields
		args   args
		want   *Database
	}{
		{
			name: "other nil",
			fields: fields{
				Host:     utils.PointerTo("host"),
				Port:     utils.PointerTo("port"),
				Db:       utils.PointerTo("db"),
				User:     utils.PointerTo("user"),
				Password: utils.PointerTo("password"),
				SslMode:  utils.PointerTo("sslmode"),
				Url:      utils.PointerTo("postgres://user:password@host:port/db?sslmode=sslmode"),
			},
			args: args{
				other: nil,
			},
			want: &Database{
				Host:     utils.PointerTo("host"),
				Port:     utils.PointerTo("port"),
				Db:       utils.PointerTo("db"),
				User:     utils.PointerTo("user"),
				Password: utils.PointerTo("password"),
				SslMode:  utils.PointerTo("sslmode"),
				Url:      utils.PointerTo("postgres://user:password@host:port/db?sslmode=sslmode"),
			},
		},
		{
			name: "other host and port",
			fields: fields{
				Host:     utils.PointerTo("host"),
				Port:     utils.PointerTo("port"),
				Db:       utils.PointerTo("db"),
				User:     utils.PointerTo("user"),
				Password: utils.PointerTo("password"),
				SslMode:  utils.PointerTo("sslmode"),
				Url:      utils.PointerTo("postgres://user:password@host:port/db?sslmode=sslmode"),
			},
			args: args{
				other: &Database{
					Host: utils.PointerTo("other-host"),
					Port: utils.PointerTo("other-port"),
				},
			},
			want: &Database{
				Host:     utils.PointerTo("other-host"),
				Port:     utils.PointerTo("other-port"),
				Db:       utils.PointerTo("db"),
				User:     utils.PointerTo("user"),
				Password: utils.PointerTo("password"),
				SslMode:  utils.PointerTo("sslmode"),
				Url:      utils.PointerTo("postgres://user:password@host:port/db?sslmode=sslmode"),
			},
		},
		{
			name: "other all",
			fields: fields{
				Host:     utils.PointerTo("host"),
				Port:     utils.PointerTo("port"),
				Db:       utils.PointerTo("db"),
				User:     utils.PointerTo("user"),
				Password: utils.PointerTo("password"),
				SslMode:  utils.PointerTo("sslmode"),
				Url:      utils.PointerTo("postgres://user:password@host:port/db?sslmode=sslmode"),
			},
			args: args{
				other: &Database{
					Host:     utils.PointerTo("other-host"),
					Port:     utils.PointerTo("other-port"),
					Db:       utils.PointerTo("other-db"),
					User:     utils.PointerTo("other-user"),
					Password: utils.PointerTo("other-password"),
					SslMode:  utils.PointerTo("other-sslmode"),
					Url:      utils.PointerTo("postgres://other-user:other-password@other-host:other-port/other-db?sslmode=other-sslmode"),
				},
			},
			want: &Database{
				Host:     utils.PointerTo("other-host"),
				Port:     utils.PointerTo("other-port"),
				Db:       utils.PointerTo("other-db"),
				User:     utils.PointerTo("other-user"),
				Password: utils.PointerTo("other-password"),
				SslMode:  utils.PointerTo("other-sslmode"),
				Url:      utils.PointerTo("postgres://other-user:other-password@other-host:other-port/other-db?sslmode=other-sslmode"),
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			d := &Database{
				Host:     testCase.fields.Host,
				Port:     testCase.fields.Port,
				Db:       testCase.fields.Db,
				User:     testCase.fields.User,
				Password: testCase.fields.Password,
				SslMode:  testCase.fields.SslMode,
				Url:      testCase.fields.Url,
			}

			got := d.Merge(testCase.args.other)

			assert.Equal(t, testCase.want, got)
		})
	}
}

func TestDatabase_String(t *testing.T) {
	type fields struct {
		Host     *string
		Port     *string
		Db       *string
		User     *string
		Password *string
		SslMode  *string
		Url      *string
	}

	testCases := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "all",
			fields: fields{
				Host:     utils.PointerTo("host"),
				Port:     utils.PointerTo("port"),
				Db:       utils.PointerTo("db"),
				User:     utils.PointerTo("user"),
				Password: utils.PointerTo("password"),
				SslMode:  utils.PointerTo("sslmode"),
				Url:      utils.PointerTo("postgres://user:password@host:port/db?sslmode=sslmode"),
			},
			want: "Database{Host: host, Port: port, Db: db, User: user, Password: password, SslMode: sslmode, Url: postgres://user:password@host:port/db?sslmode=sslmode}",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			d := &Database{
				Host:     testCase.fields.Host,
				Port:     testCase.fields.Port,
				Db:       testCase.fields.Db,
				User:     testCase.fields.User,
				Password: testCase.fields.Password,
				SslMode:  testCase.fields.SslMode,
				Url:      testCase.fields.Url,
			}

			got := d.String()

			assert.Equal(t, testCase.want, got)
		})
	}
}

func TestDatabase_GetUrl(t *testing.T) {
	type fields struct {
		Host     *string
		Port     *string
		Db       *string
		User     *string
		Password *string
		SslMode  *string
		Url      *string
	}

	type args struct {
		engine string
	}

	testCases := []struct {
		name   string
		fields fields
		args   args
		want   string
		err    error
	}{
		{
			name: "mysql",
			fields: fields{
				Host:     utils.PointerTo("host"),
				Port:     utils.PointerTo("3306"),
				Db:       utils.PointerTo("db"),
				User:     utils.PointerTo("user"),
				Password: utils.PointerTo("password"),
				SslMode:  utils.PointerTo("sslmode"),
				Url:      utils.PointerTo("root:@localhost:3306/mysql"),
			},
			args: args{
				engine: "mysql",
			},
			want: "user:password@tcp(host:3306)/db?parseTime=true",
			err:  nil,
		},
		{
			name: "postgres",
			fields: fields{
				Host:     utils.PointerTo("host"),
				Port:     utils.PointerTo("5432"),
				Db:       utils.PointerTo("db"),
				User:     utils.PointerTo("user"),
				Password: utils.PointerTo("password"),
				SslMode:  utils.PointerTo("sslmode"),
				Url:      utils.PointerTo("postgres://user:password@host:5432/db?sslmode=sslmode"),
			},
			args: args{
				engine: "postgres",
			},
			want: "postgres://user:password@host:5432/db?sslmode=sslmode",
			err:  nil,
		},
		{
			name: "invalid engine",
			fields: fields{
				Host:     utils.PointerTo("host"),
				Port:     utils.PointerTo("port"),
				Db:       utils.PointerTo("db"),
				User:     utils.PointerTo("user"),
				Password: utils.PointerTo("password"),
				SslMode:  utils.PointerTo("sslmode"),
				Url:      utils.PointerTo("postgres://user:password@host:port/db?sslmode=sslmode"),
			},
			args: args{
				engine: "invalid",
			},
			want: "",
			err:  errorx.IllegalArgument.New("invalid engine: %s", "invalid"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			d := &Database{
				Host:     testCase.fields.Host,
				Port:     testCase.fields.Port,
				Db:       testCase.fields.Db,
				User:     testCase.fields.User,
				Password: testCase.fields.Password,
				SslMode:  testCase.fields.SslMode,
				Url:      testCase.fields.Url,
			}

			got, err := d.GetUrl(testCase.args.engine)

			if testCase.err != nil {
				leftErr := testCase.err.Error()

				rightErr := err.Error()

				assert.Containsf(t, rightErr, leftErr, "expected error to contain %s", leftErr)
			} else {
				assert.Nil(t, err)

				assert.Equal(t, testCase.want, got)
			}
		})
	}
}
