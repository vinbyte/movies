#builder
FROM golang:alpine as builder
WORKDIR /home
COPY . .
RUN go build -o movie-app app/main.go

#final image
FROM alpine
COPY --from=builder /home/movie-app .
EXPOSE 5050
CMD ./movie-app