package config

// Solana Devnet endpoints
const (
	DevnetRPC = "https://api.devnet.solana.com"
	DevnetWS  = "wss://api.devnet.solana.com"
)

// Config holds Solana connection configuration
// No fancy builder pattern, no validation circus.
// If you need different networks, change the constants.
type Config struct {
	RPCURL string
	WSURL  string
}

// New returns a Config for Solana Devnet
// Want mainnet? Change the constants. Don't make me add a factory.
func New() *Config {
	return &Config{
		RPCURL: DevnetRPC,
		WSURL:  DevnetWS,
	}
}
