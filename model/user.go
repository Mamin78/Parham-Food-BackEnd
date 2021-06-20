package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID primitive.ObjectID `json:"id,omitempty" bson:"_id"`

	Name        string  `json:"name" bson:"name"`
	PhoneNumber string  `json:"phone_number,omitempty" bson:"phone_number"`
	Password    string  `json:"password" bson:"password"`
	Credit      float64 `json:"credit" bson:"credit"`
	Area        int     `json:"area" bson:"area"`
	Address     string  `json:"address" bson:"address"`

	Orders   []primitive.ObjectID `json:"orders" bson:"orders"`
	Comments []primitive.ObjectID `json:"comments" bson:"comments"`
}

func NewUser(res *User) *User {
	r := new(User)
	r.ID = res.ID
	r.Name = res.Name
	r.PhoneNumber = res.PhoneNumber
	r.Password = res.Password
	r.Credit = res.Credit
	r.Area = res.Area
	r.Address = res.Address
	return r
}

//func (u *User) HashedPassword(plain string) (string, error) {
//	if len(plain) == 0 {
//		return "", errors.New("password can npt be empty")
//	}
//	h, err := bcrypt.GenerateFromPassword([]byte(plain), bcrypt.DefaultCost)
//	return string(h), err
//}
//
//func (u *User) PasswordsAreSame(plain string) bool {
//	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(plain))
//	return err == nil
//}

//func (u *User) CheckPasswordLever() error {
//	pass := strings.ToLower(u.Password)
//	if len(pass) < 8 {
//		return fmt.Errorf("password len is < 8")
//	}
//	num := `[0-9]{1}`
//	aToz := `[a-z]{1}`
//	symbol := `[!@#~$%^&*()+|_]{1}`
//	if b, _ := regexp.MatchString(num, pass); !b {
//		return fmt.Errorf("password need number")
//	}
//	if b, _ := regexp.MatchString(aToz, pass); !b {
//		return fmt.Errorf("password need charachter")
//	}
//	if b, _ := regexp.MatchString(symbol, pass); b {
//		return fmt.Errorf("password doesn't need symbol")
//	}
//	return nil
//}
