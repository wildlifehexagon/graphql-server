package repository

type Repository interface {
	Get(string) (string, error)
	Set(string, string) error
	Exists(string) (bool, error)
	HGet(string, string) (string, error)
	HSet(string, string, string) error
	HExists(string, string) (bool, error)
}
