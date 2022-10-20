# Transaction Management with Rest-API

It is the implementation to API in https://infra.devskills.app/transaction-management/api/3.1.0

<details>
<summary>use the manual setup.</summary>
1. Update the `apiUrl` (where your backend runs) in [cypress.json](cypress.json).
2. Update the [`build`](package.json#L5) and [`start`](package.json#L6) scripts in [package.json](package.json) to respectively build and start your app.
</details>

## Running the API tests

<details>
<summary>Locally with Docker (Mac & Windows only)</summary>
## Prerequisites

- [Install Docker](https://www.docker.com/get-started)
- Start your app

### Run the tests

```bash
 docker run --add-host host.docker.internal:host-gateway -v $PWD:/e2e -w /e2e cypress/included:3.4.0
```

</details>

<details>
<summary>Locally with npm</summary>
  
#### Prerequisites

1. [Install node](https://nodejs.org/en/)
2. When in the project's root, run: `sed 's/host.docker.internal/localhost/g' cypress.json > cypress.json.tmp && mv cypress.json.tmp cypress.json`
3. Start your app

### Run the tests

```bash
 npm run test
```

</details>

## brief explanation

1. The database used is [SQLite](https://www.sqlite.org/index.html) database.
2. Cache is enabled to ensure the service GET endpoints do not slow down as the database size grows.
3. transaction is used to ensure no data lost due to a race condition when creating a new transaction on the server.
4. Unit test a couple of modules of choice.
