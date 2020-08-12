module github.com/kralamoure/d1proto

go 1.14

require (
	github.com/alexedwards/argon2id v0.0.0-20200802152012-2464efd3196b // indirect
	github.com/kralamoure/d1 v0.0.0-20200811215200-3ff36fd33625
	github.com/kralamoure/dofus v0.0.0-20200812040015-d1ce9c4da9ab
	golang.org/x/crypto v0.0.0-20200728195943-123391ffb6de // indirect
	golang.org/x/sys v0.0.0-20200810151505-1b9f1253b3ed // indirect
)

replace github.com/kralamoure/dofus => ../dofus

replace github.com/kralamoure/d1 => ../d1
