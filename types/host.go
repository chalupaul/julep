package types

import (
	"code.google.com/p/go-uuid/uuid"
	"sort"
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
	Weight int
	Hosts []Host
	ChildGroup *HostGroup
}

func (h *Host) GenID() {
	if h.Id == "" {
		h.Id = uuid.New()
	}

}

type HostById []Host

func (hs HostById) Len() int {
	return len(hs)
}

func (hs HostById) Swap(i, j int) {
	hs[i], hs[j] = hs[j], hs[i]
}

func (hs HostById) Less(i, j int) bool {
	return hs[i].Id < hs[j].Id
}

func (h HostGroup) OrderHostIds() {
	sort.Sort(HostById(h.Hosts))
	if h.ChildGroup != nil {
		h.ChildGroup.OrderHostIds()
	}
}

func (h Host) AssignHashBoundaries() {
	
}
