FROM google/golang:1.3

# Include GOPATH/bin in PATH.
ENV PATH $PATH:$GOPATH/bin

# Install Go package manager.
RUN go get github.com/mattn/gom 

# Copy the Gomfile into the image.
COPY Gomfile gopath/src/github.com/folieadrien/grounds/Gomfile

# Gom install inside app's location
RUN cd gopath/src/github.com/folieadrien/grounds && gom install

# Everything up to here was cached. This includes
# the gom install, unless the Gomfile changed.

# Now copy the app into the image.
COPY . gopath/src/github.com/folieadrien/grounds

# Set the final working dir to the app's location.
WORKDIR gopath/src/github.com/folieadrien/grounds

