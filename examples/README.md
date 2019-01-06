<!--
Copyright 2017 - The TXTDirect Authors

This work is licensed under a Creative Commons Attribution-ShareAlike 4.0 International License;
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    https://creativecommons.org/licenses/by-sa/4.0/legalcode
Unless required by applicable law or agreed to in writing, documentation
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
-->

# Configuration
**Redirect all requests to "www"-subdomain:**  
*example.com -> www.example.com*
```
txtdirect {
  enable www
}
```

**Redirect to host provided in TXT record:**  
*Default: Return 404 on empty record*
```
txtdirect {
  enable host
}
```

**Redirect to host provided in TXT record:**  
*Default: Redirect to "www"-subdomain on empty record*
```
txtdirect {
  enable host www
}
```

**Redirect to host provided in TXT record:**  
*Default: Redirect to example.com on empty record*
```
txtdirect {
  redirect https://example.com
}
```

**Enable everything except "www"-subdomain redirection:**  
```
txtdirect {
    disable www
}
```

**Enable path based redirects:**  
```
txtdirect {
    enable path
}
```

**Enable go meta/vanity redirects:**  
*pkg.example.com -> github.com/some/pkg.git*
```
txtdirect {
  enable gometa
}
```
<!--
# Placeholders
{dir} 	        The directory of the requested file (from request URI)  
{file} 	        The name of the requested file (from request URI)  
{>Header} 	    Any request header, where "Header" is the header field name  
{host} 	        The host value on the request  
{hostname} 	    The name of the host machine that is processing the request  
{hostonly} 	    Same as {host} but without port information  
{method} 	      The request method (GET, POST, etc.)  
{path} 	        The path portion of the original request URI (does not include query string or fragment)  
{path_escaped} 	Query-escaped variant of {path}  
{port} 	        The client's port  
{query} 	      The query string portion of the URL, without leading "?"  
{query_escaped} The query-escaped variant of {query}  
{?key} 	        The value of the "key" argument from the query string  
{remote} 	      The client's IP address  
{scheme} 	      The protocol/scheme used (usually http or https)  
{uri} 	        The request URI (includes path and query string)  
{uri_escaped} 	The query-escaped variant of {uri}  
-->

# TXT records
"txtdirect.example.com" is your hosted TXTDIRECT instance and is usually provided as CNAME.

