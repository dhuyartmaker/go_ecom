package repo

type IUserRepository interface {
	GetUserByEmail(email string) string
}

type userRepository struct{}

// GetUserByEmail implements IUserRepository.
func (*userRepository) GetUserByEmail(email string) string {
	return "Huynh Duc Huy"
}

func NewUserRepository() IUserRepository {
	return &userRepository{}
}
