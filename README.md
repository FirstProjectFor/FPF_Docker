## FPF_Docker

Docker使用Demo

### Dockerfile

构建Docker镜像

```shell
docker build --build-arg envType=dev -t docker_demo:v0.0.1 .
```

运行容器

```shell
docker run docker_demo:v0.0.1
```