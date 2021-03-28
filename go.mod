module github.com/kralamoure/d1proto

go 1.16

require (
	github.com/alexedwards/argon2id v0.0.0-20210326052512-e2135f7c9c77 // indirect
	github.com/kralamoure/d1 v0.0.0-20210325215504-184ee80d8398
	github.com/kralamoure/dofus v0.0.0-20200927021741-893c10151570
	golang.org/x/crypto v0.0.0-20210322153248-0c34fe9e7dc2 // indirect
	golang.org/x/sys v0.0.0-20210326220804-49726bf1d181 // indirect
)

replace github.com/kralamoure/dofus => ../dofus

replace github.com/kralamoure/d1 => ../d1
