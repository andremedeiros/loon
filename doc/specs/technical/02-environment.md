---
title: Environment
author: Andre Medeiros
timestamp: 20200608161319
status: Proposal
---

# Environment

This document proposes a specification as to how the tool manipulates the shell's environment. Specifically, how the environment can be affected from:

- The payload file
- Services
	- Configuration values (`DATABASE_URL`, `REDIS_URL`)
- Languages

The precedence is as follows:
- Inherited environment from the executing shell
- Any environment set by services
- Any environment set by languages
- Any environment set by the payload file
