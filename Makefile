CFLAGS=-Wall -O3
SRCS=sha3/aes_helper.c sha3/blake.c sha3/bmw.c sha3/groestl.c sha3/jh.c sha3/keccak.c sha3/skein.c
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

