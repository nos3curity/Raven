# AC-130

main source of truth for a cyber competition red team

https://youtu.be/p0mZmnl792o?si=IbfBX9GMjRMoLJLl

## documentation

### testing workflow

1) add a team with `/teams/add?team_name=brice`
2) add a network with the form on `/networks` or `/networks/add?network_cidr=10.100.10.0/24`
3) parse the sample nmap scan by going to `/nmap` or upload your own at `/uploads/nmap`
4) marvel at the data returned by `/networks`

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
	- [x] RenameTeam
	- [x] GetTeamNetworks

### build instructions

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