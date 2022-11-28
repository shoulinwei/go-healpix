GOPATH:=$(shell go env GOPATH)
FLAGS = -Wwrite-strings

all: test_healpix

OBJS = healpix-utils.o healpix.o starutil.o permutedsort.o mathutil.o bl.o qsort_reentrant.o
SRCS = healpix-utils.c healpix.c starutil.c permutedsort.c mathutil.c bl.c qsort_reentrant.c

HEADERS = healpix-utils.h healpix.h

$(OBJS): %.o: %.c $(HEADERS)
	gcc -fPIC -o $@ -c $<

%.o: %.c
	gcc -fPIC -o $@ -c $<

so: $(OBJS)
	gcc ${OBJS} ${FLAGS} -fPIC -lm -shared -o libhpix.so
	cp libhpix.so ../libhpix.so

dylib: $(OBJS)
	gcc ${OBJS} ${FLAGS} -fPIC -shared -o libhpix.dylib
	cp libhpix.dylib ../libhpix.dylib

test_healpix: test_healpix-main.c test_healpix.c $(OBJS) CuTest.o

example: example.c $(OBJS)

clean:
	rm -rf *.o *.so *.dylib

