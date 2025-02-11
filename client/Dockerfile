FROM node:18-alpine AS base
WORKDIR /app

FROM base AS deps
COPY package*.json ./
RUN npm ci

# Run in development
FROM deps AS dev
# TODO: is this needed?
COPY . .
EXPOSE 3000
CMD ["npm", "run", "dev"]

# Build in separate stage to minimize size of prod container
FROM base AS builder
COPY --from=deps /app/node_modules ./node_modules
COPY . .
RUN npm run build

# Run in production
FROM base AS prod

# TODO: What does this do?
ENV NODE_ENV=production

# TODO: What does this do?
RUN addgroup --system --gid 1001 nodejs
RUN adduser --system --uid 1001 nextjs

# TODO: What does this do
# Copy the bundled code from the build stage to the production image
COPY --from=builder /app/.next/standalone ./
COPY --from=builder /app/.next/static ./.next/static

# TODO: What does this do?
# Set correct permissions
RUN chown -R nextjs:nodejs /app

# TODO: What does this do?
USER nextjs

EXPOSE 3000

# server.js is created by next build from the standalone output
# https://nextjs.org/docs/pages/api-reference/config/next-config-js/output
CMD ["node", "server.js"]
