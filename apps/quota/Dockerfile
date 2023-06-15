FROM golang:1.19-buster AS build

WORKDIR /app

#COPY ./payment/go.mod ./
#COPY ./payment/go.sum ./
#COPY . .
#RUN go mod download
COPY . .
RUN go build ./payment/src/main.go
RUN echo "$PWD"

##
## Deploy
##
FROM gcr.io/distroless/base-debian10

WORKDIR /src

COPY --from=build /app/main ./main

EXPOSE 8080


ENTRYPOINT ["/src/main"]
