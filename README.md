# 1up

A 10000 hours productivity website.

## Tech Stack

The frontend uses [NextJS](https://nextjs.org/) and TypeScript along with CSS grid for responsive design and [tailwind](https://tailwindcss.com/) for easy in-line styling.

The backend is built with Go using the [Gin framework](https://gin-gonic.com/).

## Development

### Prerequisites
Make sure you have the following installed locally:
- [go](https://go.dev/doc/install)
- [pnpm](https://pnpm.io/installation)
- [docker](https://www.docker.com/products/docker-desktop/)

Frontend code is placed in the `./client/` directory whereas backend code is placed in the `./backend/` directory.

You can deploy the site locally by running: 
- `docker compose up --build`: For production
- `docker compose -f compose.yaml -f compose.dev.yaml up --build`: For development

Development will come with helpful features such as hot reloading but will be a more bloated build.

In either case, after deploying the site will be accessible from `localhost` or the IP address of your local network if accessing from a machine other than the one used to run docker.

Use the following scripts during development:

**Frontend:**

Note: All of these commands must be ran from the `./client/` directory.

After initial download you'll want to run `pnpm install` to install all client dependencies locally. These will be available in the `node_modules` directory.

- Run `pnpm run test` to run the entire frontend test suite.
- Run `pnpm run lint` to run the frontend linter.
  - The linter uses [prettier](https://prettier.io/) for code formatting

**Backend:**

Note: All of these commands must be ran from the `./backend/` directory.

- Run `go test ./...` to run the entire backend test suite.

### CI
This repo uses [Github Actions](https://github.com/features/actions) to configure CI workflows that enforce testing and code style compliance.

You can set up automatic code formatting on commit using a [pre-commit hook](https://github.com/dvochoa/1up/tree/main/.githooks/pre-commit).

To set up the hook, run:

```shell
git config core.hooksPath .githooks
```

**Note**: The pre-commit hook will add some noticeable latency to each commit and changing the hooksPath in your local git config will override the directory which looks to for hooks. If you have any pre-existing hooks that you still want to use, instead of changing the hooksPath you can copy the contents of any hooks in `.githooks/` to your existing local hooks directory at `.git/hooks`.
