# 1up

A 10000 hours productivity website.

## Tech Stack

The frontend uses [NextJS](https://nextjs.org/) and TypeScript along with CSS grid for responsive design and [tailwind](https://tailwindcss.com/) for easy in-line styling.

The backend is built with Go using the [Gin framework](https://gin-gonic.com/).

## Development

### Prerequsites
Make sure you have the following installed locally:
- [go](https://go.dev/doc/install)
- [npm](https://nodejs.org/en/download)
- [docker desktop](https://www.docker.com/products/docker-desktop/)

Frontend code is placed in the `./client/` directory whereas backend code is placed in the `./backend/` directory.

You can deploy the site locally by running `docker compose up` after which it will be accessible from `localhost:3000`.

Use the following scripts during development:

**Frontend:**

Note: All of these commands must be ran from the `./client/` directory.

After initial download you'll want to run `npm install` to install all client dependencies locally. These will be available in the `node_modules` directory.

- Run `npm run dev` to deploy the frontend locally on `localhost:3000`.
- Run `npm run test` to run the entire frontend test suite.
- Run `npm run lint` to run the frontend linter.
  - The linter uses [prettier](https://prettier.io/) for code formatting

**Backend:**

Note: All of these commands must be ran from the `./backend/` directory.

- Run `go run main.go` to deploy the backend locally on `localhost:8080`.
- Run `go test ./...` to run the entire backend test suite.

### CI
This repo uses [Github Actions](https://github.com/features/actions) to configure CI workflows that enforce testing and code style compliance.

You can set up automatic code formatting on commit using a [pre-commit hook](https://github.com/dvochoa/1up/tree/main/.githooks/pre-commit).

To set up the hook, run:

```shell
git config core.hooksPath .githooks
```

**Note**: The pre-commit hook will add some noticeable latency to each commit and changing the hooksPath in your local git config will override the directory which looks to for hooks. If you have any pre-existing hooks that you still want to use, instead of changing the hooksPath you can copy the contents of any hooks in `.githooks/` to your existing local hooks directory at `.git/hooks`.

## Troubleshooting

If you experience issues that might be related to cached files (e.g. styling or other content representing previous changes) then try the following:

1. `cd client`
2. Delete your local `node_modules` and `.next` directories: `rm -rf node_modules .next`
3. Reinstall: `npm install`
4. Re-run: `npm run dev`
