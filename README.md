racktables-fingerd
==================

This is a [finger](http://tools.ietf.org/html/rfc742) daemon that interfaces with [RackTables](http://racktables.org/) databases. It allows to quickly figure out on which host(s) a virtual machine with the given name runs.

The following command will figure out on which host servers the virtual machine named `web01.vm` is running according to RackTables:

    $ finger web01.vm@inventory.site
    web01.vm: host1.site
    web01.vm: host2.site
    $

The daemon searches the common name of each object so one can also search for all virtual machines that have `web` in their common name:

    $ finger web@inventory.site
    web01.vm: host1.site
    web01.vm: host2.site
    web02.vm: host3.site
    web03.vm: host3.site
    [...]
    $

Dependencies
------------

Besides to Go compiler you need the [mymysql library](https://github.com/ziutek/mymysql) to compile this program:

    $ go get github.com/ziutek/mymysql/autorc
    $ go get github.com/ziutek/mymysql/godrv

Installation
------------

Installation is fairly simple:

* Clone the repository
* Run `go build` inside the source directory
* Execute `racktables-fingerd`

Database setup
--------------

The code assumes that the MySQL user `finger` can access the RackTables
database on `localhost` using the password `finger`.

`grant select on racktables.* to finger@localhost identified by 'finger'`

YMMV. Adjust the constants at the top of `fingerd.go` if necessary.
