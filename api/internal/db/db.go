package db

type NoSqlDb interface {
	Ping() (string, error)
	Update(key string, value interface{}) error
	Get(key string, path string) (interface{}, error)
}
