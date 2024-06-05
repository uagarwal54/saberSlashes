package reading

type Repository interface {
	GetAllCandies() (candies []string, err error)
}

type Service interface {
	GetAllCandies() (candies []string, err error)
}

type service struct {
	r Repository
}

// Here since r is an instance of an interface so, it can take the type of any argument passed while calling this function
func NewService(r Repository) *service {
	return &service{r}
}

// GetAllCandies is the function defined in the Repository interface
func (s *service) GetAllCandies() (candies []string, err error) {
	// The below GetAllCandies function is assosiated with the object, with which, the service struct is instantiated
	// using the NewService method. In this case it is the object of the storage struct defined in the storage-mysql.go file
	if candies, err = s.r.GetAllCandies(); err != nil {
		return
	}
	return
}
