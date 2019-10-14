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
* 自动化ftp
```console
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

 
