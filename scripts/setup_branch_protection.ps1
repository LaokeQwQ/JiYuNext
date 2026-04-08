param(
  [string]$Owner = "LaokeQwQ",
  [string]$Repo = "JiYuNext"
)

$gh = "C:\Program Files\GitHub CLI\gh.exe"
if (-not (Test-Path $gh)) {
  throw "gh CLI not found at: $gh"
}

$status = & $gh auth status 2>$null
if ($LASTEXITCODE -ne 0) {
  throw "gh is not authenticated. Run: gh auth login"
}

$requiredChecks = @(
  "PR Guard / validate-branch-flow",
  "PR CI / test-and-lint"
)

$branches = @("dev", "canary", "release")

foreach ($branch in $branches) {
  Write-Host "Applying branch protection on $branch ..."

  $body = @{
    required_status_checks = @{
      strict   = $true
      contexts = $requiredChecks
    }
    enforce_admins = $true
    required_pull_request_reviews = @{
      dismiss_stale_reviews           = $true
      require_code_owner_reviews      = $false
      required_approving_review_count = 1
    }
    restrictions                    = $null
    required_linear_history         = $true
    allow_force_pushes              = $false
    allow_deletions                 = $false
    block_creations                 = $false
    required_conversation_resolution = $true
    lock_branch                     = $false
    allow_fork_syncing              = $true
  } | ConvertTo-Json -Depth 6

  $tmp = New-TemporaryFile
  Set-Content -LiteralPath $tmp.FullName -Value $body -Encoding utf8

  & $gh api `
    --method PUT `
    -H "Accept: application/vnd.github+json" `
    "/repos/$Owner/$Repo/branches/$branch/protection" `
    --input $tmp.FullName

  $code = $LASTEXITCODE
  Remove-Item -LiteralPath $tmp.FullName -Force
  if ($code -ne 0) {
    throw "Failed to apply protection for branch: $branch"
  }
}

Write-Host "Branch protection applied successfully for dev/canary/release."
