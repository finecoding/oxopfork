## Design of This Go Development Project

The code is hosted in a private repository at 

https://github.com/GoAtelier/oxo.git

There are two contributors who have forked copies of this repository

https://github.com/finecoding/oxopfork

and

https://github.com/VMaston/oxo-v-dev

The repository layout is simple and will become more involved as the code base develops.

Commands are located in subdirectories of /cmd and they will be kept small.  The entry point for each command can be found in the main.go file of each subdirectory.  These commands will be compiled into binaries for a command line interface style app and a web server.  There may be other apps for a testing and small utilities.

The commands will import code from packages.  It is not anticipated that these packages will be useful for third parties to import.  Therefore there will be no publicly accessible packages, they will be located in /internal and only accessible by from code within the module.

To begin with there will be two packages:  /internal/oxocli for a command line utility and /internal/oxoweb for the webserver code.  A third package may be used for code common to both when this becomes apparent.

The packages will be managed as a single module and the root of the module will also be the root of the repository and that is where the go.mod file will be found.

The repository layout will initially look like this:

go.mod
/cmd/cli 
/cmd/web
/internal/oxocli
/internal/oxoweb

Much of the code will be developed in files found in directories holding the oxocli and oxoweb packages.







