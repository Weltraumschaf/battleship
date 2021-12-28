# Start by building the application.
FROM golang:1.17-bullseye as build

WORKDIR /tmp/src
ADD . /tmp/src
RUN make CGO_ENABLED=0 build

FROM gcr.io/distroless/static
COPY --from=build /tmp/src/bin/server /
CMD ["/server"]
