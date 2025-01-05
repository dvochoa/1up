# 1up

A 10000 hours productivity website.

## Tech Stack

The frontend uses [NextJS](https://nextjs.org/) and TypeScript along with CSS grid for responsive design and [tailwind](https://tailwindcss.com/) for easy in-line styling.

The backend is built with Go using the [Gin framework](https://gin-gonic.com/).

## Development

Frontend code is placed in the `./frontend/` directory whereas backend code is placed in the `./backend/` directory.

After initial download you'll want to `cd frontend` and run `npm install` to install all frontend dependencies locally. These will be available in the `node_modules` directory.

Use the following scripts during development:

**Frontend:**

Note: All of these commands must be ran from the `./frontend/` directory.

- Run `npm run dev` to start a local instance of the frontend on `localhost:3000`.
- Run `npm run test` to run the entire frontend test suite.
- Run `npm run lint` to run the frontend linter.
  - The linter uses [prettier](https://prettier.io/) for code formatting

**Backend:**

Note: All of these commands must be ran from the `./backend/` directory.

- Run `go run backend` to start a local instance of the backend on port `8080`.
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

1. `cd frontend`
2. Delete your local `node_modules` and `.next` directories: `rm -rf node_modules .next`
3. Reinstall: `npm install`
4. Re-run: `npm run dev`
