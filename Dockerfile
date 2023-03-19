# go编译缓存使用了BuildKit, 可以通过环境变量开启
#     DOCKER_BUILDKIT=1 docker build .

ARG GoVersion=1.18

# 编译
FROM golang:${GoVersion} AS compiler
WORKDIR /go/src/cde
RUN go env -w GOCACHE=/go-build
COPY . /go/src/cde

# 不同的tag缓存不一样, 可以分开构建
FROM compiler AS gs
RUN --mount=type=cache,sharing=shared,target=/go-build,go build -mod vendor -o bin/main github.com/cde

# 发布
FROM golang:${GoVersion}
WORKDIR /go/src/cde
ARG GitCommit=unknown
LABEL GitCommit=${GitCommit}
COPY --from=main /go/src/cde/bin ./

# 默认入口是 main
ENTRYPOINT ["./main"]