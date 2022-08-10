# @bitnetwork/core-utils

## What is this?

`@bitnetwork/core-utils` contains the Bit Virtual Machine core utilities.

## Getting started

### Building and usage

After cloning and switching to the repository, install dependencies:

```bash
$ yarn
```

Use the following commands to build, use, test, and lint:

```bash
$ yarn build
$ yarn start
$ yarn test
$ yarn lint
```

### L2 Fees

`TxGasLimit` can be used to `encode` and `decode` the L2 Gas Limit
locally.

```typescript
import { TxGasLimit } from '@bitnetwork/core-utils'
import { JsonRpcProvider } from 'ethers'

const L2Provider = new JsonRpcProvider('https://mainnet.optimism.io')
const L1Provider = new JsonRpcProvider('http://127.0.0.1:8545')

const l2GasLimit = await L2Provider.send('eth_estimateExecutionGas', [tx])
const l1GasPrice = await L1Provider.getGasPrice()

const encoded = TxGasLimit.encode({
  data: '0x',
  l1GasPrice,
  l2GasLimit,
  l2GasPrice: 10000000,
})

const decoded = TxGasLimit.decode(encoded)
assert(decoded.eq(gasLimit))

const estimate = await L2Provider.estimateGas()
assert(estimate.eq(encoded))
```