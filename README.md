# cfg2env
Simple tool to set enviroment variable from config file.

## Usage
Specify your config file as an argument.  
Example: `cfg2env config.ini`

## Format of the config file
cfg2env supports this format except indent. See below.  
https://docs.python.org/3/library/configparser.html#supported-ini-file-structure  
**Note:** Sections (enclosed by `[]`) will be ignored.

## Install
I haven't prepared the way to install.  
Please build `main.go` myself and move the binary file to `/usr/bin` or `~/bin`. 

## Requirement
Go 1.14 or above (Probably It works on older Go, but unwarranted)

## License
MIT License.
