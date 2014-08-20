package dbutil

type Row interface {
    Scan(dest ...interface{}) error
}

type ParseRowFunc func(row Row) (Object, error)

type Object interface {
    MarshalJSON() ([]byte, error)
}

