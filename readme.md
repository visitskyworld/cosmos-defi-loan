# DeFi Loan

**DeFi loan** is a peer to peer loan blockchain built using Cosmos SDK and Tendermint and created with [Starport](https://starport.com). The blockchain allows you to create loan request. With the loan request you must submit the loan amount, fee you are willing to pay for the loan, and the amount of collateral that will back the loan. Once the loan request is created anyone can choose to approve your loan and lend you tokens if they agree to your collateral and fee terms. If you don't pay the loan back in a certain time it can be liqudated by the user that lent you the tokens. If the loan is liquidated all collateral is sent to the user\address that lent you the tokens. You can also cancel a loan request before it has been approved. This is simply a proof-of-concept.

# Readme file in development

## Get started

```
starport chain serve
```

`serve` command installs dependencies, builds, initializes, and starts your blockchain in development.

### Configure

Your blockchain in development can be configured with `config.yml`. To learn more, see the [Starport docs](https://docs.starport.com).

### Web Frontend

Starport has scaffolded a Vue.js-based web app in the `vue` directory. Run the following commands to install dependencies and start the app:

```
cd vue
npm install
npm run serve
```

The frontend app is built using the `@starport/vue` and `@starport/vuex` packages. For details, see the [monorepo for Starport front-end development](https://github.com/tendermint/vue).

## Release

To release a new version of your blockchain, create and push a new tag with `v` prefix. A new draft release with the configured targets will be created.

```
git tag v0.1
git push origin v0.1
```

After a draft release is created, make your final changes from the release page and publish it.

### Install

To install the latest version of your blockchain node's binary, execute the following command on your machine:

```
curl https://get.starport.com/sheldonlsides/loan@latest! | sudo bash
```

`sheldonlsides/loan` should match the `username` and `repo_name` of the Github repository to which the source code was pushed. Learn more about [the install process](https://github.com/allinbits/starport-installer).

## Learn more

- [Starport](https://starport.com)
- [Tutorials](https://docs.starport.com/guide)
- [Starport docs](https://docs.starport.com)
- [Cosmos SDK docs](https://docs.cosmos.network)
- [Developer Chat](https://discord.gg/H6wGTY8sxw)
