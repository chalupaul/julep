package types

import (
	"code.google.com/p/go-uuid/uuid"
)

type Host struct {
    Id string
	Weight int
    Hostname string
	ServiceIp string
	ServicePort int
	HashStart string
	HashEnd string
}

type HostGroup struct {
	Id string
	Weight int
	Hosts []Host
	HostGroups []HostGroup
}
func (h *Host) GenID() {
	if h.Id == "" {
		h.Id = uuid.New()
	}

}