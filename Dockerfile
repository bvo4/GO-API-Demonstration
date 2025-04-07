# Specifies a parent image
FROM golang:latest

# Creates an app directory to hold your app's source code
WORKDIR /app

# Tracks changes within go.mod file
COPY go.mod .

# Copies everything from root directory into /app
COPY . .

COPY appsetting.json ./app/appsetting.json

# Installs Go Dependencies
RUN go mod download 

# Builds app
RUN go build -o /main

#Opens port
EXPOSE 8090

ENTRYPOINT ["/main"]