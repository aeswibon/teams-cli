# Maintainers

This document describes how the project is maintained and how decisions are made.

## Decision Making

- We prefer small, reviewable changes over large rewrites.
- Most decisions are made by consensus in PR reviews and GitHub discussions.
- If consensus is unclear, maintainers will choose the simplest option that preserves scope, safety, and maintainability.
- Security issues follow the private reporting process in `SECURITY.md`.

## Reviews And Expectations

- All non-trivial changes should go through a PR.
- PRs should be focused: one intent, minimal diff, clear description of why.
- Keep CI green: tests, formatting, and lint should pass before merge.
- Backward-incompatible changes should be explicitly called out in the PR description and changelog.

## Path To Maintainer

Maintainers are added by existing maintainers based on sustained, high-quality contributions.

Signals we look for:

- Consistent, constructive participation in issues/PRs
- Good judgment about scope and tradeoffs
- Helpful reviews and an ability to unblock others
- Care for security, docs, and testing (not only code changes)

Typical steps:

- Start by contributing small PRs (docs, tests, fixes)
- Help triage issues and reproduce bugs
- Review PRs and propose improvements with empathy and precision
- After a track record is established, a maintainer may invite you to join
