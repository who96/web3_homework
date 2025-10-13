package events

import (
	"context"
	"fmt"
	"log"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/gagliardetto/solana-go/rpc/ws"
)

// Listener handles WebSocket subscriptions for Solana transactions
// No fancy event bus. Just a simple subscription handler.
type Listener struct {
	wsClient *ws.Client
}

// NewListener creates a WebSocket listener
func NewListener(wsURL string) (*Listener, error) {
	client, err := ws.Connect(context.Background(), wsURL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to WebSocket: %w", err)
	}

	return &Listener{wsClient: client}, nil
}

// SubscribeSignature subscribes to a specific transaction signature
// Blocks until transaction is finalized or error occurs.
// Returns the final transaction status.
func (l *Listener) SubscribeSignature(ctx context.Context, sig solana.Signature) error {
	sub, err := l.wsClient.SignatureSubscribe(
		sig,
		rpc.CommitmentFinalized,
	)
	if err != nil {
		return fmt.Errorf("failed to subscribe to signature: %w", err)
	}
	defer sub.Unsubscribe()

	log.Printf("Subscribed to transaction: %s", sig)
	log.Printf("View on explorer: https://explorer.solana.com/tx/%s?cluster=devnet", sig)

	// Wait for notification
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()

		default:
			got, err := sub.Recv(ctx)
			if err != nil {
				return fmt.Errorf("subscription error: %w", err)
			}

			if got == nil {
				continue
			}

			// Transaction confirmed
			if got.Value.Err != nil {
				log.Printf("❌ Transaction failed: %v", got.Value.Err)
				return fmt.Errorf("transaction failed: %v", got.Value.Err)
			}

			log.Printf("✅ Transaction finalized!")
			log.Printf("   Slot: %d", got.Context.Slot)
			return nil
		}
	}
}

// SubscribeAccount subscribes to account changes
// Useful for monitoring token balance changes
func (l *Listener) SubscribeAccount(ctx context.Context, account solana.PublicKey) error {
	sub, err := l.wsClient.AccountSubscribe(
		account,
		rpc.CommitmentFinalized,
	)
	if err != nil {
		return fmt.Errorf("failed to subscribe to account: %w", err)
	}
	defer sub.Unsubscribe()

	log.Printf("Monitoring account: %s", account)

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()

		default:
			got, err := sub.Recv(ctx)
			if err != nil {
				return fmt.Errorf("subscription error: %w", err)
			}

			if got == nil {
				continue
			}

			log.Printf("Account updated:")
			log.Printf("  Lamports: %d", got.Value.Lamports)
			log.Printf("  Owner: %s", got.Value.Owner)
			log.Printf("  Slot: %d", got.Context.Slot)
		}
	}
}

// Close closes the WebSocket connection
func (l *Listener) Close() error {
	l.wsClient.Close()
	return nil
}
