## This is a demo app that shows information about pod, such as name, and IP, and has prometheus metrics enabled by default.

### build

```
docker build -t emedeiros/hello-app-go-sample:latest .
```

### run

```
docker run -p 8080:8080 emedeiros/hello-app-go-sample:latest
```

### checking the app
```
curl http://localhost:8080
curl http://localhost:8080/metrics
```