---
name: Bug report
about: Report something that does not work as documented or expected
title: ''
labels: bug
---

## Summary

<!-- Clear short description -->

## Environment

- teams-cli version (or commit): 
- OS (e.g. macOS 15, Ubuntu 24.04): 
- Go version (`go version`): 
- API backend (`graph` / `chatsvc` from config or `teams-cli doctor`): 
- Token type (required: `delegated` / `app-only` (MSAL)): 

### Graph request details (required when API backend is `graph`)

- Graph endpoint path being called (required, e.g. `/v1.0/me/joinedTeams`):
- HTTP status (required, e.g. `401`, `403`, `429`):
- `request-id` response header (required if present):

## Steps to reproduce

1. 
2. 

## Expected behavior

## Actual behavior

## Configuration / redacted details

<!-- Do NOT paste real tokens. You may redact IDs and show only shape, e.g. `api = "chatsvc"`. -->

- [ ] I confirm I did not paste real tokens/secrets.

```text
(paste redacted doctor output or config fields if useful)
```

## Notes

<!-- Optional: anything else that helps narrow this down -->
