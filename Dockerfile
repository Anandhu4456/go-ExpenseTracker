FROM golang:1.21-alpine3.18 as build
WORKDIR /project
COPY . ./
RUN go mod download
RUN go build -v -o /api ./cmd/api

FROM gcr.io/distroless/static-debian11
COPY --from=build /api /api
COPY --from=build /project/.env /
EXPOSE 3000
CMD [ "/api" ]