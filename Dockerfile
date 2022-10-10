FROM golang:1.19
WORKDIR $GOPATH/src/github.com/avummet/stock-tracker


COPY .  .

# Download all the dependencies
RUN go get -d -v ./...

# Install the package
RUN go install -v ./...

EXPOSE 7443

#USER "$UID"

CMD ["stock-tracker"]
