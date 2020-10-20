package repository

type Repository struct {
	postgresql string
	cassandra string
	hbase string
}

func NewRepository() *Repository {
	return &Repository{"psql", "csdr", "hb"}
}