FROM golang:1.22.3

ENV LANG=C.UTF-8

WORKDIR /app

# Installing packages
RUN apt-get -y install bash make

# Copy project files
COPY . .

# Install dependencies
RUN make go-install
# Generate swagger docs
RUN make go-swagger
# Build
RUN make go-build

# Run
CMD ["/go-api/build/main"]


WORKDIR /app


# Copy and add permissions to entrypoint script
RUN cp ./docker/api/entrypoint.sh ./
RUN chmod +x ./entrypoint.sh

EXPOSE 9000

ENTRYPOINT ["/app/entrypoint.sh"]
