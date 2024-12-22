# 1up

A 10000 hours productivity website.

## Development

Built using [NextJS](https://nextjs.org/) and TypeScript along with CSS grid for a responsive design and [tailwind](https://tailwindcss.com/) for easy in-line styling.

After initial download you'll want to run `npm install` to install all dependencies locally. These will be available in the `node_modules` directory.

Use the following scripts during development:

- Run `npm run dev` to start a local instance of the web application on `localhost:3000`.
- Run `npm run test` to run the entire test suite.
- Run `npm run lint` to run the linter.
  - The linter uses [prettier](https://prettier.io/) for code formatting

## Troubleshooting

If you experience issues that might be related to cached files (e.g. styling or other content representing previous changes) then try the following:

1. Delete your local `node_modules` and `.next` directories: `rm -rf node_modules .next`
2. Reinstall: `npm install`
3. Re-run: `npm run dev`
