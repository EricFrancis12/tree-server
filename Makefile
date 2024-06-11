build:
	go build .

run: build
	go run . $(ARGS)

# How to pass in arguments:
#
# make run ARGS="-PORT=4000 -WD=C:\Users\Jim\Desktop\my\folder"
