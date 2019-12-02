FROM golang:latest
WORKDIR /app
# Now just add the binary
ADD clientbin /app/
RUN chmod 777 clientbin
ENTRYPOINT ["./clientbin"]
