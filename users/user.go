package users

import (
	"fmt"
	"godesde0/modelos"
	"time"
)

func AltaUsuario() {
	usuario := new(modelos.User)
	usuario.AddUser(10, "Rufino", time.Now(), true)
	fmt.Println(usuario)
}
