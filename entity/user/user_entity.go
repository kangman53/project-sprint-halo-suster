package user_entity

type User struct {
	Id                  string
	Nip                 string
	Name                string
	Role                string
	Password            string
	IdentityCardScanImg string
}

type UserRepository struct {
	Id        string
	Name      string
	Nip       string
	Role      string
	Limit     int
	Offset    int
	CreatedAt string
}
