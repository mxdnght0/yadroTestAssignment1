package service_test

import (
	"errors"
	"testing"
	"yadroTestAssignment/server/internal/application/service"
	"yadroTestAssignment/server/internal/domain"
)

type mockStorage struct {
	records []string
	saveErr error
	delErr  error
	getErr  error
}

func (m *mockStorage) SaveDNS(dns string) error {
	if m.saveErr != nil {
		return m.saveErr
	}
	for _, r := range m.records {
		if r == dns {
			return domain.ErrDNSAlreadyExists
		}
	}
	m.records = append(m.records, dns)
	return nil
}

func (m *mockStorage) DeleteDNS(dns string) error {
	if m.delErr != nil {
		return m.delErr
	}
	for i, r := range m.records {
		if r == dns {
			m.records = append(m.records[:i], m.records[i+1:]...)
			return nil
		}
	}
	return domain.ErrDNSNotFound
}

func (m *mockStorage) GetAllDNS() ([]string, error) {
	if m.getErr != nil {
		return nil, m.getErr
	}
	return m.records, nil
}

func TestSaveDNS_Success(t *testing.T) {
	svc := service.NewDNS(&mockStorage{})
	if err := svc.SaveDNS("8.8.8.8"); err != nil {
		t.Fatalf("expected nil, got %v", err)
	}
}

func TestSaveDNS_Duplicate(t *testing.T) {
	storage := &mockStorage{records: []string{"8.8.8.8"}}
	svc := service.NewDNS(storage)
	err := svc.SaveDNS("8.8.8.8")
	if !errors.Is(err, service.ErrDNSAlreadyExists) {
		t.Fatalf("expected ErrDNSAlreadyExists, got %v", err)
	}
}

func TestSaveDNS_StorageError(t *testing.T) {
	storageErr := errors.New("disk full")
	svc := service.NewDNS(&mockStorage{saveErr: storageErr})
	err := svc.SaveDNS("8.8.8.8")
	if !errors.Is(err, storageErr) {
		t.Fatalf("expected storage error, got %v", err)
	}
}

func TestDeleteDNS_Success(t *testing.T) {
	storage := &mockStorage{records: []string{"8.8.8.8"}}
	svc := service.NewDNS(storage)
	if err := svc.DeleteDNS("8.8.8.8"); err != nil {
		t.Fatalf("expected nil, got %v", err)
	}
}

func TestDeleteDNS_NotFound(t *testing.T) {
	svc := service.NewDNS(&mockStorage{})
	err := svc.DeleteDNS("1.1.1.1")
	if !errors.Is(err, service.ErrDNSNotFound) {
		t.Fatalf("expected ErrDNSNotFound, got %v", err)
	}
}

func TestDeleteDNS_FileNotCreated(t *testing.T) {
	svc := service.NewDNS(&mockStorage{delErr: domain.ErrFileIsNotCreated})
	err := svc.DeleteDNS("8.8.8.8")
	if !errors.Is(err, service.ErrFileIsNotCreated) {
		t.Fatalf("expected ErrFileIsNotCreated, got %v", err)
	}
}

func TestGetAllDNS_Success(t *testing.T) {
	storage := &mockStorage{records: []string{"8.8.8.8", "1.1.1.1"}}
	svc := service.NewDNS(storage)
	lines, err := svc.GetAllDNS()
	if err != nil {
		t.Fatalf("expected nil, got %v", err)
	}
	if len(lines) != 2 {
		t.Fatalf("expected 2 records, got %d", len(lines))
	}
}

func TestGetAllDNS_Empty(t *testing.T) {
	svc := service.NewDNS(&mockStorage{})
	lines, err := svc.GetAllDNS()
	if err != nil {
		t.Fatalf("expected nil, got %v", err)
	}
	if len(lines) != 0 {
		t.Fatalf("expected 0 records, got %d", len(lines))
	}
}

func TestGetAllDNS_FileNotCreated(t *testing.T) {
	svc := service.NewDNS(&mockStorage{getErr: domain.ErrFileIsNotCreated})
	_, err := svc.GetAllDNS()
	if !errors.Is(err, service.ErrFileIsNotCreated) {
		t.Fatalf("expected ErrFileIsNotCreated, got %v", err)
	}
}
