# Bash-Parser
А program that accepts Bash commands from a user and executes them.  


Move into the directory  `BashParser`.To test the program, use the `go run` command:  
`$ go run main.go -cmd "cat commands/test_files/abc.txt"`  

Create a Go module:  
`$ go mod init BashParser`  

Generate an executable binary for the application:  
`$ go build`  

The executable was added to the current directory. Run the `ls` command to check.   
Output:  
`$ go.mod BashParser main.go`  

Run it to make sure the binary has been built correctly. Execute the following command:  
`$ ./BashParser`  

Run the command from your `$HOME/bin` directory,where the executable file `BashParser` may be located.  
The executable files of the individual commands must be in the same directory. This will yield the following  
`$ BashParser`  

This will yield the following:  
`$ BashParser`   
 `Usage: ProgramName -cmd`    
      `-cmd string`  
    	`Please, specify command.`  
      
Test files are located in the `/tmp/test-files` directory.  
The following commands can be executed:   
`$ BashParser -cmd "cd /"`  
`$ BashParser -cmd "cat /tmp/test_files/<filename>.txt"`  
`$ BashParser -cmd "grep <pattern> /tmp/test_files/*.*"`   
`$ BashParser -cmd "rm /tmp/test_files/test/rm_file.txt"`  
`$ BashParser -cmd "ls /tmp/test_files"`  
`$ BashParser -cmd "sort /tmp/test_files/<filename>.txt"`  
`$ BashParser -cmd "tail /tmp/test_files/<filename>.txt"`  
`$ BashParser -cmd "mv /tmp/test_files/<filename>.txt /tmp/test_files/test"`  
