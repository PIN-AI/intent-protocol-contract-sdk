# Documentation Navigation

This SDK provides Go language interaction capabilities with PIN Protocol smart contracts for RootLayer and Subnet.

**[中文文档](README-zh.md)** | English

## Recommended Reading Order

1. **[Quick Start](quickstart.md)** — Get started quickly with common use cases and code examples
2. **[Roles and Workflows](roles-and-workflows.md)** — Understand the Intent lifecycle and participant roles (Requester, Matcher, Agent, Validator)
3. **[Configuration & Environment](config.md)** — Environment variables, priority, and .env examples
4. **[Contract Addresses & Networks](addresses.md)** — Network configuration (Base mainnet/testnet, local dev)
5. **[Signing Specification](signing.md)** — EIP-191 signing and digest construction
6. **[Transaction Management](txmanager.md)** — Configurable TxManager (EIP-1559, nonce, replacement strategy)
7. **[API Reference](api-reference.md)** — Complete Service Layer API documentation
8. **[Best Practices](best-practices.md)** — Production recommendations and optimization tips
9. **[Troubleshooting](troubleshooting.md)** — Common issues and solutions

## Architecture Overview

```
┌─────────────────────────────────────────────────────────────┐
│                    Third-party Application                   │
└───────────────────────────┬─────────────────────────────────┘
                            │
┌───────────────────────────▼─────────────────────────────────┐
│                         SDK Client                           │
│  • Signer (EIP-191)  • TxManager (EIP-1559)  • AddressBook  │
└───────────────────────────┬─────────────────────────────────┘
                            │
┌───────────────────────────▼─────────────────────────────────┐
│                      Service Layer (7 Services)              │
│  • IntentService (22 methods)                                │
│  • AssignmentService                                         │
│  • ValidationService                                         │
│  • SubnetService (27 methods)                                │
│  • SubnetFactoryService (30 methods)                         │
│  • StakingService (21 methods)                               │
│  • CheckpointService (18 methods)                            │
└───────────────────────────┬─────────────────────────────────┘
                            │
┌───────────────────────────▼─────────────────────────────────┐
│              Contract Bindings (abigen generated)            │
│  • IntentManager  • SubnetFactory  • Subnet                  │
│  • StakingManager  • CheckpointManager                       │
└───────────────────────────┬─────────────────────────────────┘
                            │
┌───────────────────────────▼─────────────────────────────────┐
│                   Blockchain Networks                        │
│  • Base Mainnet (8453)  • Base Sepolia (84532)               │
│  • Local Development (31337)                                 │
└─────────────────────────────────────────────────────────────┘
```

## Quick Links

- **Getting Started**: Start with [Quick Start](quickstart.md) for installation and basic usage
- **Integration Guide**: See [Roles and Workflows](roles-and-workflows.md) to understand how to integrate as a specific participant type
- **API Documentation**: Complete method reference in [API Reference](api-reference.md)
- **Production Deployment**: Review [Best Practices](best-practices.md) before going live

## Additional Resources

- **Contract Bindings Generation**: Refer to main [README.md](../README.md#generating-contract-bindings-developers)
- **Example Scripts**: See `examples/` directory for working code samples
- **Environment Setup**: Use `.env.example` as a template for configuration

## Support

- **Issues**: Report bugs and feature requests on [GitHub Issues](https://github.com/PIN-AI/intent-protocol-contract-sdk/issues)
- **Questions**: For integration questions, please refer to [Troubleshooting](troubleshooting.md) first
- **Updates**: Follow the repository for the latest updates and releases
