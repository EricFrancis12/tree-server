build:
	templ generate
	go build .

# How to pass in arguments:
#
# make run ARGS="-PORT=4000 -WD=C:\Users\Jim\Desktop\my\folder"
run: build
	./tree-server $(ARGS)
