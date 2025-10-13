package token

import (
	"context"
	"fmt"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/programs/system"
	"github.com/gagliardetto/solana-go/programs/token"
	"github.com/gagliardetto/solana-go/rpc"
)

// SwapConfig holds the swap pool configuration
type SwapConfig struct {
	PoolAuthorityKey solana.PrivateKey
	PoolATokenAccount solana.PublicKey  // Pool's a_t account
	PoolBTokenAccount solana.PublicKey  // Pool's b_t account
	MintA             solana.PublicKey  // a_t mint
	MintB             solana.PublicKey  // b_t mint
	ExchangeRate      uint64            // 1 a_t = ExchangeRate b_t (default: 100)
}

// SwapAtoB swaps a_t tokens for b_t tokens at fixed rate 1:100
// User sends amount_a of a_t, receives amount_a * 100 of b_t
func (c *Client) SwapAtoB(
	ctx context.Context,
	userKey solana.PrivateKey,
	userAAccount solana.PublicKey,  // User's a_t token account
	userBAccount solana.PublicKey,  // User's b_t token account
	amountA uint64,                 // Amount of a_t to swap
	swapConfig *SwapConfig,
) (solana.Signature, error) {
	// Calculate amount of b_t to receive (1:100 ratio)
	amountB := amountA * swapConfig.ExchangeRate

	// Get recent blockhash
	recent, err := c.rpc.GetLatestBlockhash(ctx, rpc.CommitmentFinalized)
	if err != nil {
		return solana.Signature{}, fmt.Errorf("failed to get recent blockhash: %w", err)
	}

	// Create transaction instructions
	instructions := []solana.Instruction{}

	// 1. Transfer a_t from user to pool
	transferAInstruction := token.NewTransferInstruction(
		amountA,
		userAAccount,
		swapConfig.PoolATokenAccount,
		userKey.PublicKey(),
		[]solana.PublicKey{},
	).Build()
	instructions = append(instructions, transferAInstruction)

	// 2. Transfer b_t from pool to user
	transferBInstruction := token.NewTransferInstruction(
		amountB,
		swapConfig.PoolBTokenAccount,
		userBAccount,
		swapConfig.PoolAuthorityKey.PublicKey(),
		[]solana.PublicKey{},
	).Build()
	instructions = append(instructions, transferBInstruction)

	// Build transaction
	tx, err := solana.NewTransaction(
		instructions,
		recent.Value.Blockhash,
		solana.TransactionPayer(userKey.PublicKey()),
	)
	if err != nil {
		return solana.Signature{}, fmt.Errorf("failed to create transaction: %w", err)
	}

	// Sign transaction (both user and pool authority)
	_, err = tx.Sign(func(key solana.PublicKey) *solana.PrivateKey {
		if key.Equals(userKey.PublicKey()) {
			return &userKey
		}
		if key.Equals(swapConfig.PoolAuthorityKey.PublicKey()) {
			return &swapConfig.PoolAuthorityKey
		}
		return nil
	})
	if err != nil {
		return solana.Signature{}, fmt.Errorf("failed to sign transaction: %w", err)
	}

	// Send transaction
	sig, err := c.rpc.SendTransaction(ctx, tx)
	if err != nil {
		return solana.Signature{}, fmt.Errorf("failed to send transaction: %w", err)
	}

	return sig, nil
}

// CreateTokenAccount creates a new SPL token account for a given mint
func (c *Client) CreateTokenAccount(
	ctx context.Context,
	payer solana.PrivateKey,
	mint solana.PublicKey,
	owner solana.PublicKey,
) (solana.PublicKey, error) {
	// Generate new account keypair
	newAccount := solana.NewWallet()

	// Get minimum balance for rent exemption
	// SPL Token account size is 165 bytes
	const tokenAccountSize = 165
	rentExemption, err := c.rpc.GetMinimumBalanceForRentExemption(
		ctx,
		tokenAccountSize,
		rpc.CommitmentFinalized,
	)
	if err != nil {
		return solana.PublicKey{}, fmt.Errorf("failed to get rent exemption: %w", err)
	}

	// Get recent blockhash
	recent, err := c.rpc.GetLatestBlockhash(ctx, rpc.CommitmentFinalized)
	if err != nil {
		return solana.PublicKey{}, fmt.Errorf("failed to get recent blockhash: %w", err)
	}

	// Create instructions
	instructions := []solana.Instruction{}

	// 1. Create account
	createAccountInstruction := system.NewCreateAccountInstruction(
		rentExemption,
		tokenAccountSize,
		solana.TokenProgramID,
		payer.PublicKey(),
		newAccount.PublicKey(),
	).Build()
	instructions = append(instructions, createAccountInstruction)

	// 2. Initialize token account
	initAccountInstruction := token.NewInitializeAccountInstruction(
		newAccount.PublicKey(),
		mint,
		owner,
		solana.SysVarRentPubkey,
	).Build()
	instructions = append(instructions, initAccountInstruction)

	// Build and sign transaction
	tx, err := solana.NewTransaction(
		instructions,
		recent.Value.Blockhash,
		solana.TransactionPayer(payer.PublicKey()),
	)
	if err != nil {
		return solana.PublicKey{}, fmt.Errorf("failed to create transaction: %w", err)
	}

	_, err = tx.Sign(func(key solana.PublicKey) *solana.PrivateKey {
		if key.Equals(payer.PublicKey()) {
			return &payer
		}
		if key.Equals(newAccount.PublicKey()) {
			return &newAccount.PrivateKey
		}
		return nil
	})
	if err != nil {
		return solana.PublicKey{}, fmt.Errorf("failed to sign transaction: %w", err)
	}

	// Send transaction
	_, err = c.rpc.SendTransaction(ctx, tx)
	if err != nil {
		return solana.PublicKey{}, fmt.Errorf("failed to send transaction: %w", err)
	}

	return newAccount.PublicKey(), nil
}
