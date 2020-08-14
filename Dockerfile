FROM node:14.8.0-alpine3.11 as nodejs-builder
RUN apk add --update make git
COPY . /unsee
RUN make -C /unsee webpack

FROM 1.15.0-alpine3.12 as go-builder
COPY --from=nodejs-builder /unsee /go/src/github.com/cameronkollwitz/unsee
ARG VERSION
RUN apk add --update make git
RUN CGO_ENABLED=0 make -C /go/src/github.com/cameronkollwitz/unsee VERSION="${VERSION:-dev}" unsee

FROM gcr.io/distroless/base
COPY --from=go-builder /go/src/github.com/cameronkollwitz/unsee/unsee /unsee
EXPOSE 8080
CMD ["/unsee"]
