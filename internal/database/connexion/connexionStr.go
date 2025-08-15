package connexion

import "strings"

// connectionStringComponent represents a component of a connection string in the key=value format
type connectionStringComponent interface {
	key() string
	value() string
}

type dbName string

func (d dbName) key() string {
	return "dbname"
}

func (d dbName) value() string {
	return string(d)
}

type user string

func (u user) key() string {
	return "user"
}
func (u user) value() string {
	return string(u)
}

type password string

func (u password) key() string {
	return "password"
}

func (u password) value() string {
	return string(u)
}

type host string

func (u host) key() string {
	return "host"
}

func (u host) value() string {
	return string(u)
}

type port string

func (p port) key() string {
	return "port"
}

func (p port) value() string {
	return string(p)
}

type sslMode string

func (s sslMode) key() string {
	return "sslmode"
}

func (s sslMode) value() string {
	return string(s)
}

const verifyFull = sslMode("verify-full")
const noSsl = sslMode("disable")

func generateCnxnString(comp ...connectionStringComponent) string {
	builder := strings.Builder{}
	for _, v := range comp {
		builder.WriteString(v.key())
		builder.WriteRune('=')
		builder.WriteString(v.value())
		builder.WriteRune(' ')
	}
	return builder.String()
}
