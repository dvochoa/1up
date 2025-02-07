FROM node:18-alpine AS base
WORKDIR /app
COPY package*.json ./
RUN npm ci
COPY . .

# Run in development
FROM base AS dev
EXPOSE 3000
CMD ["npm", "run", "dev"]

# Run in production
FROM base AS prod
RUN npm run build
EXPOSE 3000
CMD ["npm", "run", "start"]
