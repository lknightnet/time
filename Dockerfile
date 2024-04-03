FROM ubuntu

WORKDIR /app

COPY server /app

EXPOSE 8000

CMD ["./server"]