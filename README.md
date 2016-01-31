# magic-filetransfer
Transfer files like a magican!

magic-filetransfer is for people who want to transfer files without the hassle
of configuration or IP-addresses. It magically detects the other side (which
wants to recieve the file or which wants to transmit the file)

# How do i use this?
Clone it and change into the directory:
```sh
git clone https://github.com/jannickfahlbusch/magic-filetransfer.git
cd magic-filetransfer
```

Compile and execute it:
```sh
go build -o mft src/*.go
./mft
```
OR
Run it instantly:
```sh
go run src/*
```

# License
This proram is licensed under the GPLv3
