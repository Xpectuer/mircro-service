# Product Images

## Uploading

Note: need to use `--data-binary` to ensure file is not converted to text

```
curl -vv localhost:9090/1/go.mod -X PUT --data-binary @test.png
curl localhost:9091/1/abc.png -X POST -d @abc.png
curl localhost:
```
