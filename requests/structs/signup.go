package structs

import (
    "github.com/alexkarpovich/mowy-api-go/database"
    "github.com/alexkarpovich/mowy-api-go/database/regs"
    "github.com/alexkarpovich/mowy-api-go/database/users"
)

type Signup struct {
    Base
    Email string `json:"email"`
}

func (s *Signup) Validate() bool {
    var user users.User
    var reg regs.Registration
    s.Errors = make(map[string]string)

    database.Db.Where("email=?", s.Email).First(&user)
    database.Db.Where("email=?", s.Email).First(&reg)

    if user.ID != 0 || reg.ID != 0 {
        s.Errors["email"] = "The email is already in use."
    }

    return len(s.Errors) == 0
}
