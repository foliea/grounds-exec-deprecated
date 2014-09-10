FROM google/golang:1.3

# Include GOPATH/bin in PATH.
ENV PATH $PATH:$GOPATH/bin

# Install Go package manager.
RUN go get github.com/mattn/gom 

# Copy the Gomfile into the image.
COPY Gomfile /grounds/Gomfile

# Link app's location to gopath
RUN mkdir -p gopath/src/github.com/folieadrien
RUN ln -s /grounds gopath/src/github.com/folieadrien/grounds

# Gom install inside app's location
RUN cd /grounds && gom install

# Everything up to here was cached. This includes
# the gom install, unless the Gomfile changed.

# Now copy the app into the image.
COPY . /grounds

# Set the final working dir to the app's location.
WORKDIR /grounds
