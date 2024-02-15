# syntax=docker/dockerfile:1
# Alpine is chosen for its small footprint
# The base go-image
#FROM golang:1.18-alpine
 
# Create a directory for the app
#RUN mkdir /app
 
# Copy all files from the current directory to the app directory
#COPY . /app
 
# Set working directory
#WORKDIR /app
 
# Run command as described:
# go build will build an executable file named azurepoc in the current directory
#RUN go build -o azurepoc . 
 
# Run the server executable
#CMD [ "/app/azurepoc" ]

# syntax=docker/dockerfile:1
# Alpine is chosen for its small footprint
# compared to Ubuntu
FROM golang:1.18-alpine
## We create an /app directory within our
## image that will hold our application source
## files
RUN mkdir /app
## We copy everything in the root directory
## into our /app directory
ADD . /app
# Set working directory
WORKDIR /app
#Copy go mod and sum files
COPY go.mod ./
COPY go.sum ./
#Run go mod under app directory
RUN go mod download
#Copy all go files
COPY *.go ./

# Run command as described:
# go build will build an executable file named azurepoc in the current directory
RUN go build -o /azurepoc
# Run the server executable
CMD [ "/azurepoc" ]