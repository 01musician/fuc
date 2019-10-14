# Expect

## 起源
Expect is an extension to the TCL scripting language written by Don Libes. The program automates interactions with programs that expose a text terminal interface. Expect, originally written in 1990 for the Unix platform, has since become available for Microsoft Windows and other systems.

## 功能
Expect is used to automate control of interactive applications such as Telnet, FTP, passwd, fsck, rlogin, tip, SSH, and others. Expect uses pseudo terminals (Unix) or emulates a console (Windows), starts the target program, and then communicates with it, just as a human would, via the terminal or console interface. Tk, another Tcl extension, can be used to provide a GUI.

## 用法
Expect serves as a "glue" to link existing utilities together. The general idea is to figure out how to make Expect use the system's existing tools rather than figure out how to solve a problem inside of Expect.

A key usage of Expect involves commercial software products. Many of these products provide some type of command-line interface, but these usually lack the power needed to write scripts. They were built to service the users administering the product, but the company often does not spend the resources to fully implement a robust scripting language. An Expect script can spawn a shell, look up environmental variables, perform some Unix commands to retrieve more information, and then enter into the product's command-line interface armed with the necessary information to achieve the user's goal. After retrieving information by interacting with the product via its command-line interface, the script can make intelligent decisions about what action to take, if any.

Every time an Expect operation is completed, the results are stored in a local variable called $expect_out. This allows the script to harvest information to feedback to the user, and it also allows conditional behavior of what to send next based on the circumstances.

## 使用事例
* 自动化telnet
```bash
# Assume $remote_server, $my_user_id, $my_password, and $my_command were read in earlier
# in the script.
# Open a telnet session to a remote server, and wait for a username prompt.
spawn telnet $remote_server
expect "username:"
# Send the username, and then wait for a password prompt.
send "$my_user_id\r"
expect "password:"
# Send the password, and then wait for a shell prompt.
send "$my_password\r"
expect "%"
# Send the prebuilt command, and then wait for another shell prompt.
send "$my_command\r"
expect "%"
# Capture the results of the command into a variable. This can be displayed, or written to disk.
set results $expect_out(buffer)
# Exit the telnet session, and wait for a special end-of-file character.
send "exit\r"
expect eof
```
* 自动化ftp
```bash
# Set timeout parameter to a proper value.
# For example, the file size is indeed big and the network speed is really one problem,
# you'd better set this parameter a value.
set timeout -1
# Open an ftp session to a remote server, and wait for a username prompt.
spawn ftp $remote_server
expect "username:"
# Send the username, and then wait for a password prompt.
send "$my_user_id\r"
expect "password:"
# Send the password, and then wait for an ftp prompt.
send "$my_password\r"
expect "ftp>"
# Switch to binary mode, and then wait for an ftp prompt.
send "bin\r"
expect "ftp>"
# Turn off prompting.
send "prompt\r"
expect "ftp>"
# Get all the files
send "mget *\r"
expect "ftp>"
# Exit the ftp session, and wait for a special end-of-file character.
send "bye\r"
expect eof
```

* 自动化sftp
```bash
#!/usr/bin/env expect -f

# procedure to attempt connecting; result 0 if OK, 1 otherwise
proc connect {passw} {
  expect {
    "Password:" {
      send "$passw\r"
        expect {
          "sftp*" {
            return 0
          }
        }
    }
  }
  # timed out
  return 1
}

#read the input parameters
set user [lindex $argv 0]
set passw [lindex $argv 1]
set host [lindex $argv 2]
set location [lindex $argv 3]
set file1 [lindex $argv 4]
set file2 [lindex $argv 5]

#puts "Argument data:\n";
#puts "user: $user";
#puts "passw: $passw";
#puts "host: $host";
#puts "location: $location";
#puts "file1: $file1";
#puts "file2: $file2";

#check if all were provided
if { $user == "" || $passw == "" || $host == "" || $location == "" || $file1 == "" || $file2 == "" }  {
  puts "Usage: <user> <passw> <host> <location> <file1 to send> <file2 to send>\n"
  exit 1
}

#sftp to specified host and send the files
spawn sftp $user@$host

set rez [connect $passw]
if { $rez == 0 } {
  send "cd $location\r"
  set timeout -1
  send "put $file2\r"
  send "put $file1\r"
  send "ls -l\r"
  send "quit\r"
  expect eof
  exit 0
}
puts "\nError connecting to server: $host, user: $user and password: $passw!\n"
exit 1
```

* 自动化ssh登陆
```bash
#timeout is a predefined variable in expect which by default is set to 10 sec
#spawn_id is another default variable in expect.
#It is good practice to close spawn_id handle created by spawn command
set timeout 60
spawn ssh $user@machine
while {1} {
  expect {

    eof                          {break}
    "The authenticity of host"   {send "yes\r"}
    "password:"                  {send "$password\r"}
    "*\]"                        {send "exit\r"}
  }
}
wait
close $spawn_id
```

