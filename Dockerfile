# TODO(@kylejb): revisit directory structure to simplify embedding of web and set non-root user for security
FROM node:lts-alpine AS web_builder

RUN npm i -g @go-task/cli

WORKDIR /app

COPY .taskfiles .taskfiles/
COPY Taskfile.yml ./
COPY web web/

RUN task web:deps
RUN task web:build


FROM golang:1-alpine AS api_builder

RUN go install github.com/go-task/task/v3/cmd/task@latest

WORKDIR /app

COPY . .
COPY --from=web_builder /app/web/dist /app/web/dist/

RUN task api:build


FROM alpine AS production

RUN apk upgrade --no-cache

WORKDIR /

COPY --from=api_builder /app/api /app/swellhub

EXPOSE 4000

CMD ["/app/swellhub"]
