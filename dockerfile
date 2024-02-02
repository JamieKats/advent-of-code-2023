FROM golang:alpine3.19 AS build
WORKDIR /src
COPY . .
RUN go build -o /out/example .

# scratch is completely emnpty file system, only has /
FROM scratch AS bin
COPY --from=build /out/example /