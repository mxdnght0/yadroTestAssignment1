package file

import (
	"bufio"
	"errors"
	"os"
	"sync"
	"yadroTestAssignment/server/internal/application/contracts"
	"yadroTestAssignment/server/internal/config"
)

type Storage struct {
	path string
	mu   sync.Mutex
}

func NewStorage(cfg *config.Config) contracts.DNSStorage {
	return &Storage{
		path: cfg.GetPath(),
	}
}

func (s *Storage) fileExists() bool {
	_, err := os.Stat(s.path)
	return !errors.Is(err, os.ErrNotExist)
}

func (s *Storage) SaveDNS(dns string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	file, err := os.OpenFile(s.path, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer func() { _ = file.Close() }()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if scanner.Text() == dns {
			return ErrDNSAlreadyExists
		}
	}
	if err := scanner.Err(); err != nil {
		return err
	}

	_, err = file.WriteString(dns + "\n")

	return err
}

func (s *Storage) DeleteDNS(dns string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if !s.fileExists() {
		return ErrFileIsNotCreated
	}

	file, err := os.Open(s.path)
	if err != nil {
		return err
	}
	defer func() { _ = file.Close() }()

	var lines []string
	found := false
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == dns {
			found = true
		} else {
			lines = append(lines, line)
		}
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	if !found {
		return ErrDNSNotFound
	}

	err = os.WriteFile(s.path, []byte(""), 0666)
	if err != nil {
		return err
	}
	out, err := os.OpenFile(s.path, os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	defer func() { _ = out.Close() }()
	for _, line := range lines {
		_, err = out.WriteString(line + "\n")
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *Storage) GetAllDNS() ([]string, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if !s.fileExists() {
		return nil, ErrFileIsNotCreated
	}

	file, err := os.Open(s.path)
	if err != nil {
		return nil, err
	}
	defer func() { _ = file.Close() }()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}