For more details take a look at our [TXT-record specification](/docs/README.md#specification).

## host
**Host based redirect using default redirection code**
*example.com -> about.example.com 301*
```
example.com                   3600 IN A      127.0.0.1
_redirect.example.com         3600 IN TXT    "v=txtv0;to=https://about.example.com;type=host"
```

**Host based redirect using explicit 301 redirection code**
*www.example.com -> about.example.com 301*
```
www.example.com               3600 IN CNAME  txtdirect.example.com.
_redirect.www.example.com     3600 IN TXT    "v=txtv0;to=https://about.example.com;type=host;code=301"
```

**Host based redirect using explicit 302 redirection code**
*www.example.com -> about.example.com 302*
```
www.example.com               3600 IN CNAME  txtdirect.example.com.
_redirect.www.example.com     3600 IN TXT    "v=txtv0;to=https://about.example.com;type=host;code=302"
```

**Host based redirect including path**
*gophers.example.com -> example.com/gophers*
```
gophers.example.com           3600 IN CNAME  txtdirect.example.com.
_redirect.gophers.example.com 3600 IN TXT    "v=txtv0;to=https://example.com/gophers;type=host"
```

<!--
**Host based redirect including {url} placeholder**
*gophers.example.com/1 -> example.com/gophers/1*
```
gophers.example.com           3600 IN CNAME  txtdirect.example.com.
_redirect.gophers.example.com 3600 IN TXT    "v=txtv0;to=https://example.com/gophers{uri};type=host"
```

**Host based redirect including {file} placeholder**
*placeholder.example.com/cat.png -> example.com/cat.png*
```
placeholder.example.com              3600 IN CNAME  txtdirect.example.com.
_redirect.placeholder.example.com    3600 IN TXT    "v=txtv0;to=https://example.com/{file};type=host"
```
-->

## path
**Path based redirect using default ordering**
*example.com/firstMatch/secondMatch -> about.example.com*
*example.com/firstMatch/noMatch -> defaults or 404*
```
example.com                                   3600 IN A      127.0.0.1
_redirect.example.com                         3600 IN TXT    "v=txtv0;type=path"
_redirect.secondmatch.firstmatch.example.com  3600 IN TXT    "v=txtv0;to=https://about.example.com/;type=host"
```

**Path based redirect using default ordering**
*example.com/firstMatch -> about.example.com*
*example.com/noMatch -> defaults or 404*
```
example.com                                   3600 IN A      127.0.0.1
_redirect.example.com                         3600 IN TXT    "v=txtv0;type=path"
_redirect.firstmatch.example.com              3600 IN TXT    "v=txtv0;to=https://about.example.com/;type=host"
```

**Path based redirect using explicit ordering**
*example.com/firstMatch/secondMatch -> about.example.com*
*example.com/firstMatch/noMatch -> defaults or 404*
```
example.com                                   3600 IN A      127.0.0.1
_redirect.example.com                         3600 IN TXT    "v=txtv0;from=/$1/$2/;type=path"
_redirect.secondmatch.firstmatch.example.com  3600 IN TXT    "v=txtv0;to=https://about.example.com/;type=host"
```

**Path based redirect using modified ordering**
*example.com/firstMatch/secondMatch -> about.example.com*
*example.com/firstMatch/noMatch -> defaults or 404*
```
example.com                                   3600 IN A      127.0.0.1
_redirect.example.com                         3600 IN TXT    "v=txtv0;from=/$2/$1/;type=path"
_redirect.firstmatch.secondmatch.example.com  3600 IN TXT    "v=txtv0;to=https://about.example.com/;type=host"
```

**Path based redirect fallback on root/index**
*example.com/ -> root.example.com*
```
example.com                                   3600 IN A      127.0.0.1
_redirect.example.com                         3600 IN TXT    "v=txtv0;from=/$1/$2;root=https://root.example.com;type=path"
_redirect.secondmatch.firstmatch.example.com  3600 IN TXT    "v=txtv0;to=https://about.example.com/;type=host"
```

**Path based redirect path record fallback**
*example.com/firstMatch/noMatch -> fallback.example.com*
```
example.com                                   3600 IN A      127.0.0.1
_redirect.example.com                         3600 IN TXT    "v=txtv0;to=https://fallback.example.com;type=path"
```

**Path based redirect using wildcard**
*example.com/firstMatch/secondMatch -> about.example.com*
*example.com/firstMatch/randomString -> wildcard.example.com*
*example.com/randomString/randomString -> full-wildcard.example.com*

```
example.com                                   3600 IN A      127.0.0.1
_redirect.example.com                         3600 IN TXT    "v=txtv0;type=path"
_redirect.secondmatch.firstmatch.example.com  3600 IN TXT    "v=txtv0;to=https://about.example.com/;type=host"
_redirect._.firstmatch.example.com            3600 IN TXT    "v=txtv0;to=https://wildcard.example.com/;type=host"
_redirect._._.example.com                     3600 IN TXT    "v=txtv0;to=https://full-wildcard.example.com/;type=host"
```

<!--
*example.com/firstMatch/secondMatch -> about.example.com/secondMatch/firstMatch*
*example.com/firstMatch/noMatch -> 404*
```
example.com                                   3600 IN A      127.0.0.1
_redirect.example.com                         3600 IN TXT    "v=txtv0;from=/$1/$2;type=path"
_redirect.secondmatch.firstmatch.example.com  3600 IN TXT    "v=txtv0;to=https://about.example.com/{2}/{1};type=host"
```

*example.com/firstMatch/secondMatch -> about.example.com/secondMatch/firstMatch*
*example.com/firstMatch/noMatch -> fallback.example.com*
```
example.com                                   3600 IN A      127.0.0.1
_redirect.example.com                         3600 IN TXT    "v=txtv0;from=/$1/$2;to=https://fallback.example.com;type=path"
_redirect.secondmatch.firstmatch.example.com  3600 IN TXT    "v=txtv0;to=https://about.example.com/{2}/{1};type=host"
```

*example.com/firstMatch/secondMatch -> about.example.com*
*example.com/firstMatch/noMatch -> fallback.example.com*
```
example.com                                   3600 IN A      127.0.0.1
_redirect.example.com                         3600 IN TXT    "v=txtv0;re=\/(.*)\/(.*);to=https://fallback.example.com;type=path"
_redirect.secondmatch.firstmatch.example.com  3600 IN TXT    "v=txtv0;to=https://about.example.com;type=host"
```

*example.com/some/thing -> catchall.example.com*
*example.com/another/thing -> catchall.example.com*
*example.com/so/many/things -> catchall.example.com/things*
```
example.com                           3600 IN A      127.0.0.1
_redirect.example.com                 3600 IN TXT    "v=txtv0;from=/$1/$2;type=path"
_redirect._._.example.com             3600 IN TXT    "v=txtv0;to=https://catchall.example.com{uri};type=host"
```
-->

## gometa
*pkg.example.com -> github.com/some/repo*
```
pkg.example.com               3600 IN CNAME  txtdirect.example.com.
_redirect.pkg.example.com     3600 IN TXT    "v=txtv0;to=https://github.com/some/repo;type=gometa"
```

<!--
*example.com/somePackage -> github.com/some/repo/somePackage*
```
pkg.example.com               3600 IN CNAME  txtdirect.example.com.
_redirect.pkg.example.com     3600 IN TXT    "v=txtv0;to=https://github.com/some/repo{uri};type=gometa"
```

## gometa + path
*example.com/pkg/fmt -> github.com/pkg/fmt*
```
example.com                     3600 IN A      127.0.0.1
_redirect.example.com           3600 IN TXT    "v=txtv0;from=/$1/$2;to=https://fallback.example.com;type=path"
_redirect.fmt.pkg.example.com   3600 IN TXT    "v=txtv0;to=https://github.com/somePackage/someFmt;type=gometa"
_redirect._._.example.com       3600 IN TXT    "v=txtv0;to=https://fallback.example.com;type=gometa"

```

*example.com/firstMatch/secondMatch -> github.com/somePackage/SomeFmt*
```
example.com                                   3600 IN A      127.0.0.1
_redirect.example.com                         3600 IN TXT    "v=txtv0;re=\/(.*)\/(.*);to=https://fallback.example.com;type=path"
_redirect.secondmatch.firstmatch.example.com  3600 IN TXT    "v=txtv0;to=https://github.com/somePackage/someFmt;type=gometa"
_redirect._._.example.com                     3600 IN TXT    "v=txtv0;to=https://fallback.example.com;type=gometa"
```

*example.com/pkg/fmt -> github.com/somePackage/fmt*
*example.com/pkg2/fmt -> github.com/anotherRepo/fmt/anotherPackage*
```
example.com                     3600 IN A      127.0.0.1
_redirect.example.com           3600 IN TXT    "v=txtv0;from=/$1/$2;to=https://fallback.example.com;type=path"
_redirect._.pkg.example.com     3600 IN TXT    "v=txtv0;to=https://github.com/somePackage/{2};type=gometa"
_redirect._.pkg2.example.com    3600 IN TXT    "v=txtv0;to=https://github.com/anotherRepo/{2}/anotherPackage;type=gometa"
```

*example.com/pkg/fmt -> github.com/pkg/fmt*
*example.com/pkg/fmt23 -> github.com/pkg/fmt23*
*example.com/pkg23/fmt42 -> github.com/pkg23/fmt42*
*example.com/pkg/area51 -> github.com/pkg23/fmt42*
```
example.com                       3600 IN A      127.0.0.1
_redirect.example.com             3600 IN TXT    "v=txtv0;from=/$1/$2;to=https://fallback.example.com;type=path"
_redirect._._.example.com         3600 IN TXT    "v=txtv0;to=https://github.com/{1}/{2};type=gometa"
_redirect.area51.pkg.example.com  3600 IN TXT    "v=txtv0;to=https://github.com/secret/package;type=gometa"
```

## dep
*pkg.example.com/somePackage -> github.com/user/repo/somePackage*
```
import "pkg.example.com"

_redirect.pkg.example.com     3600 IN TXT    "v=txtv0;to=https://github.com/some/repo;type=dep"
```

*example.com/somePackage -> github.com/user/repo/somePackage*
```
import "example.com/somePackage"

_redirect.example.com     3600 IN TXT    "v=txtv0;to=https://github.com/user/repo;type=dep"
```

## dep + path
*example.com/user/fmt -> github.com/user/fmt*
*example.com/user/somePackage/fmt/sub -> github.com/user/repo/fmt/sub*
*example.com/user/somePackage/fmt42 -> github.com/user/repo/fmt42*
*example.com/user/anotherPackage/fmt42 -> mirror.example.com/deps/fmt42*
*example.com/test -> 404* <- 404 for dep tool
*example.com/test -> fallback.example.com* <- website for user traffic
```
_redirect.example.com                     3600 IN TXT    "v=txtv0;from=/$1/$2;to=https://fallback.example.com;type=path"
_redirect.fmt.user.example.com            3600 IN TXT    "v=txtv0;to=https://github.com/user/fmt;type=dep"
_redirect.somePackage.user.example.com    3600 IN TXT    "v=txtv0;to=https://github.com/user/repo;type=dep"
_redirect._.user.example.com              3600 IN TXT    "v=txtv0;to=https://mirror.example.com/deps;type=dep"
```

*example.com/firstMatch/secondMatch -> github.com/somePackage/SomeFmt*
*example.com/failure/test -> 404*
```
_redirect.example.com                           3600 IN TXT    "v=txtv0;re=\/(.*)\/(.*);type=path"
_redirect.secondmatch.firstmatch.example.com    3600 IN TXT    "v=txtv0;to=https://github.com/user/repo;type=dep"
```

*example.com/pkg/area51/test -> github.com/secret/package/test*
*example.com/pkg/fmt -> github.com/pkg/fmt*
*example.com/pkg/fmt23 -> github.com/pkg/fmt23*
*example.com/pkg23/fmt42 -> github.com/pkg23/fmt42*
*example.com/pkg/area51 -> github.com/pkg23/fmt42*
```
_redirect.example.com             3600 IN TXT    "v=txtv0;from=/$1/$2;type=path"
_redirect.area51.pkg.example.com  3600 IN TXT    "v=txtv0;to=https://github.com/secret/package;type=dep"
_redirect._._.example.com         3600 IN TXT    "v=txtv0;to=https://github.com/{1}/{2};type=dep"
```

## dockerv2
*container.example.com -> hub.example.com/some/container*
```
container.example.com             3600 IN CNAME  txtdirect.example.com
_redirect.container.example.com   3600 IN TXT    "v=txtv0;https://hub.example.com/some/container;type=dockerv2"
```

*container.example.com/image -> hub.example.com/some/image*
*container.example.com/image42 -> hub.example.com/some/image42*
```
container.example.com             3600 IN CNAME  txtdirect.example.com
_redirect.container.example.com   3600 IN TXT    "v=txtv0;https://hub.example.com/some{uri};type=dockerv2"
```

## dockerv2 + path
*example.com/con/img -> hub.docker.com/con/img*
*example.com/con/img23 -> hub.docker.com/con/img23*
*example.com/con23/img42 -> hub.docker.com/con23/img42*
*example.com/con/image51 -> hub.docker.com/secret/image*
```
example.com                       3600 IN A      127.0.0.1
_redirect.example.com             3600 IN TXT    "v=txtv0;from=/$1/$2;to=https://fallback.example.com;type=path"
_redirect._._.example.com         3600 IN TXT    "v=txtv0;to=https://hub.docker.com/{1}/{2};type=dockerv2"
_redirect.area51.con.example.com  3600 IN TXT    "v=txtv0;to=https://hub.docker.com/secret/image;type=dockerv2"
```
-->
---

We are happy to get new contributions.

See how you can contribute with our [contribution guide](/CONTRIBUTING.md).
