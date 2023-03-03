package exchange

// Manager is an interface exchange of application
type Manager interface {
	// CloseConnection is closing connection
	CloseConnection() error
	// WriteData command write data to exchange connection
	WriteData(message []byte) (int, error)
	// ReadData command is reading from receiver data
	ReadData() ([]byte, error)
}
