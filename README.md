# Slow Go HTTP Server

I needed a simple HTTP server that would take a long time to serve responses

## Run

```bash
go run slow-server.go
```

## Example

```
http://localhost:3000/       # Usage example /foo/5
http://localhost:3000/pugs   # I love pugs! (no delay)
http://localhost:3000/pugs/5 # I love pugs! (delayed by 5 seconds)
```
