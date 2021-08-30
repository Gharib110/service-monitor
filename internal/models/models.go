package models

import (
	"errors"
	"time"
)

var (
	// ErrNoRecord no record found in database error
	ErrNoRecord = errors.New("models: no matching record found")
	// ErrInvalidCredentials invalid username/password error
	ErrInvalidCredentials = errors.New("models: invalid credentials")
	// ErrDuplicateEmail duplicate email error
	ErrDuplicateEmail = errors.New("models: duplicate email")
	// ErrInactiveAccount inactive account error
	ErrInactiveAccount = errors.New("models: Inactive Account")
)

// User model
type User struct {
	ID          int
	FirstName   string
	LastName    string
	UserActive  int
	AccessLevel int
	Email       string
	Password    []byte
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
	Preferences map[string]string
}

// Preference model
type Preference struct {
	ID         int
	Name       string
	Preference []byte
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

// Host uses for holding data about Each Host
type Host struct {
	ID            int
	HostName      string
	CanonicalName string
	URL           string
	IP            string
	IPV6          string
	Location      string
	OS            string
	Active        int
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

// Services uses for holding data about Each Service
type Services struct {
	ID          int
	Icon        string
	ServiceName string
	Active      int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// HostServices uses for managing and holding data about each Host and Services
type HostServices struct {
	ID             int
	ServiceID      int
	HostID         int
	Active         int
	ScheduleNumber int
	Status         string
	ScheduleUnit   string
	LastCheck      time.Time
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
