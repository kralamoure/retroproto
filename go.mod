module github.com/kralamoure/d1proto

go 1.16

require (
	github.com/kralamoure/d1 v0.0.0-20201007013438-9eb50ef5ce7d
	github.com/kralamoure/dofus v0.0.0-20200927021741-893c10151570
	golang.org/x/crypto v0.0.0-20201002170205-7f63de1d35b0 // indirect
	golang.org/x/sys v0.0.0-20201008064518-c1f3e3309c71 // indirect
)

replace github.com/kralamoure/dofus => ../dofus

replace github.com/kralamoure/d1 => ../d1
