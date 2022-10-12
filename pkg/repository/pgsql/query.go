package pgsql

type Row interface {
	Scan(dest ...interface{}) error
}
