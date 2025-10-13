package token

import (
	"context"
	"fmt"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/programs/token"
	"github.com/gagliardetto/solana-go/rpc"
)

// Client wraps token operations
type Client struct {
	rpc *rpc.Client
}

// NewClient creates a token client
func NewClient(rpcURL string) (*Client, error) {
	client := rpc.New(rpcURL)
	return &Client{rpc: client}, nil
}

// GetTokenBalance gets SPL token balance for an account
func (c *Client) GetTokenBalance(ctx context.Context, tokenAccount solana.PublicKey) (float64, error) {
	resp, err := c.rpc.GetTokenAccountBalance(ctx, tokenAccount, rpc.CommitmentFinalized)
	if err != nil {
		return 0, fmt.Errorf("failed to get token balance: %w", err)
	}

	if resp.Value.UiAmount == nil {
		return 0, nil
	}
	return *resp.Value.UiAmount, nil
}

// GetTokenAccountsByOwner gets all token accounts owned by a wallet
func (c *Client) GetTokenAccountsByOwner(ctx context.Context, owner solana.PublicKey) ([]solana.PublicKey, error) {
	tokenProgramID := solana.TokenProgramID
	resp, err := c.rpc.GetTokenAccountsByOwner(
		ctx,
		owner,
		&rpc.GetTokenAccountsConfig{
			ProgramId: &tokenProgramID,
		},
		&rpc.GetTokenAccountsOpts{
			Encoding: solana.EncodingBase64,
		},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get token accounts: %w", err)
	}

	var accounts []solana.PublicKey
	for _, acc := range resp.Value {
		accounts = append(accounts, acc.Pubkey)
	}

	return accounts, nil
}

// TransferToken transfers SPL tokens
func (c *Client) TransferToken(
	ctx context.Context,
	from solana.PrivateKey,
	fromTokenAccount solana.PublicKey,
	toTokenAccount solana.PublicKey,
	amount uint64,
) (solana.Signature, error) {

	// Get recent blockhash
	recent, err := c.rpc.GetLatestBlockhash(ctx, rpc.CommitmentFinalized)
	if err != nil {
		return solana.Signature{}, fmt.Errorf("failed to get blockhash: %w", err)
	}

	// Build transfer instruction
	instruction := token.NewTransferInstruction(
		amount,
		fromTokenAccount,
		toTokenAccount,
		from.PublicKey(),
		[]solana.PublicKey{},
	).Build()

	// Create transaction
	tx, err := solana.NewTransaction(
		[]solana.Instruction{instruction},
		recent.Value.Blockhash,
		solana.TransactionPayer(from.PublicKey()),
	)
	if err != nil {
		return solana.Signature{}, fmt.Errorf("failed to create transaction: %w", err)
	}

	// Sign transaction
	_, err = tx.Sign(func(key solana.PublicKey) *solana.PrivateKey {
		if key.Equals(from.PublicKey()) {
			return &from
		}
		return nil
	})
	if err != nil {
		return solana.Signature{}, fmt.Errorf("failed to sign transaction: %w", err)
	}

	// Send transaction
	sig, err := c.rpc.SendTransactionWithOpts(
		ctx,
		tx,
		rpc.TransactionOpts{
			SkipPreflight:       false,
			PreflightCommitment: rpc.CommitmentFinalized,
		},
	)
	if err != nil {
		return solana.Signature{}, fmt.Errorf("failed to send transaction: %w", err)
	}

	return sig, nil
}
