package contracts

type DNSService interface {
	Add(ip string) error
	Delete(ip string) error
	GetAll() ([]string, error)
}
