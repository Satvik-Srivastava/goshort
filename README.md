# GoShort

GoShort is a lightweight URL shortener: a Go/Fiber backend with Redis storage, and a React + Tailwind frontend served by Nginx in Docker.

## What we built

The stack now includes a complete frontend that connects to the existing backend API. The system provides:
- `POST /api/v1/shorten` — create a short URL
- `GET /:url` — resolve the short URL and redirect to the original target
- Rate limiting using Redis
- Expiry support for shortened links
- A minimal, interactive React frontend (paste a URL, optionally provide a custom short slug, choose expiry, get a short link)

## Frontend

- Built with React 19 and Tailwind CSS (v3).
- Built UI features: animated inputs, loading state, success/error notifications, copy-to-clipboard, open-in-new-tab, and a compact rate-limit display.
- Served by Nginx (default at port `3001`) with the API proxied at `/api/` → `http://api:3000` inside Docker.

## What was fixed (summary)

- Docker build and user creation issues in the API image were fixed (correct `adduser` usage).
- Go build now runs inside the correct `WORKDIR` so `go.mod` is found.
- Port normalization: `.env` `APP_PORT` was normalized to avoid double-colon listen addresses.
- Redis lookup logic for redirects was fixed so successful lookups redirect correctly.
- Custom `custom_short` values are normalized to store only the slug portion.

## How to run the full stack (local, Docker Compose)

From the project root you can build and run all services (api, frontend, redis) with one command. This will build images and start containers:

```bash
sudo docker compose up -d --build
```

Services and default ports:
- Frontend (Nginx serving the React app): http://localhost:3001
- API (Go/Fiber): http://localhost:3000
- Redis: 6379 (internal container, exposed when using compose)

If you only change frontend code and want to rebuild it:

```bash
sudo docker compose build frontend --no-cache
sudo docker compose up -d frontend
```

Example API request (create short URL):

```bash
curl -X POST http://localhost:3000/api/v1/shorten \
  -H "Content-Type: application/json" \
  -d '{"url":"https://example.com", "expiry":24}'
```

The response contains the shortened URL and metadata.

## Development notes

- Frontend is compatible with React 19; the app uses `ReactDOM.createRoot(...)`.
- Tailwind v3 is used to avoid PostCSS plugin compatibility issues.
- Nginx config contains the SPA fallback (`try_files`) and proxies `/api/` to the API service.

## Future enhancements

- Analytics and click counts
- User accounts to manage links
- QR code generation and preview cards
- Vanity domains and custom domains support

## Contributing / Workflow

1. Create a branch for your change: `git checkout -b feature/your-change`.
2. Make changes and run `sudo docker compose up -d --build` to test locally.
3. Commit and push your branch, then open a pull request.

---

If you want, I can push this README update to GitHub on a branch and open a PR for you.
