package service

import (
	"yadroTestAssignment/client/internal/contracts"
)

type DNS struct {
	client contracts.DNSClient
}

func NewDNS(client contracts.DNSClient) contracts.DNSService {
	return &DNS{client: client}
}

func (s *DNS) Add(ip string) error {
	return s.client.Add(ip)
}

func (s *DNS) Delete(ip string) error {
	return s.client.Delete(ip)
}

func (s *DNS) GetAll() ([]string, error) {
	return s.client.GetAll()
}
