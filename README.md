# raven

main source of truth for a cyber competition red team

https://youtu.be/p0mZmnl792o?si=IbfBX9GMjRMoLJLl

## testing workflow

1) go to http://localhost:8080/
2) login using the password from the console or db
3) go to http://localhost:8080/teams
4) add a team and call it `brice`
5) add a network to team `brice` with cidr `10.100.10.0/24`
6) go to http://localhost:8080/uploads
7) upload sample `example.xml` file as nmap xml
8) go to http://localhost:8080/teams/1

## build instructions

clone or download zip:
```
git clone https://github.com/nos3curity/raven
```

install the bee cli tool:
```
go install github.com/beego/bee/v2@latest
```

cd into folder:
```
cd raven/
```

init the module and download libraries:
```
go mod init raven
go mod tidy
```

run the app:
```
bee run
```

## Docker

### Development Environment

1. Clone

2. Build Docker Image

```
docker build -t raven-dev .
```

3. Init the modules and download libraries
```
go mod init raven
go mod tidy
go mod vendor
```

5. Run

```
docker run -it --rm -p 8080:8080 -v $PWD/src:/go/src/raven raven-dev
```

### Production Environment

1. Clone

2. Build Production Docker Image
```
docker build -t raven-prod -f Dockerfile.production
```

3. Run
```
docker run -it -p 8080:8080 raven-prod
```

## documentation

### primary functions

these functions exist to manipulate model objects like add, remove, get, etc.

the function names should be self-explanatory. see model code for usage:
- [x] networks
	- [x] AddNetwork
	- [x] DeleteNetwork
	- [x] GetNetwork
	- [x] GetAllNetworks
	- [x] GetNetworkSystems
- [x] system
	- [x] AddSystem
	- [x] GetSystem
	- [x] DeleteSystem
	- [x] ChangeOs
	- [x] ChangeHostname
	- [x] GetSystemPorts
- [x] team
	- [x] AddTeam
	- [x] DeleteTeam
	- [x] GetTeam
	- [x] GetAllTeams
	- [x] RenameTeam
	- [x] GetTeamNetworks
