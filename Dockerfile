FROM alpine:3.20.1 as builder

WORKDIR /go/src/github.com/systemli/dereferrer

ENV USER=appuser
ENV UID=10001

RUN adduser \
    --disabled-password \
    --gecos "" \
    --home "/nonexistent" \
    --shell "/sbin/nologin" \
    --no-create-home \
    --uid "${UID}" \
    "${USER}"

FROM scratch

COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group
COPY dereferrer /dereferrer

USER appuser:appuser

EXPOSE 8080

ENTRYPOINT ["/dereferrer"]
