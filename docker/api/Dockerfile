FROM golang:1.23.2 as builder

ENV LANG=C.UTF-8

WORKDIR /app

# Installing packages
RUN apt-get -y install bash make

COPY . .

RUN make clean-cache
RUN make install
RUN make build

CMD ["/go-api/build/main"]

WORKDIR /app


# Copy and add permissions to entrypoint script
RUN cp ./docker/api/entrypoint.sh ./
RUN chmod +x ./entrypoint.sh

EXPOSE 9000

ENTRYPOINT ["/app/entrypoint.sh"]
