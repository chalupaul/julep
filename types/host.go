package types

import (
	"code.google.com/p/go-uuid/uuid"
)

type Host struct {
    Id string
	Weight int
    Name string
	ServiceIp string
	ServicePort int
	HashStart string
	HashEnd string
}

type HostGroup struct {
	Id string
	Name string
	Weight int
	Hosts []Host
	Groups []HostGroup
}
func (h *Host) GenID() {
	if h.Id == "" {
		h.Id = uuid.New()
	}

}