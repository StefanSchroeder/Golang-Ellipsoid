GOFMT=gofmt -s -spaces=true -tabindent=false -tabwidth=4

GOFILES=\
  hello_world.go\

TARGET=hello_world

ifeq ($(GOOS),windows)
  TARGET=hello_world.exe
endif
  
all:
	8g ellipsoid/ellipsoid.go
	8g hello_world.go && 8l -o $(TARGET) hello_world.8 && ./$(TARGET) 
	@echo Expected results:
	@echo "  Distance = 543044.190419953 Bearing = 137.50134015496275"
	@echo "  lat3 = 37.74631054036373 lon3 = -122.21438161492877"
	@echo "  x = 1.1042590709397183e+06"
	@echo "  y = -4.824765955871677e+06"
	@echo "  z = 4.0093940281868847e+06"
	@echo "  lat5 = 39.197807000000005 lon5 = -77.10857400000002 alt3 = 55.00000000093132"

format:
	${GOFMT} -w ${GOFILES}

clean:
	rm -f hello_world.8 hello_world.exe hello_world *~ ellipsoid.8
	rm -f ellipsoid/*~ ellipsoid/*.8
