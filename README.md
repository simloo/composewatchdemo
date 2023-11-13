# Compose Watch demo

This is a demo of `docker compose watch` command

1. start compose watch
```bash
docker compose watch
```
2. open http://localhost:8080
3. edit `static/index.html`
4. see the change by refresh in browser
5. open http://localhost:8080/version
6. Change version number in `internal/composewatchdemo/serve.go`
7. see the change by refresh in browser
