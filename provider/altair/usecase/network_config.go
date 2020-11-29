package usecase

import "time"

// A NetworkConfig for altair client
type NetworkConfig struct {
	Cfg struct {
		Host          string
		Username      string
		Password      string
		Timeout       time.Duration
		Retry         int
		SleepDuration time.Duration
	}
}

// Host of altair
func (n *NetworkConfig) Host() string {
	return n.Cfg.Host
}

// Username of altair basic auth
func (n *NetworkConfig) Username() string {
	return n.Cfg.Username
}

// Password of altair basic auth
func (n *NetworkConfig) Password() string {
	return n.Cfg.Password
}

// Timeout of client http call
func (n *NetworkConfig) Timeout() time.Duration {
	return n.Cfg.Timeout
}

// Retry times if request fail
func (n *NetworkConfig) Retry() int {
	return n.Cfg.Retry
}

// RetrySleepDuration sleep duration for every retry
func (n *NetworkConfig) RetrySleepDuration() time.Duration {
	return n.Cfg.SleepDuration
}
