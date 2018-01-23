FROM golang:latest
MAINTAINER Xavier Fontrodona

ENV APP_DIR=/app
RUN mkdir $APP_DIR
WORKDIR $APP_DIR
ADD ./app/* $APP_DIR/

RUN go build -o main .

EXPOSE 8080
CMD [ "/app/main" ]

