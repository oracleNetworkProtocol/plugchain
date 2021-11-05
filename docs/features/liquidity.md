# Liquidity

## Liquidity Module

The liquidity module serves Automated Market Maker (AMM)-style decentralized liquidity by providing liquidity activities and coin swap functions.

The module enables users to create a liquidity pool, make deposits and withdrawals, and request coin swaps from the liquidity pool.

## Key features

**Combination of traditional orderbook-based model and new AMM model**

- With multiple advantages over order book-based models, the liquidity module combines a batch-based order book matching algorithm with AMM to create enriched utilities for more potential users.
- The liquidity module redefines the concept of a “swap order” in AMM as a “limit order with a short lifetime” in an order book-based exchange. By combining these concepts from two different models as one united model, the function supports both ways to participate in trading and liquidity-providing activities.
- Limit order options are not supported in the first version of the liquidity module, but the base structure of the codebase anticipates and supports feature expansion.
- Advantages of the combined model
    - More freedom on ways to provide liquidity, planned expansion for limit orders
    - The combination of pool liquidity and limit order liquidity provide users with a more enriched trading environment
