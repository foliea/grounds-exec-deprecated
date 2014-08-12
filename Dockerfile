FROM google/golang:1.3

# Include GOPATH/bin in PATH
ENV PATH $PATH:$GOPATH/bin

# Install Go package manager
RUN go get github.com/mattn/gom 

COPY Gomfile gopath/src/github.com/folieadrien/grounds/Gomfile

RUN cd gopath/src/github.com/folieadrien/grounds && gom install

COPY . gopath/src/github.com/folieadrien/grounds

WORKDIR gopath/src/github.com/folieadrien/grounds

