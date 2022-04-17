FROM alpine:3.15

RUN mkdir -p /app

RUN mkdir /app/logs

RUN mkdir /app/conf

WORKDIR /app

ADD ./dist/main /app/main

ADD ./conf /app/conf

ENV GIN_MODE=release PORT=8082

EXPOSE 8082

CMD ["./main"]