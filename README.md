# TODO Cli App  
Your classic todo app but from the command line and written in Go. 

Available command line tasks:
- add [task]
- delete [id]
- complete [id]
- list
- help




## Things Learned from this Project
This was my first time using the Go encoding/csv package in a meaningful way so most of the interactions with a csv file were completely new to me.  I'm sure there is a better method for updating or delete rows but I wanted to try and figure out as much as I could without copy/pasting from something like stack overflow.

I also did not use a cli package to help format the table of the output with the list command so attempting to have some minor formatting on the output was difficult.

## Future Enhancements
- Fix formatting on list command
    - possibly use a package
- Setup build binary to install to terminal so the project can be called with something like:
    - task add [task]
    - task delete 1
    - task list