FROM golang:latest
WORKDIR /app
# Now just add the binary
ADD rego.pem /root/go/src/github.com/190930-UTA-CW-Go/project3/
ADD config ~/.ssh/
ADD . /app/
EXPOSE 9000
ENTRYPOINT ["./clientbin"]
