FROM golang:1.22
COPY . /app
WORKDIR /app
RUN go install github.com/wailsapp/wails/v2/cmd/wails@latest
RUN wails build -platform windows/amd64,darwin/arm64