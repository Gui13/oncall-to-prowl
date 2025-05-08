# running from scratch will not include any certificates, we need to fetch them from elsewhere
FROM alpine:latest as certificates
RUN apk add --no-cache ca-certificates

FROM scratch
LABEL authors="guillaume"

COPY oncall-to-prowl /
COPY --from=certificates /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

ENTRYPOINT ["/oncall-to-prowl"]