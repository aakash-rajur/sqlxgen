package types

import (
	"testing"

	"github.com/aakash-rajur/sqlxgen/internal/utils"
	"github.com/joomcode/errorx"
	"github.com/stretchr/testify/assert"
)

func TestGetPgUrl(t *testing.T) {
	type args struct {
		host     *string
		port     *string
		db       *string
		user     *string
		password *string
		sslMode  *string
		url      *string
	}

	testCases := []struct {
		name string
		args args
		want string
	}{
		{
			name: "all",
			args: args{
				host:     utils.PointerTo("host"),
				port:     utils.PointerTo("port"),
				db:       utils.PointerTo("db"),
				user:     utils.PointerTo("user"),
				password: utils.PointerTo("password"),
				sslMode:  utils.PointerTo("sslmode"),
				url:      nil,
			},
			want: "postgres://user:password@host:port/db?sslmode=sslmode",
		},
		{
			name: "no port",
			args: args{
				host:     utils.PointerTo("host"),
				port:     nil,
				db:       utils.PointerTo("db"),
				user:     utils.PointerTo("user"),
				password: utils.PointerTo("password"),
				sslMode:  utils.PointerTo("sslmode"),
				url:      nil,
			},
			want: "postgres://user:password@host:5432/db?sslmode=sslmode",
		},
		{
			name: "no db",
			args: args{
				host:     utils.PointerTo("host"),
				port:     utils.PointerTo("port"),
				db:       nil,
				user:     utils.PointerTo("user"),
				password: utils.PointerTo("password"),
				sslMode:  utils.PointerTo("sslmode"),
				url:      nil,
			},
			want: "postgres://user:password@host:port/postgres?sslmode=sslmode",
		},
		{
			name: "no user",
			args: args{
				host:     utils.PointerTo("host"),
				port:     utils.PointerTo("port"),
				db:       utils.PointerTo("db"),
				user:     nil,
				password: utils.PointerTo("password"),
				sslMode:  utils.PointerTo("sslmode"),
				url:      nil,
			},
			want: "postgres://postgres:password@host:port/db?sslmode=sslmode",
		},
		{
			name: "no password",
			args: args{
				host:     utils.PointerTo("host"),
				port:     utils.PointerTo("port"),
				db:       utils.PointerTo("db"),
				user:     utils.PointerTo("user"),
				password: nil,
				sslMode:  utils.PointerTo("sslmode"),
				url:      nil,
			},
			want: "postgres://user:@host:port/db?sslmode=sslmode",
		},
		{
			name: "no sslmode",
			args: args{
				host:     utils.PointerTo("host"),
				port:     utils.PointerTo("port"),
				db:       utils.PointerTo("db"),
				user:     utils.PointerTo("user"),
				password: utils.PointerTo("password"),
				sslMode:  nil,
				url:      nil,
			},
			want: "postgres://user:password@host:port/db?sslmode=disable",
		},
		{
			name: "url with host",
			args: args{
				host:     utils.PointerTo("host"),
				port:     utils.PointerTo("port"),
				db:       utils.PointerTo("db"),
				user:     utils.PointerTo("user"),
				password: utils.PointerTo("password"),
				sslMode:  utils.PointerTo("sslmode"),
				url:      utils.PointerTo("postgres://user1:password1@host1:port1/db?sslmode=sslmode1"),
			},
			want: "postgres://user:password@host:port/db?sslmode=sslmode",
		},
		{
			name: "only url",
			args: args{
				host:     nil,
				port:     nil,
				db:       nil,
				user:     nil,
				password: nil,
				sslMode:  nil,
				url:      utils.PointerTo("postgres://user1:password1@host1:port1/db?sslmode=sslmode1"),
			},
			want: "postgres://user1:password1@host1:port1/db?sslmode=sslmode1",
		},
		{
			name: "no args",
			args: args{
				host:     nil,
				port:     nil,
				db:       nil,
				user:     nil,
				password: nil,
				sslMode:  nil,
				url:      nil,
			},
			want: "postgres://postgres:@localhost:5432/postgres?sslmode=disable",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			db := &Database{
				Host:     testCase.args.host,
				Port:     testCase.args.port,
				Db:       testCase.args.db,
				User:     testCase.args.user,
				Password: testCase.args.password,
				SslMode:  testCase.args.sslMode,
				Url:      testCase.args.url,
			}

			got := getPgUrl(db)

			assert.Equal(t, testCase.want, got)
		})
	}
}

