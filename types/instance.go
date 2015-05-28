package types

import (
	"code.google.com/p/go-uuid/uuid"
	"math/big"
	"crypto/md5"
	"encoding/hex"
	"fmt"
)


type Instance struct {
    Id string
    Hostname string
}

func (i *Instance) GenID() *big.Int {
	if i.Id == "" {
		i.Id = uuid.New()
	}
	id := big.NewInt(0)
	h := md5.New()
	h.Write([]byte(i.Id))
	idHex := hex.EncodeToString(h.Sum(nil))
	if _, ok := id.SetString(idHex, 16); ok {
		fmt.Printf("number = %v\n", id)
	} else {
		fmt.Printf("instance id %#v too large\n", id)
	}
	return id

}