package types

import (
	"code.google.com/p/go-uuid/uuid"
)

type Host struct {
    Id string
    Hostname string
	ServiceIp string
	HashStart string
	HashEnd string
}

func (h *Host) GenID() {
	if h.Id == "" {
		h.Id = uuid.New()
	}

}