package chain

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/gagliardetto/solana-go"
)

// LoadKeypairFromFile loads a Solana keypair from JSON file
// The file format is the standard Solana CLI format: [byte, byte, ...]
func LoadKeypairFromFile(path string) (solana.PrivateKey, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return solana.PrivateKey{}, fmt.Errorf("failed to read keypair file: %w", err)
	}

	var keyBytes []byte
	if err := json.Unmarshal(data, &keyBytes); err != nil {
		return solana.PrivateKey{}, fmt.Errorf("failed to parse keypair JSON: %w", err)
	}

	if len(keyBytes) != 64 {
		return solana.PrivateKey{}, fmt.Errorf("invalid keypair length: expected 64 bytes, got %d", len(keyBytes))
	}

	privateKey := solana.PrivateKey(keyBytes)
	return privateKey, nil
}

// GetDefaultKeypairPath returns the default Solana CLI keypair path
func GetDefaultKeypairPath() string {
	home, _ := os.UserHomeDir()
	return home + "/.config/solana/id.json"
}
