#bin/bash

# go build command
# classic
# go build -o ansible/build/website main.go
# raspberry
env GOOS=linux GOARCH=arm GOARM=5 go build -o build/website ../main.go


ansible webserver -m ping
ansible-playbook -i hosts deploy-app.yml