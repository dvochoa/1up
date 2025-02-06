FROM node:18-alpine as base
WORKDIR /app
COPY package*.json ./
RUN npm ci
COPY . .
EXPOSE 3000

# Run in development
FROM base as dev
CMD ["npm", "run", "fart"]

# Run in production
FROM base as prod
RUN npm run build
CMD ["npm", "run", "start"]
