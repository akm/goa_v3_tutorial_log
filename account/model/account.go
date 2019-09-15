package accountmodel

type Account struct {
	Email    string `datastore:"email" boom:"id"`
	JwtToken string `datastore:"jwt_token"`
}
