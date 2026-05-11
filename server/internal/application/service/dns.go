package service

import (
	"errors"
	contracts2 "yadroTestAssignment/server/internal/application/contracts"
	"yadroTestAssignment/server/internal/domain"
)

type DNS struct {
	storage contracts2.DNSStorage
}

func NewDNS(storage contracts2.DNSStorage) contracts2.DNSServer {
	return &DNS{
		storage: storage,
	}
}

func (D *DNS) SaveDNS(dns string) error {
	err := D.storage.SaveDNS(dns)
	if errors.Is(err, domain.ErrDNSAlreadyExists) {
		return ErrDNSAlreadyExists
	}
	return err
}

func (D *DNS) DeleteDNS(dns string) error {
	err := D.storage.DeleteDNS(dns)
	if errors.Is(err, domain.ErrFileIsNotCreated) {
		return ErrFileIsNotCreated
	}
	if errors.Is(err, domain.ErrDNSNotFound) {
		return ErrDNSNotFound
	}
	return err
}

func (D *DNS) GetAllDNS() ([]string, error) {
	lines, err := D.storage.GetAllDNS()
	if errors.Is(err, domain.ErrFileIsNotCreated) {
		return nil, ErrFileIsNotCreated
	}
	if err != nil {
		return nil, err
	}
	return lines, nil
}
