package mapper

import (
	"net"
	"sync"
)

type Mapper struct {
	lookup func(string) ([]net.IP, error)
	cidr   *Cidr
	domain *domain

	lookupLock sync.RWMutex
}

func (x *Mapper) SetLookup(f func(string) ([]net.IP, error)) {
	x.lookupLock.Lock()
	defer x.lookupLock.Unlock()
	x.lookup = f
}

func (x *Mapper) Insert(str string, mark interface{}) {
	if str == "" {
		return
	}

	_, ipNet, err := net.ParseCIDR(str)
	if err != nil {
		x.domain.Insert(str, mark)
	} else {
		x.cidr.InsertCIDR(ipNet, mark)
	}
}

func (x *Mapper) Search(str string) (mark interface{}) {
	if ip := net.ParseIP(str); ip != nil {
		mark, _ = x.cidr.SearchIP(ip)
		goto _end
	}

	mark, _ = x.domain.Search(str)
	if mark != nil {
		goto _end
	}

	x.lookupLock.RLock()
	defer x.lookupLock.RUnlock()
	if x.lookup == nil {
		goto _end
	}
	if dns, err := x.lookup(str); err == nil {
		mark, _ = x.cidr.SearchIP(dns[0])
	}

_end:
	return mark
}

func (x *Mapper) Clear() {
	x.cidr = NewCidrMapper()
	x.domain = NewDomainMapper()
}

func NewMapper(lookup func(string) ([]net.IP, error)) (matcher *Mapper) {
	return &Mapper{
		cidr:   NewCidrMapper(),
		domain: NewDomainMapper(),
		lookup: lookup,
	}
}
