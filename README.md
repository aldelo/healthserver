# healthserver
 Stand alone health server for linux or windows, reporting http status 200 when url path /health is called. Works with load balancer health check such as AWS ELB target group health checks.
 
# deploy
server.yaml file must be deployed together with the binary in the same root folder.
yaml file name should be server.yaml, unless name is changed within init code.

# reconfigure
web server changes can be done via edit of server.yaml settings

# install service on linux or windows
for linux, using systemd service configuration, and sudo systemctl start xyz.service method (details see /common project's /systemd wrapper comments)
for windows, using sc.exec command, for example "c:\sc.exe create xyzService binpath= c:\xyzFolder\xyz.exe type= own start= auto"

