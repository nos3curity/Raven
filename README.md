# AC-130

main source of truth for a cyber competition red team

https://youtu.be/p0mZmnl792o?si=IbfBX9GMjRMoLJLl

## testing workflow

1) go to http://localhost:8080/teams
2) add a team and call it `brice`
3) add a network to team `brice` with cidr `10.100.10.0/24`
4) go to http://localhost:8080/uploads
5) upload sample `example.xml` file as nmap xml
6) go to http://localhost:8080/teams/1

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

## documentation

### primary functions

these functions exist to manipulate model objects like add, remove, get, etc.

the function names should be self-explanatory. see controller code for usage:
- [x] networks
	- [x] AddNetwork
	- [x] DeleteNetwork
	- [x] GetNetwork
	- [x] GetAllNetworks
	- [x] GetNetworkSystems
- [ ] system
	- [x] AddSystem
	- [x] GetSystem
	- [x] DeleteSystem
	- [ ] ChangeOs
	- [ ] ChangeHostname
	- [x] GetSystemPorts
- [x] team
	- [x] AddTeam
	- [x] DeleteTeam
	- [x] GetTeam
	- [x] GetAllTeams
	- [x] RenameTeam
	- [x] GetTeamNetworks