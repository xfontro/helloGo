FROM golang:latest AS build-env

ENV APP_DIR=/app
RUN mkdir $APP_DIR
WORKDIR $APP_DIR
ADD . $APP_DIR/

RUN go build -o main .


FROM scratch

WORKDIR /app
COPY --from=build-env /app/main /app/

EXPOSE 8080

ENTRYPOINT [ "./main" ]
