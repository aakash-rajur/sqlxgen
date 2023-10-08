package types

import (
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/joomcode/errorx"
)

type Database struct {
	Host     *string `json:"host" yaml:"host"`
	Port     *string `json:"port" yaml:"port"`
	Db       *string `json:"db" yaml:"db"`
	User     *string `json:"user" yaml:"user"`
	Password *string `json:"password" yaml:"password"`
	SslMode  *string `json:"sslmode" yaml:"sslmode"`
	Url      *string `json:"url" yaml:"url"`
}

func (d *Database) String() string {
	content := strings.Join(
		[]string{
			fmt.Sprintf("Host: %v", *d.Host),
			fmt.Sprintf("Port: %v", *d.Port),
			fmt.Sprintf("Db: %v", *d.Db),
			fmt.Sprintf("User: %v", *d.User),
			fmt.Sprintf("Password: %v", *d.Password),
			fmt.Sprintf("SslMode: %v", *d.SslMode),
			fmt.Sprintf("Url: %v", *d.Url),
		},
		", ",
	)

	return fmt.Sprintf("Database{%s}", content)
}

func (d *Database) GetUrl(engine string) (string, error) {
	if engine == "mysql" {
		return getMysqlUrl(d), nil
	}

	if engine == "postgres" {
		return getPgUrl(d), nil
	}

	return "", errorx.IllegalArgument.New("invalid engine: %s", engine)
}

func (d *Database) Connect(engine string) (*sqlx.DB, error) {
	connectionUrl, err := d.GetUrl(engine)

	if err != nil {
		return nil, errorx.IllegalState.Wrap(err, "failed to get connection url")
	}

	return sqlx.Connect(engine, connectionUrl)
}

func (d *Database) Merge(other *Database) *Database {
	if other == nil {
		return d
	}

	if d == nil {
		return other
	}

	if other.Host != nil {
		d.Host = other.Host
	}

	if other.Port != nil {
		d.Port = other.Port
	}

	if other.Db != nil {
		d.Db = other.Db
	}

	if other.User != nil {
		d.User = other.User
	}

	if other.Password != nil {
		d.Password = other.Password
	}

	if other.SslMode != nil {
		d.SslMode = other.SslMode
	}

	if other.Url != nil {
		d.Url = other.Url
	}

	return d
}

func getPgUrl(d *Database) string {
	isUrlEmpty := d.Url == nil || *d.Url == ""

	isHostEmpty := d.Host == nil || *d.Host == ""

	if isHostEmpty && !isUrlEmpty {
		return *d.Url
	}

	port := safeValue("5432", d.Port)

	host := safeValue("localhost", d.Host)

	db := safeValue("postgres", d.Db)

	user := safeValue("postgres", d.User)

	password := safeValue("", d.Password)

	sslMode := safeValue("disable", d.SslMode)

	safeUrl := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		user,
		password,
		host,
		port,
		db,
		sslMode,
	)

	return safeUrl
}

func getMysqlUrl(d *Database) string {
	isUrlEmpty := d.Url == nil || *d.Url == ""

	isHostEmpty := d.Host == nil || *d.Host == ""

	if isHostEmpty && !isUrlEmpty {
		return *d.Url
	}

	port := safeValue("3306", d.Port)

	host := safeValue("localhost", d.Host)

	db := safeValue("mysql", d.Db)

	user := safeValue("root", d.User)

	password := safeValue("", d.Password)

	safeUrl := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true",
		user,
		password,
		host,
		port,
		db,
	)

	return safeUrl
}

func safeValue(left string, right *string) string {
	if right != nil && *right != "" {
		return *right
	}

	return left
}
