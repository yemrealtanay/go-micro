FROM alpine:latest

RUN mkdir /app

COPY openAIServiceApp /app
COPY .env /app


CMD [ "/app/openAIServiceApp" ]

