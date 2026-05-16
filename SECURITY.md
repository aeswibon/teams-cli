# Security

## Supported versions

Security fixes are applied to the **default branch** (`master`) and released as new semver tags when appropriate. Use the latest tagged release or commit on `master` when possible.

## Reporting a vulnerability

**Please do not open a public issue** for security vulnerabilities.

Instead:

1. Use **[GitHub private vulnerability reporting](https://github.com/aeswibon/teams-cli/security/advisories/new)** for this repository if it is enabled for your account, **or**
2. Contact the repository maintainers privately (for example via GitHub profile contact options your organization allows).

Include:

- Description of the issue and impact  
- Steps to reproduce (if possible)  
- Affected versions or commits (if known)

We aim to acknowledge reports within a few business days and coordinate a fix and disclosure timeline with you.

## Scope notes

- This CLI handles **session tokens and configuration** on the user’s machine. Treat any bug that could leak tokens, write config to unintended locations, or weaken permissions (`config.toml`, env vars) as security-relevant.
- **Third-party services** (Microsoft Graph, chatsvc, Azure identity endpoints) are outside this repo; report product issues to the appropriate vendor when needed.
