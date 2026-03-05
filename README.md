# WireWeave v6

---

## Goals

The point of this project is to create something similar to the millions of other VPN management solutions 
out there (e.g., NetBird, wg-easy, headscale, pritunl, firezone, etc), but something that is designed
for IPv6-only or dual stack environments. While most of these applications have IPv6 support, its almost always
an afterthought and makes IPv6-only networks almost impossible to manage. 

For many administrators, IPv6 is one of those things that doesn't really make sense (as it more or less eliminates
all of the problems that came with the overutilization of IPv4), and in some respects, may appear to be too good
to be true. Especially for home labbers, being able to have more than one IP address with a normal residential
internet plan may seem liberating but confusing. However, the fact of the matter is that IPv6 allows you to eliminate
all of your private DNS servers, NAT, and local subnets. 

This project aims to create a VPN server/client that's designed around these ideas. We don't need to host a private DNS
server, we don't need to NAT traffic (even if we may choose to do so in some cases), and we can get rid of the local
subnet and focus strictly on routing via public IPv6 addresses. 
