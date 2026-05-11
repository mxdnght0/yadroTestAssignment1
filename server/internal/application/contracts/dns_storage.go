package contracts

type DNSStorage interface {
	SaveDNS(dns string) error
	DeleteDNS(dns string) error
	GetAllDNS() ([]string, error)
}
