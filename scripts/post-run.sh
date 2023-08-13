apt update
apt upgrade -y
git config --global user.name "Serhioromano"
git config --global user.email "serg4172@mail.ru"
cd /workspaces/gotest/app
go get github.com/githubnemo/CompileDaemon
go install github.com/githubnemo/CompileDaemon
cd /workspaces/gotest/app/cmd/api/
CompileDaemon -command="./api"
