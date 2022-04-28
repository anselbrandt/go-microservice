# Product Images

## Uploading

`--data-binary` ensures file is not converted to text

```
curl -vv localhost:9090/1/go.mod -X PUT --data-binary @test.png
```

```
curl -v localhost:9090/images/1/test.png --data-binary @test.png
```

```
curl localhost:9090/images/1/test.png --output downloaded.png
```

Accept gzipped content:

```
curl -v localhost:9090/images/1/test.png --compressed --output downloaded.png
```

Should result in the following request headers:

```
> GET /images/1/test.png HTTP/1.1
> Host: localhost:9090
> User-Agent: curl/7.79.1
> Accept: */*
> Accept-Encoding: deflate, gzip
```
