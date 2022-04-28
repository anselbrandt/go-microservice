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
