ARG NODE_VERSION=20
FROM node:${NODE_VERSION} AS builder

WORKDIR /usr/src/app

COPY package*.json ./
RUN npm ci

COPY . .

# Production image, copy all the files and run next
FROM node:${NODE_VERSION} AS production

WORKDIR /usr/src/app

COPY --from=builder /usr/src/app .

EXPOSE 3000

CMD ["npm", "start"]