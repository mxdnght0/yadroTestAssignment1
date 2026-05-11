package file

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
	"sync"
	"yadroTestAssignment/server/internal/application/contracts"
	"yadroTestAssignment/server/internal/config"
)

const nameserverPrefix = "nameserver "

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

func toLine(ip string) string {
	return nameserverPrefix + ip
}

func fromLine(line string) (string, bool) {
	if strings.HasPrefix(line, nameserverPrefix) {
		return strings.TrimPrefix(line, nameserverPrefix), true
	}
	return "", false
}

func (s *Storage) SaveDNS(dns string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	file, err := os.OpenFile(s.path, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer func() { _ = file.Close() }()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if ip, ok := fromLine(scanner.Text()); ok && ip == dns {
			return ErrDNSAlreadyExists
		}
	}
	if err := scanner.Err(); err != nil {
		return err
	}

	_, err = fmt.Fprintln(file, toLine(dns))
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

	var kept []string
	found := false
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if ip, ok := fromLine(line); ok && ip == dns {
			found = true
		} else {
			kept = append(kept, line)
		}
	}
	scanErr := scanner.Err()
	_ = file.Close()
	if scanErr != nil {
		return scanErr
	}
	if !found {
		return ErrDNSNotFound
	}

	out, err := os.OpenFile(s.path, os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer func() { _ = out.Close() }()

	w := bufio.NewWriter(out)
	for _, line := range kept {
		_, err = fmt.Fprintln(w, line)
		if err != nil {
			return err
		}
	}
	return w.Flush()
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

	var ips []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if ip, ok := fromLine(scanner.Text()); ok {
			ips = append(ips, ip)
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return ips, nil
}
