Docker image
---
Current image is based on Alpine and uses runit for supervising simple test Go service app.

[App source code](test-service/main.go)

[Dockerfile](test-service/Dockerfile)

[Image on Docker hub](https://hub.docker.com/r/vitaliihurin/test-service/)



Test on my AWS EC2
---
http://54.93.56.17:8686/test



Local setup
---
```bash
sudo docker run --name test-service -p 8686:80 -d vitaliihurin/test-service
curl http://localhost:8686/test
```
