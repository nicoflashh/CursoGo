package users

import (
	"fmt"
	"time"

	"github.com/Cursogo/modelos"
)

func AltaUsuario() {
	u := new(modelos.User)
	u.AddUser(22, "Pablo", time.Now(), true)
	fmt.Println(u)
}
