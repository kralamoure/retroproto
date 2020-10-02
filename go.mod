module github.com/kralamoure/d1proto

go 1.15

require (
	github.com/alexedwards/argon2id v0.0.0-20200802152012-2464efd3196b // indirect
	github.com/kralamoure/d1 v0.0.0-20200917030335-f23076eacc5c
	github.com/kralamoure/dofus v0.0.0-20200917024449-5e4b76236af8
	golang.org/x/crypto v0.0.0-20200820211705-5c72a883971a // indirect
)

replace github.com/kralamoure/dofus => ../dofus

replace github.com/kralamoure/d1 => ../d1
