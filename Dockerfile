FROM golang:1.22.0 as builder

# RUN apk add --no-cache tzdata
ENV TZ=Asia/Jakarta

WORKDIR /app

COPY go.mod go.sum ./
RUN go install github.com/air-verse/air@latest
COPY . .

RUN go mod download
# RUN curl -fLo install.sh https://raw.githubusercontent.com/cosmtrek/air/master/install.sh \
#     && chmod +x install.sh && sh install.sh && cp ./bin/air /bin/air
# To prevent error message buildvcs, dont remove this
# ENV GOFLAGS="-buildvcs=false"
# RUN go install github.com/cosmtrek/air@latest
# RUN go build -o ./tmp/main .

CMD ["air", "-c", ".air.toml"]
# CMD ["air"]