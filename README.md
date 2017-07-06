collectd-nsd: NSD statistics with collectd
==================================================

This utility allows you to access [NSD][0] statistics from [collectd][1] using the [exec][2] plugin. It hasn't been thoroughly tested, use it at your own risk...

Building
--------

```
go get -u github.com/falzm/collectd-nsd
```

Configuration
-------------

Once the `collectd-nsd` binary is compiled, copy it wherever you want and add the following lines to you collectd configuration:

```
LoadPlugin exec
<Plugin "exec">
    Exec "nsd" "/usr/bin/collectd-nsd"
</Plugin>
```

Note: the utility executes the command `nsd-control stats_noreset` to fetch the statistics: make sure the user specified in your `exec` plugin block has the permissions to execute the command. To verify this (replace *nsd* with your user of choice):

```
$ sudo -u nsd nsd-control stats_noreset
server0.queries=62
num.queries=62
time.boot=47397.898856
time.elapsed=4272.464956
size.db.disk=50176
size.db.mem=18672
size.xfrd.mem=20990440
size.config.disk=0
size.config.mem=2160
num.type.A=33
num.type.NS=2
num.type.MD=0
num.type.MF=0
num.type.CNAME=0
num.type.SOA=3
num.type.MB=0
num.type.MG=0
num.type.MR=0
num.type.NULL=0
num.type.WKS=0
num.type.PTR=1
num.type.HINFO=0
num.type.MINFO=0
num.type.MX=0
num.type.TXT=0
num.type.RP=0
num.type.AFSDB=0
num.type.X25=0
num.type.ISDN=0
num.type.RT=0
num.type.NSAP=0
num.type.SIG=0
num.type.KEY=0
num.type.PX=0
num.type.AAAA=23
num.type.LOC=0
num.type.NXT=0
num.type.SRV=0
num.type.NAPTR=0
num.type.KX=0
num.type.CERT=0
num.type.DNAME=0
num.type.OPT=0
num.type.APL=0
num.type.DS=0
num.type.SSHFP=0
num.type.IPSECKEY=0
num.type.RRSIG=0
num.type.NSEC=0
num.type.DNSKEY=0
num.type.DHCID=0
num.type.NSEC3=0
num.type.NSEC3PARAM=0
num.type.TLSA=0
num.type.CDS=0
num.type.CDNSKEY=0
num.type.OPENPGPKEY=0
num.type.CSYNC=0
num.type.SPF=0
num.type.NID=0
num.type.L32=0
num.type.L64=0
num.type.LP=0
num.type.EUI48=0
num.type.EUI64=0
num.opcode.QUERY=62
num.class.IN=62
num.rcode.NOERROR=59
num.rcode.FORMERR=0
num.rcode.SERVFAIL=0
num.rcode.NXDOMAIN=2
num.rcode.NOTIMP=0
num.rcode.REFUSED=0
num.rcode.YXDOMAIN=0
num.edns=53
num.ednserr=0
num.udp=56
num.udp6=5
num.tcp=1
num.tcp6=0
num.answer_wo_aa=0
num.rxerr=0
num.txerr=0
num.raxfr=0
num.truncated=0
num.dropped=0
zone.master=2
zone.slave=0
```

See the _STATISTIC COUNTERS_ section from the [nsd-control][4] command manpage for information regarding metrics signification.

TODO
----

 * As of now the plugin only processes the metrics prefixed with `num.`. Feel free to hack this according to your needs.

License
-------

Copyright (c) 2017, Marc Falzon.
All rights reserved.

Redistribution and use in source and binary forms, with or without
modification, are permitted provided that the following conditions
are met:

 * Redistributions of source code must retain the above copyright
   notice, this list of conditions and the following disclaimer.

 * Redistributions in binary form must reproduce the above copyright
   notice, this list of conditions and the following disclaimer in the
   documentation and/or other materials provided with the distribution.

 * Neither the name of the authors nor the names of its contributors
   may be used to endorse or promote products derived from this software
   without specific prior written permission.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE
ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE
LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR
CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF
SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS
INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN
CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE)
ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE
POSSIBILITY OF SUCH DAMAGE.


[0]: https://nlnetlabs.nl/projects/nsd/
[1]: https://collectd.org/
[2]: https://collectd.org/documentation/manpages/collectd-exec.5.shtml
[3]: https://github.com/octo/go-collectd/
[4]: https://nlnetlabs.nl/projects/nsd/nsd-control.8.html
