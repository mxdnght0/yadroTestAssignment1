package contracts

type DNSServer interface {
	SaveDNS(dns string) error
	DeleteDNS(dns string) error
	GetAllDNS() ([]string, error)
}
