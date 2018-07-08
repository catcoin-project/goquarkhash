CFLAGS=-Wall -O3
SRCS=aes_helper.c blake.c bmw.c groestl.c jh.c keccak.c skein.c
OBJS=$(SRCS:%.c=%.o)
CGOFILES=quarkhash.go

%.o : %.c
	$(CC) $(CFLAGS) -c $< -o $@

build: libsha3.a

all: clean build test
	@gofmt -e -s -w .
	@go build

test:
	@go test

clean:
	rm -f $(OBJS) libsha3.a goquarkhash

libsha3.a: $(OBJS)
	$(AR) -rcs $@ $(OBJS)