func TestGetMysqlUrl(t *testing.T) {
	type args struct {
		host     *string
		port     *string
		db       *string
		user     *string
		password *string
		sslMode  *string
		url      *string
	}

	testCases := []struct {
		name string
		args args
		want string
	}{
		{
			name: "all",
			args: args{
				host:     utils.PointerTo("host"),
				port:     utils.PointerTo("port"),
				db:       utils.PointerTo("db"),
				user:     utils.PointerTo("user"),
				password: utils.PointerTo("password"),
				sslMode:  nil,
				url:      nil,
			},
			want: "user:password@tcp(host:port)/db?parseTime=true",
		},
		{
			name: "no port",
			args: args{
				host:     utils.PointerTo("host"),
				port:     nil,
				db:       utils.PointerTo("db"),
				user:     utils.PointerTo("user"),
				password: utils.PointerTo("password"),
				sslMode:  nil,
				url:      nil,
			},
			want: "user:password@tcp(host:3306)/db?parseTime=true",
		},
		{
			name: "no db",
			args: args{
				host:     utils.PointerTo("host"),
				port:     utils.PointerTo("port"),
				db:       nil,
				user:     utils.PointerTo("user"),
				password: utils.PointerTo("password"),
				sslMode:  nil,
				url:      nil,
			},
			want: "user:password@tcp(host:port)/mysql?parseTime=true",
		},
		{
			name: "no user",
			args: args{
				host:     utils.PointerTo("host"),
				port:     utils.PointerTo("port"),
				db:       utils.PointerTo("db"),
				user:     nil,
				password: utils.PointerTo("password"),
				sslMode:  nil,
				url:      nil,
			},
			want: "root:password@tcp(host:port)/db?parseTime=true",
		},
		{
			name: "no password",
			args: args{
				host:     utils.PointerTo("host"),
				port:     utils.PointerTo("port"),
				db:       utils.PointerTo("db"),
				user:     utils.PointerTo("user"),
				password: nil,
				sslMode:  nil,
				url:      nil,
			},
			want: "user:@tcp(host:port)/db?parseTime=true",
		},
		{
			name: "url with host",
			args: args{
				host:     utils.PointerTo("host"),
				port:     utils.PointerTo("port"),
				db:       utils.PointerTo("db"),
				user:     utils.PointerTo("user"),
				password: utils.PointerTo("password"),
				sslMode:  nil,
				url:      utils.PointerTo("user1:password1@tcp(host1:port1)/db1?parseTime=true"),
			},
			want: "user:password@tcp(host:port)/db?parseTime=true",
		},
		{
			name: "url without host",
			args: args{
				host:     nil,
				port:     nil,
				db:       nil,
				user:     nil,
				password: nil,
				sslMode:  nil,
				url:      utils.PointerTo("user1:password1@tcp(host1:port1)/db1?parseTime=true"),
			},
			want: "user1:password1@tcp(host1:port1)/db1?parseTime=true",
		},
		{
			name: "no args",
			args: args{
				host:     nil,
				port:     nil,
				db:       nil,
				user:     nil,
				password: nil,
				sslMode:  nil,
				url:      nil,
			},
			want: "root:@tcp(localhost:3306)/mysql?parseTime=true",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			db := &Database{
				Host:     testCase.args.host,
				Port:     testCase.args.port,
				Db:       testCase.args.db,
				User:     testCase.args.user,
				Password: testCase.args.password,
				SslMode:  testCase.args.sslMode,
				Url:      testCase.args.url,
			}

			got := getMysqlUrl(db)

			assert.Equal(t, testCase.want, got)
		})
	}
}

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
