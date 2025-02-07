FROM node:18-alpine as base
WORKDIR /app
COPY package*.json ./
RUN npm ci
COPY . .

# Run in development
FROM base as dev
EXPOSE 3000
CMD ["npm", "run", "dev"]

# Run in production
FROM base as prod
RUN npm run build
EXPOSE 3000
CMD ["npm", "run", "start"]
