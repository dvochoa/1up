# 1up

A 10000 hours productivity website.

## Tech Stack

The frontend uses [NextJS](https://nextjs.org/) and TypeScript along with CSS grid for responsive design and [tailwind](https://tailwindcss.com/) for easy in-line styling.

The backend is built with Go using the [Gin framework](https://gin-gonic.com/).

## Development

After initial download you'll want to run `npm install` to install all dependencies locally. These will be available in the `node_modules` directory.

Use the following scripts during development:

**Frontend:**

- Run `npm run dev` to start a local instance of the frontend on `localhost:3000`.
- Run `npm run test` to run the entire frontend test suite.
- Run `npm run lint` to run the frontend linter.
  - The linter uses [prettier](https://prettier.io/) for code formatting

**Backend:**

- Navigate to `/backend` and run `go run backend` to start a local instance of the backend on port `8080`.

## Troubleshooting

If you experience issues that might be related to cached files (e.g. styling or other content representing previous changes) then try the following:

1. Delete your local `node_modules` and `.next` directories: `rm -rf node_modules .next`
2. Reinstall: `npm install`
3. Re-run: `npm run dev`
