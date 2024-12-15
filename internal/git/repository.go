package git

// Repository represents git repository
type Repository struct {
	path string
}

// Path returns path to git repository root
func (r Repository) Path() string {
	return r.path
}

// NewRepository returns new git repository
func NewRepository(path string) Repository {
	return Repository{path: path}
}

// OpenRepository opens and checks git repository
func OpenRepository(path string) (Repository, error) {
	err := AssertRepository(path)
	if err != nil {
		return Repository{}, err
	}
	return NewRepository(path), nil
}
