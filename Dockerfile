FROM golang:latest
WORKDIR /app
# Now just add the binary
ADD rego.pem /root/go/src/github.com/190930-UTA-CW-Go/project3/
ADD ssh_config /etc/ssh/
ADD . /app/
RUN chmod 400 rego.pem
EXPOSE 9000
ENTRYPOINT ["./clientbin"]
