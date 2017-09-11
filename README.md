# Playground for Pipe / GRPC experiment.


Thoughts are create a fifo file /tmp/`hostname`.

1) Have Apache/Tomcat/Syslog output such file. 
2) client daemon reads from pipe, sends data to server for processing.  Client would only know to write to a file.
   a) Example "echo "helllo from web server" > /tmp/mylog-server
   
 
 Problems.
 
 1) Pipe blocks both read/write.  Write is an issue,  write cant happen unless read is available.
 2) Client goes down, halts write on apache log writes.
 3) logrotate issue with it trying to roll the log.
 
 
 
 
