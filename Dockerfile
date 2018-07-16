#build stage
FROM daxiang6506/tensorflow:1.8.0-go AS builder
WORKDIR /go/src/app
COPY . .
RUN go install 

#final stage
FROM daxiang6506/tensorflow:1.8.0-go
WORKDIR /app
COPY --from=builder /go/bin/app /app/app
COPY --from=builder /go/src/app /app/
ENTRYPOINT ./app
LABEL Name=tensorflow-demo Version=0.0.1
EXPOSE 1323
