# Dockerfile.builder
FROM gcc:4.9
COPY main.c /main.c
CMD ["gcc", "-std=c99", "-static", "-o", "/out/hello", "/main.c"]
