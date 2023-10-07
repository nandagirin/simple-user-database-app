FROM docker.io/library/node:18.18.0-alpine3.18 as builder

WORKDIR /app/src
COPY package*.json ./
RUN npm install
COPY . .

FROM gcr.io/distroless/nodejs:18 as runner
COPY --from=builder /app/src /app/src
WORKDIR /app/src

CMD ["./bin/www"]
