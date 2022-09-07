# Transaction Management Backend - Level 3

Your task is to **build a backend app** that **fulfills the [Transaction Management API](https://infra.devskills.app/transaction-management/api/3.1.0)** and **make the provided API tests pass**.

Please agree with your hiring team regarding the tech stack choice.

## Before you get started

### If you run into a problem

Need help? Head over to [our community on GitHub](https://github.com/orgs/DevSkillsHQ/discussions/categories/help) to get assistance.

### Import boilerplate

Follow [this link](https://docs.devskills.co/collections/85-the-interview-process/articles/342-importing-challenge-boilerplate) to get the boilerplate code for your tech stack to configure a minimal setup for running the E2E tests.

<details>
<summary>Alternatively, use the manual setup.</summary>

1. Update the `apiUrl` (where your backend runs) in [cypress.json](cypress.json).
2. Update the [`build`](package.json#L5) and [`start`](package.json#L6) scripts in [package.json](package.json) to respectively build and start your app.

</details>

### Get familiar with the API

Follow [this link](https://infra.devskills.app/transaction-management/api/3.1.0) to find the API documentation. Feel free to try out a few requests to better understand how the API should work.

### Try running the API tests

<details>
<summary>Remotely on the pipeline</summary>

Create and switch to a new `implementation` branch and push your code. This will trigger a new pipeline run which will execute the tests.
  
Check the 'Actions' tab to see the historical runs.

</details>


<details>
<summary>Locally with Docker (Mac & Windows only)</summary>
  
#### Prerequisites

- [Install Docker](https://www.docker.com/get-started)
- Start your app
  
#### Run the tests
```bash
 docker run --add-host host.docker.internal:host-gateway -v $PWD:/e2e -w /e2e cypress/included:3.4.0
```

You can either use the console output or generated screenshots/videos (*check the newly created files that appear after a test run*) to troubleshoot the test results.


</details>

<details>
<summary>Locally with npm</summary>
  
#### Prerequisites

1. [Install node](https://nodejs.org/en/)
2. When in the project's root, run: `sed 's/host.docker.internal/localhost/g' cypress.json > cypress.json.tmp && mv cypress.json.tmp cypress.json`  
3. Start your app
  
#### Run the tests
```bash
 npm run test
```

You can either use the console output or generated screenshots/videos (*check the newly created files that appear after a test run*) to troubleshoot the test results.

</details>


### What we expect from you

1. Make the provided API tests pass.
2. Keep server data in a [SQLite](https://www.sqlite.org/index.html) database. We want to see how you design the database schema and SQL queries.
3. Ensure that the service GET endpoints do not slow down as the database size grows.
4. Ensure that no data can be lost due to a race condition when creating a new transaction on the server.
5. Avoid duplication and extract re-usable modules where it makes sense. We want to see your approach to creating a codebase that is easy to maintain.
6. Unit test one module of choice. There is no need to test the whole app, as we only want to understand what you take into consideration when writing unit tests.
7. Push your code to the new `implementation` branch. We encourage you to commit and push your changes regularly as it's a good way for you to showcase your thinking process.
8. Create a new pull request, but please **do not merge it**.
9. Document the tech decisions you've made by creating a new review on your PR ([see how](https://www.loom.com/share/94ae305e7fbf45d592099ac9f40d4274)). In particular, document how you implemented the mechanism that protects the service against race conditions.
10. Await further instructions from the hiring team.

## Time estimate

About **3 hours** depending on your experience level + the time to set up the project/environment (go with one of the provided boilerplates to move faster).

Also, there is no countdown. The estimate is for you to plan your time.

---

Authored by [Rafa Paliwoda](https://devskills.co/authors/rafa-paliwoda) via [DevSkills](https://devskills.co).

How was your experience? **[Don't hesitate to give us a shout](https://github.com/orgs/DevSkillsHQ/discussions/categories/feedback)**.
