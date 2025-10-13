package chain

import (
	"context"
	"fmt"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
)

// Client wraps the Solana RPC client.
// No fancy abstraction. It's just an RPC client.
type Client struct {
	rpc *rpc.Client
}

// NewClient creates a Solana RPC client
// If connection fails, you get an error. Handle it.
func NewClient(rpcURL string) (*Client, error) {
	client := rpc.New(rpcURL)
	return &Client{rpc: client}, nil
}

// GetLatestBlockhash gets the most recent blockhash
// Solana blockhashes expire in ~60-90 seconds. Don't cache this.
func (c *Client) GetLatestBlockhash(ctx context.Context) (solana.Hash, error) {
	resp, err := c.rpc.GetLatestBlockhash(ctx, rpc.CommitmentFinalized)
	if err != nil {
		return solana.Hash{}, fmt.Errorf("failed to get blockhash: %w", err)
	}

	return resp.Value.Blockhash, nil
}

// GetBalance queries account balance in lamports
// 1 SOL = 1,000,000,000 lamports
// Don't do unit conversion here. That's the caller's job.
func (c *Client) GetBalance(ctx context.Context, addr solana.PublicKey) (uint64, error) {
	resp, err := c.rpc.GetBalance(ctx, addr, rpc.CommitmentFinalized)
	if err != nil {
		return 0, fmt.Errorf("failed to get balance for %s: %w", addr, err)
	}

	return resp.Value, nil
}

// Close closes the RPC client connection
func (c *Client) Close() error {
	// The gagliardetto client doesn't need explicit close
	// but keeping this for consistency with other projects
	return nil
}
