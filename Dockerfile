FROM golang:1.20

WORKDIR /app
COPY . /app

RUN go build -o generator ./cmd

RUN cp /app/generator /generator
RUN rm -rf /app

CMD /generator
