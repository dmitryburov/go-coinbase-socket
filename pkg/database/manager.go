package database

// Manager is an interface database of application
type Manager interface {
	// PingCheck check health connection
	PingCheck() error
	// CloseConnect is closing current connection
	CloseConnect()
}
