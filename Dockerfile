FROM alpine:latest
COPY ./bin/mux ./mux
EXPOSE 3211
CMD ["./mux"]