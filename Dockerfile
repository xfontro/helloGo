FROM golang:latest

ENV APP_DIR=/app
RUN mkdir $APP_DIR
WORKDIR $APP_DIR
ADD . $APP_DIR/

RUN go build -o main .

EXPOSE 8080
CMD [ "/app/main" ]

