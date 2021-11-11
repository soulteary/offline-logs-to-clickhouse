FROM soulteary/lognormalizer:2.0.6 AS Builder
RUN apk add go
COPY datepatcher.go /
RUN go build datepatcher.go

FROM soulteary/lognormalizer:2.0.6
COPY --from=Builder /datepatcher /bin/