package schema

type Person  struct {
	ID         int    `json:"id"`
	FirstName  string `index:"2" size:"fixed" maxlen:"50" json:"first_name"`
	LastName   string `index:"3" size:"fixed" maxlen:"50" json:"last_name"`
	Age        int    `index:"4" size:"fixed" maxlen:"3" json:"age"`
	Email      string `index:"5" size:"fixed" maxlen:"100" json:"email"`
	PhoneNumber string `index:"6" size:"fixed" maxlen:"15" json:"phone_number"`
	Address    string `index:"7" size:"llvar" maxlen:"255" json:"address"`
	City       string `index:"8" size:"fixed" maxlen:"100" json:"city"`
	State      string `index:"9" size:"fixed" maxlen:"50" json:"state"`
	ZipCode    string `index:"10" size:"fixed" maxlen:"10" json:"zip_code"`
}
