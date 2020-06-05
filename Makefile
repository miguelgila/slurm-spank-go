#gcc -pthread bar.c foo.a -o foo
#go build -buildmode=c-archive foo.go

#override CFLAGS += -O2 -Wall -Wvla -Werror -Wfatal-errors -pthread -shared -fPIC 
override CFLAGS += -pthread 
CC = /usr/bin/gcc

# SLURM
HAS_SQUEUE := $(shell (which squeue 2> /dev/null))
ifneq ($(strip $(HAS_SQUEUE)),)
       HAS_SLURM = $(shell dirname $(HAS_SQUEUE))
       SLURM_ROOT = $(shell dirname $(HAS_SLURM))
       override CFLAGS += -I $(SLURM_ROOT)/include
       override LDFLAGS += -L $(SLURM_ROOT)/lib64 -lslurm
else
	$(error Unable to build without SLURM support)
endif


all: foo bar

foo: foo.go
	go build -buildmode=c-shared -o foo.o foo.go

bar: foo
	@$(CC) $(CFLAGS) $(LDFLAGS) -Wall -std=gnu99 -shared -fPIC -o bar.o -c bar.c
	@$(CC) $(CFLAGS) $(LDFLAGS) -shared -o bar.so foo.o bar.o

.PHONY: install
install: 
	install bar.so /usr/lib64
	install foo.o /usr/lib64

.PHONY: uninstall
uninstall:
	rm -f /usr/lib64/bar.so
	rm -f /usr/lib64/foo.o

clean:
	rm -f foo.a foo.h bar core foo foo.so foo.h bar.so
