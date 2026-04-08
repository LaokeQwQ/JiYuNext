# Branch Policy

## Branches
- `dev`: daily integration branch.
- `canary`: pre-release verification branch.
- `release`: production release branch.

## PR Flow Rules
- PR to `dev`: allowed from feature/fix/chore branches (not from `canary`/`release`).
- PR to `canary`: only from `dev`.
- PR to `release`: only from `canary`.

## Required Checks
- `PR Guard / validate-branch-flow`
- `PR CI / test-and-lint`

## Protection Baseline
- Require pull request before merge.
- Require 1 approving review.
- Dismiss stale reviews on new commits.
- Require conversation resolution before merge.
- No force push, no delete.
- Require linear history.

## Apply Protection
After `gh auth login`, run:

```powershell
powershell -ExecutionPolicy Bypass -File scripts/setup_branch_protection.ps1
```
