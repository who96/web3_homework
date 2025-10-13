use solana_program::{
    account_info::{next_account_info, AccountInfo},
    entrypoint,
    entrypoint::ProgramResult,
    msg,
    program::{invoke, invoke_signed},
    program_error::ProgramError,
    program_pack::Pack,
    pubkey::Pubkey,
    sysvar::{rent::Rent, Sysvar},
};
use spl_token::state::Account as TokenAccount;

entrypoint!(process_instruction);

/// Fixed exchange rate: 1 a_t = 100 b_t
const EXCHANGE_RATE: u64 = 100;

/// Instruction enum
pub enum TokenSwapInstruction {
    /// Swap a_t for b_t
    /// Accounts expected:
    /// 0. `[writable]` User's a_t token account (source)
    /// 1. `[writable]` User's b_t token account (destination)
    /// 2. `[writable]` Pool's a_t token account
    /// 3. `[writable]` Pool's b_t token account
    /// 4. `[signer]` User authority
    /// 5. `[]` Token program
    SwapAtoBtoB { amount_a: u64 },
}

impl TokenSwapInstruction {
    pub fn unpack(input: &[u8]) -> Result<Self, ProgramError> {
        let (&tag, rest) = input.split_first().ok_or(ProgramError::InvalidInstructionData)?;

        match tag {
            0 => {
                if rest.len() != 8 {
                    return Err(ProgramError::InvalidInstructionData);
                }
                let amount_a = u64::from_le_bytes(rest[0..8].try_into().unwrap());
                Ok(Self::SwapAtoBtoB { amount_a })
            }
            _ => Err(ProgramError::InvalidInstructionData),
        }
    }
}

pub fn process_instruction(
    program_id: &Pubkey,
    accounts: &[AccountInfo],
    instruction_data: &[u8],
) -> ProgramResult {
    let instruction = TokenSwapInstruction::unpack(instruction_data)?;

    match instruction {
        TokenSwapInstruction::SwapAtoBtoB { amount_a } => {
            msg!("Instruction: Swap a_t to b_t");
            process_swap(program_id, accounts, amount_a)
        }
    }
}

fn process_swap(
    _program_id: &Pubkey,
    accounts: &[AccountInfo],
    amount_a: u64,
) -> ProgramResult {
    let account_info_iter = &mut accounts.iter();

    let user_a_account = next_account_info(account_info_iter)?;  // User's a_t account
    let user_b_account = next_account_info(account_info_iter)?;  // User's b_t account
    let pool_a_account = next_account_info(account_info_iter)?;  // Pool's a_t account
    let pool_b_account = next_account_info(account_info_iter)?;  // Pool's b_t account
    let user_authority = next_account_info(account_info_iter)?;  // User (signer)
    let token_program = next_account_info(account_info_iter)?;   // Token program

    // Verify user is signer
    if !user_authority.is_signer {
        msg!("Error: User authority must be a signer");
        return Err(ProgramError::MissingRequiredSignature);
    }

    // Calculate amount of b_t to send
    let amount_b = amount_a
        .checked_mul(EXCHANGE_RATE)
        .ok_or(ProgramError::InvalidArgument)?;

    msg!("Swapping {} a_t for {} b_t", amount_a, amount_b);

    // Transfer a_t from user to pool
    let transfer_a_to_pool_ix = spl_token::instruction::transfer(
        token_program.key,
        user_a_account.key,
        pool_a_account.key,
        user_authority.key,
        &[],
        amount_a,
    )?;

    invoke(
        &transfer_a_to_pool_ix,
        &[
            user_a_account.clone(),
            pool_a_account.clone(),
            user_authority.clone(),
            token_program.clone(),
        ],
    )?;

    msg!("Transferred {} a_t from user to pool", amount_a);

    // Transfer b_t from pool to user
    // Note: In production, pool authority should be a PDA controlled by this program
    // For simplicity, we assume pool_authority is available in accounts[6]
    let pool_authority = next_account_info(account_info_iter)?;

    if !pool_authority.is_signer {
        msg!("Error: Pool authority must be a signer");
        return Err(ProgramError::MissingRequiredSignature);
    }

    let transfer_b_to_user_ix = spl_token::instruction::transfer(
        token_program.key,
        pool_b_account.key,
        user_b_account.key,
        pool_authority.key,
        &[],
        amount_b,
    )?;

    invoke(
        &transfer_b_to_user_ix,
        &[
            pool_b_account.clone(),
            user_b_account.clone(),
            pool_authority.clone(),
            token_program.clone(),
        ],
    )?;

    msg!("Transferred {} b_t from pool to user", amount_b);
    msg!("Swap completed successfully!");

    Ok(())
}
