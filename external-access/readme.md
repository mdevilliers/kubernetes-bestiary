

 ---------
|         |
|    P    | <---- install an HAProxy / Kubernentes API listener
|         |
 ---------

 ---------
|         |
|  K8 M   |
|         |
 ---------

 ---------
|         |
| K8 Node |
|         |
 ---------

Logic
-----

Watch All Services in K8

	Iterate

		- look for labels e.g.
			- external:public=true
			- external:type=TCP | HTTP
			- external:uri= /api/myservice1

Write to HAProxy
HAProxy reload