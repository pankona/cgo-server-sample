
UNAME := $(shell uname)

ifeq ($(UNAME), Darwin)
BUILD_PARAM := -framework CoreFoundation -framework Security
endif

all:
	make -C ../
	gcc -o main main.c ../out/libsample.a -lpthread $(BUILD_PARAM)

