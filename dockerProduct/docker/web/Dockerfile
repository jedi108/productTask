# Telling to use Docker's golang ready image
FROM golang

MAINTAINER Vadim Tsurkov <vadimka108@gmail.com>

EXPOSE 8080 80 443

# Set pwd to the go folder
WORKDIR ${GOPATH}

RUN apt-get update -y && \
	apt-get install -y \
	vim mysql-client

# Install and build 
RUN go get -d github.com/jedi108/skylib/...
RUN go get github.com/jedi108/productTask
RUN go get github.com/pressly/goose

# Copy the entrypoint file
ADD /docker/web/infinite_loop.sh /infinite_loop.sh
RUN chmod +x /infinite_loop.sh
RUN cp /go/src/github.com/jedi108/productTask/config.json /go/bin/


# Start process
ENTRYPOINT ["/infinite_loop.sh"]