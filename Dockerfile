FROM golang:1.11-alpine AS build
COPY ./ /go/src/study-cicd/
WORKDIR /go/src/study-cicd/
RUN go build -o study-cicd

FROM alpine:3.7
#RUN apk --no-cache --update upgrade \
#  && apk add --update --no-cache tzdata ca-certificates \
#  && update-ca-certificates --fresh
RUN mkdir -p /app
COPY --from=build /go/src/study-cicd/study-cicd /app/study-cicd

EXPOSE 80
EXPOSE 443

ENTRYPOINT ["/app/study-cicd"]