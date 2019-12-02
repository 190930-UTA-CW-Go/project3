FROM golang:latest
WORKDIR /app
# Now just add the binary
ADD clientbin /app/
ENTRYPOINT ["./clientbin"]
