FROM golang:1.22.2-bullseye

WORKDIR /app
COPY . .

ARG DB_USER=postgres
ARG DB_PASSWORD=postgres
ARG DB_NAME=query_counter
ARG DB_URL=db

ENV DB_USER=${DB_USER}
ENV DB_PASSWORD=${DB_PASSWORD}
ENV DB_NAME=${DB_NAME}
ENV DB_URL=${DB_URL}


RUN go mod download

RUN go build -o /godocker

EXPOSE 8080

CMD [ "/godocker" ]