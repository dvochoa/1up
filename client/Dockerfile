FROM node:18-alpine AS base
WORKDIR /app
# Enables use of pnpm
RUN corepack enable
# TODO: Remove this once latest pnpm is fixed (https://github.com/nodejs/corepack/issues/612
ENV COREPACK_DEFAULT_TO_LATEST=0

FROM base AS deps
COPY package.json pnpm-lock.yaml ./
RUN --mount=type=cache,id=pnpm,target=/root/.local/share/pnpm/store pnpm install --frozen-lockfile

## Run in development
FROM deps AS dev
COPY . .
EXPOSE 3000
CMD pnpm run dev

## Build in separate stage to minimize size of prod container
FROM base AS builder
COPY --from=deps /app/node_modules ./node_modules
COPY . .
ENV NEXT_TELEMETRY_DISABLED=1
RUN pnpm run build

## Run in production
FROM base AS prod
ENV NODE_ENV=production
ENV NEXT_TELEMETRY_DISABLED=1

# Add system group & user
RUN addgroup --system --gid 1001 nodejs
RUN adduser --system --uid 1001 nextjs

# Copy the bundled code from the build stage to the production image
COPY --from=builder /app/.next/standalone ./
COPY --from=builder /app/.next/static ./.next/static
COPY --from=builder /app/public ./public

# Set ownership and permission
RUN chown -R nextjs:nodejs /app

# Switch to non-root user
USER nextjs

EXPOSE 3000

# server.js is created by next build from the standalone output
# https://nextjs.org/docs/pages/api-reference/config/next-config-js/output
CMD node server.js
