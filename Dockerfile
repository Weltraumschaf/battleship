# Start by building the application.
FROM golang:1.17-bullseye as build

WORKDIR /tmp/src
ADD . /tmp/src
RUN make

FROM gcr.io/distroless/base-debian11
COPY --from=build /tmp/src/bin/server /
CMD ["/server"]
