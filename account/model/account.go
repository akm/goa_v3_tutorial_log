package accountmodel

type Account struct {
	ID       int64 `datastore:"-" boom:"id"`
	Email    string
	JwtToken string
}
