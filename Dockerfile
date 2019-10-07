FROM golang:latest AS build-env

ENV APP_DIR=/app
RUN mkdir $APP_DIR
WORKDIR $APP_DIR
ADD . $APP_DIR/

RUN CGO_ENABLED=0 go build -o main .


FROM scratch AS production

WORKDIR /app
COPY --from=build-env /app/main /app/

EXPOSE 8080

ENTRYPOINT [ "/app/main" ]
