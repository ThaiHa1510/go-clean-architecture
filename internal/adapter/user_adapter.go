package adapter

type Adapter interface {
    UserRepository() UserRepository
}

type InMemoryAdapter struct {
    users []*User
}

func NewInMemoryAdapter() *InMemoryAdapter {
    return &InMemoryAdapter{users: []*User{}}
}

func (adapter *InMemoryAdapter) UserRepository() UserRepository {
    return &inMemoryUserRepository{adapter: adapter}
}
