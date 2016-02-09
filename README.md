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


This are the commandline flags you can use:
```
usage: mft [<flags>]

Flags:
      --help                   Show context-sensitive help (also try --help-long and --help-man).
  -f, --fileName=FILENAME      Name of the file to transmit
  -o, --outputDir=OUTPUTDIR    Directory to save the retrieved file in
      --outputName=OUTPUTNAME  Save the recieved file under a different name
  -h, --usage                  Print the usage for magic-filetransfer and exit
  -v, --version                Print the version for magic-filetransfer
```

# License
This program is licensed under the GPLv3
