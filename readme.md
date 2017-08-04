# Running locally
```sh
make
docker run -d -p8080:8080 generator
```

Show generated map output:
```sh
curl http://127.0.0.1:8080/api/maps/generate/flat
```

