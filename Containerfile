FROM registry.access.redhat.com/ubi8/ubi AS builder
RUN dnf install -y go
WORKDIR /opt/
COPY goping.go .
RUN ["go", "build", "goping.go"]

FROM registry.access.redhat.com/ubi8/ubi-minimal
WORKDIR /opt/
COPY --from=builder /opt/goping .
ENTRYPOINT ["./goping"]
