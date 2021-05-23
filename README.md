# Smart Contract

## Requirements
- [Truffle](https://www.trufflesuite.com/docs/truffle/overview)
- [Ganache CLI](https://github.com/trufflesuite/ganache-cli)
- Node: v8.9.4 or later
- [Jest](https://jestjs.io/) or [Mocha](https://mochajs.org/)

## Configuration
- truffle.config.js

## Initialize project
```
truffle init
```

## Compile constracts
```
npm run compile
```

## Deploy constracts
```
npx truffle migrate --network test
npm run migrate-testnet
```

## Run test
```
npx truffle test --network test
```