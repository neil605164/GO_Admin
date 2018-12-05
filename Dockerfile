# 第一層基底
FROM golang:1.11.2-alpine AS build

# 複製原始碼
COPY . /go/src/GO_Admin
WORKDIR /go/src/GO_Admin

# 進行編譯(名稱為：go_admin)
RUN go build -o go_admin

# 最終運行golang 的基底
FROM alpine

COPY --from=build /go/src/GO_Admin/conf /app/conf
COPY --from=build /go/src/GO_Admin/go_admin /app/go_admin
WORKDIR /app

ENTRYPOINT [ "./go_admin" ]