FROM golang:1.21-alpine

WORKDIR /opt/code/
ADD ./ /opt/code/

# apk - алипина
RUN apk update && apk upgrade && \
    apk add --no-cache git

# скачивание зависимостей из go.mod файла
RUN go mod download

# сборка бинарника в папку bin/goserver путь контейнера и путь к файлу, который нужно компилировать
RUN go build -o bin/goserver cmd/app/main.go

# инструкция куда должен обратиться контейнер при запуске
ENTRYPOINT [ "bin/goserver" ]