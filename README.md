# GoShort

GoShort is a simple URL shortening service built with Go, Fiber, Redis, and Docker.

## What we built

This project provides a backend API that accepts a long URL and returns a shortened URL. When a user opens the shortened link, the service redirects to the original URL.

The service includes:
- `POST /api/v1/shorten` to create a short URL
- `GET /:url` to resolve the short URL and redirect to the original target
- rate limiting using Redis
- expiry support for shortened links

## What was wrong and how it was fixed

While building and running the project, several issues were identified and resolved:

1. Docker build errors
   - The API Dockerfile used an invalid `adduser` command in the second stage. This was fixed by creating a proper user with `adduser -S -D -H appuser`.
   - The Go build stage was running `go build` outside of the copied source directory. Adding `WORKDIR /build` before `go build` ensured the module and source files were available.

2. Port configuration bug
   - `APP_PORT` in `.env` was set to `:3000`, which caused Fiber to build an invalid listen address like `::3000`. The value was normalized to `3000` and the code now strips any leading colon before listening.

3. Redirect handler logic bug
   - The resolve route incorrectly treated every successful Redis lookup as if the short code was missing. The lookup logic was corrected so the service only returns `404` when Redis returns `redis.Nil`.

4. Custom short URL normalization
   - User-provided `custom_short` values like `localhost:3000/aa99fe` were converted to just the short slug `aa99fe`, so the stored key and returned redirect URL now work consistently.

## Technologies used

- Go
- Fiber web framework
- Redis (for storage and rate limiting)
- Docker
- Docker Compose
- Alpine Linux base images
- Go modules (`go.mod`)

## How it works

1. The client sends a JSON payload to `POST /api/v1/shorten`.
2. The service validates the URL and checks rate limits.
3. It creates a short key and stores the mapping in Redis.
4. The response includes a shortened URL.
5. When a client requests the short URL, the service looks it up in Redis and redirects to the original URL.

## Setup and run

From the project root:

```bash
sudo docker compose build api
sudo docker compose up -d
```

Then test the API with Postman or curl:

```bash
curl -X POST http://localhost:3000/api/v1/shorten \
  -H "Content-Type: application/json" \
  -d '{"url":"https://example.com"}'
```

The response will include a `custom_short` or generated short link. Open that link in a browser to verify redirection.

## Future scaling ideas

This project can scale in several ways:

- Move from a single Redis instance to a Redis cluster for higher availability and larger dataset support.
- Add persistent storage for analytics and long-term history beyond Redis expiry.
- Add authentication and user-specific URLs so users can manage their own shortened links.
- Use a dedicated load balancer and horizontally scale API containers.
- Add metrics, monitoring, and observability for production readiness.
- Support a custom domain service and vanity URL management.

## Notes

- The current implementation is focused on backend URL shortening and redirection.
- The dockerized setup makes it easy to run locally.
- The project is ready for extension with frontend, analytics, and production-grade deployment.
