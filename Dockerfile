FROM golang
WORKDIR /opt
COPY . .
RUN go build -o showip
CMD /opt/showip