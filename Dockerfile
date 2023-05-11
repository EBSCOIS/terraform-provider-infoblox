from golang:alpine as golang

RUN apk update

RUN apk add terraform make

WORKDIR /terraform-provider-infoblox

COPY ./ ./

RUN go mod vendor

RUN go test -i $(go list ./... |grep -v 'vendor') || exit 1 \
    echo $(go list ./... |grep -v 'vendor') | \
    xargs -t -n4 go test  -timeout=30s -parallel=4

RUN CGO_ENABLED=0 go build -o build-output/terraform-provider-infoblox -mod=vendor

# RUN TF_LOG=info terraform plan

CMD ["/bin/sh"]

from scratch as export-stage
copy --from=golang /terraform-provider-infoblox/build-output .