package dom

// Admin model stores admin credentials
type Admin struct {
	Username string
	Password string
}

type AdminModel struct {
	AdminID  uint
	Username string
	Email    string
	Role     string
	Password string
}
