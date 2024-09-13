package inmemory

type InMemoryStorage struct {
	storage map[string]string
}

func New() *InMemoryStorage {
	return &InMemoryStorage{
		storage: make(map[string]string),
	}
}
