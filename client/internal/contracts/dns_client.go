package contracts

type DNSClient interface {
	Add(ip string) error
	Delete(ip string) error
	GetAll() ([]string, error)
}
