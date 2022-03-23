# Nginx Source

Secondary source for Nginx. This source gathers information from Nginx servers and injects `nginx` items into requests that conatiner details of nginx servers. Thie source is triggered by "triggers" rather than being queried directly. The items that can trigger this source are detailed below

## Triggers

This source can be triggered by the following item types being found as the reasult of another query:

* `service`: Any service with a name of `nginx` or `nginx.service` will trigger this source. 
  * This is designed to be triggered against systemd based serices at the moment. This trigger is able to parse out the location of the nginx binary as well as any included arguments, and will use these when querying data. This means that it should support non-standard installs automatcially

## Types

### `nginx`

Returns information about the Nginx server e.g.

```json
{
    "type": "nginx",
    "uniqueAttribute": "configHash",
    "attributes": {
        "attrStruct": {
            "builtBy": "gcc 9.3.0 (Ubuntu 9.3.0-10ubuntu2) ",
            "config": [
                {
                    "Errors": [
                        {
                            "Error": "[Errno 2] No such file or directory: '/tmp/mime.types'",
                            "Line": 18
                        },
                        {
                            "Error": "\"types\" directive is not allowed here in /tmp/crossplane4107191222:73",
                            "Line": 73
                        }
                    ],
                    "File": "/tmp/crossplane4107191222",
                    "Parsed": [
                        {
                            "Args": [
                                "www-data"
                            ],
                            "Directive": "user",
                            "Line": 3
                        },
                        {
                            "Args": [
                                "auto"
                            ],
                            "Directive": "worker_processes",
                            "Line": 4
                        },
                        {
                            "Args": [
                                "1024"
                            ],
                            "Directive": "worker_rlimit_nofile",
                            "Line": 5
                        },
                        {
                            "Args": [
                                "/var/run/nginx.pid"
                            ],
                            "Directive": "pid",
                            "Line": 7
                        },
                        {
                            "Args": [
                                "/etc/nginx/modules-enabled/*.conf"
                            ],
                            "Directive": "include",
                            "Inlcudes": [],
                            "Line": 8
                        },
                        {
                            "Args": [],
                            "Block": [
                                {
                                    "Args": [
                                        "on"
                                    ],
                                    "Directive": "accept_mutex",
                                    "Line": 11
                                },
                                {
                                    "Args": [
                                        "500ms"
                                    ],
                                    "Directive": "accept_mutex_delay",
                                    "Line": 12
                                },
                                {
                                    "Args": [
                                        "1024"
                                    ],
                                    "Directive": "worker_connections",
                                    "Line": 13
                                }
                            ],
                            "Directive": "events",
                            "Line": 10
                        },
                        {
                            "Args": [],
                            "Block": [
                                {
                                    "Args": [
                                        "mime.types"
                                    ],
                                    "Directive": "include",
                                    "Inlcudes": [],
                                    "Line": 18
                                },
                                {
                                    "Args": [
                                        "application/octet-stream"
                                    ],
                                    "Directive": "default_type",
                                    "Line": 19
                                },
                                {
                                    "Args": [
                                        "/var/log/nginx/access.log"
                                    ],
                                    "Directive": "access_log",
                                    "Line": 21
                                },
                                {
                                    "Args": [
                                        "/var/log/nginx/error.log",
                                        "error"
                                    ],
                                    "Directive": "error_log",
                                    "Line": 22
                                },
                                {
                                    "Args": [
                                        "on"
                                    ],
                                    "Directive": "sendfile",
                                    "Line": 25
                                },
                                {
                                    "Args": [
                                        "on"
                                    ],
                                    "Directive": "server_tokens",
                                    "Line": 26
                                },
                                {
                                    "Args": [
                                        "1024"
                                    ],
                                    "Directive": "types_hash_max_size",
                                    "Line": 28
                                },
                                {
                                    "Args": [
                                        "512"
                                    ],
                                    "Directive": "types_hash_bucket_size",
                                    "Line": 29
                                },
                                {
                                    "Args": [
                                        "64"
                                    ],
                                    "Directive": "server_names_hash_bucket_size",
                                    "Line": 31
                                },
                                {
                                    "Args": [
                                        "512"
                                    ],
                                    "Directive": "server_names_hash_max_size",
                                    "Line": 32
                                },
                                {
                                    "Args": [
                                        "65s"
                                    ],
                                    "Directive": "keepalive_timeout",
                                    "Line": 34
                                },
                                {
                                    "Args": [
                                        "100"
                                    ],
                                    "Directive": "keepalive_requests",
                                    "Line": 35
                                },
                                {
                                    "Args": [
                                        "60s"
                                    ],
                                    "Directive": "client_body_timeout",
                                    "Line": 36
                                },
                                {
                                    "Args": [
                                        "60s"
                                    ],
                                    "Directive": "send_timeout",
                                    "Line": 37
                                },
                                {
                                    "Args": [
                                        "5s"
                                    ],
                                    "Directive": "lingering_timeout",
                                    "Line": 38
                                },
                                {
                                    "Args": [
                                        "on"
                                    ],
                                    "Directive": "tcp_nodelay",
                                    "Line": 39
                                },
                                {
                                    "Args": [
                                        "/run/nginx/client_body_temp"
                                    ],
                                    "Directive": "client_body_temp_path",
                                    "Line": 42
                                },
                                {
                                    "Args": [
                                        "10m"
                                    ],
                                    "Directive": "client_max_body_size",
                                    "Line": 43
                                },
                                {
                                    "Args": [
                                        "128k"
                                    ],
                                    "Directive": "client_body_buffer_size",
                                    "Line": 44
                                },
                                {
                                    "Args": [
                                        "/run/nginx/proxy_temp"
                                    ],
                                    "Directive": "proxy_temp_path",
                                    "Line": 45
                                },
                                {
                                    "Args": [
                                        "90s"
                                    ],
                                    "Directive": "proxy_connect_timeout",
                                    "Line": 46
                                },
                                {
                                    "Args": [
                                        "90s"
                                    ],
                                    "Directive": "proxy_send_timeout",
                                    "Line": 47
                                },
                                {
                                    "Args": [
                                        "90s"
                                    ],
                                    "Directive": "proxy_read_timeout",
                                    "Line": 48
                                },
                                {
                                    "Args": [
                                        "32",
                                        "4k"
                                    ],
                                    "Directive": "proxy_buffers",
                                    "Line": 49
                                },
                                {
                                    "Args": [
                                        "8k"
                                    ],
                                    "Directive": "proxy_buffer_size",
                                    "Line": 50
                                },
                                {
                                    "Args": [
                                        "Host",
                                        "$host"
                                    ],
                                    "Directive": "proxy_set_header",
                                    "Line": 51
                                },
                                {
                                    "Args": [
                                        "X-Real-IP",
                                        "$remote_addr"
                                    ],
                                    "Directive": "proxy_set_header",
                                    "Line": 52
                                },
                                {
                                    "Args": [
                                        "X-Forwarded-For",
                                        "$proxy_add_x_forwarded_for"
                                    ],
                                    "Directive": "proxy_set_header",
                                    "Line": 53
                                },
                                {
                                    "Args": [
                                        "X-Forwarded-Proto",
                                        "$scheme"
                                    ],
                                    "Directive": "proxy_set_header",
                                    "Line": 54
                                },
                                {
                                    "Args": [
                                        "Proxy",
                                        ""
                                    ],
                                    "Directive": "proxy_set_header",
                                    "Line": 55
                                },
                                {
                                    "Args": [
                                        "64"
                                    ],
                                    "Directive": "proxy_headers_hash_bucket_size",
                                    "Line": 56
                                },
                                {
                                    "Args": [
                                        "shared:SSL:10m"
                                    ],
                                    "Directive": "ssl_session_cache",
                                    "Line": 58
                                },
                                {
                                    "Args": [
                                        "5m"
                                    ],
                                    "Directive": "ssl_session_timeout",
                                    "Line": 59
                                },
                                {
                                    "Args": [
                                        "TLSv1",
                                        "TLSv1.1",
                                        "TLSv1.2"
                                    ],
                                    "Directive": "ssl_protocols",
                                    "Line": 60
                                },
                                {
                                    "Args": [
                                        "ECDHE-ECDSA-CHACHA20-POLY1305:ECDHE-RSA-CHACHA20-POLY1305:ECDHE-ECDSA-AES128-GCM-SHA256:ECDHE-RSA-AES128-GCM-SHA256:ECDHE-ECDSA-AES256-GCM-SHA384:ECDHE-RSA-AES256-GCM-SHA384:DHE-RSA-AES128-GCM-SHA256:DHE-RSA-AES256-GCM-SHA384:ECDHE-ECDSA-AES128-SHA256:ECDHE-RSA-AES128-SHA256:ECDHE-ECDSA-AES128-SHA:ECDHE-RSA-AES256-SHA384:ECDHE-RSA-AES128-SHA:ECDHE-ECDSA-AES256-SHA384:ECDHE-ECDSA-AES256-SHA:ECDHE-RSA-AES256-SHA:DHE-RSA-AES128-SHA256:DHE-RSA-AES128-SHA:DHE-RSA-AES256-SHA256:DHE-RSA-AES256-SHA:ECDHE-ECDSA-DES-CBC3-SHA:ECDHE-RSA-DES-CBC3-SHA:EDH-RSA-DES-CBC3-SHA:AES128-GCM-SHA256:AES256-GCM-SHA384:AES128-SHA256:AES256-SHA256:AES128-SHA:AES256-SHA:DES-CBC3-SHA:!DSS"
                                    ],
                                    "Directive": "ssl_ciphers",
                                    "Line": 61
                                },
                                {
                                    "Args": [
                                        "on"
                                    ],
                                    "Directive": "ssl_prefer_server_ciphers",
                                    "Line": 62
                                },
                                {
                                    "Args": [
                                        "off"
                                    ],
                                    "Directive": "ssl_stapling",
                                    "Line": 63
                                },
                                {
                                    "Args": [
                                        "off"
                                    ],
                                    "Directive": "ssl_stapling_verify",
                                    "Line": 64
                                },
                                {
                                    "Args": [
                                        "/etc/nginx/conf.d/*.conf"
                                    ],
                                    "Directive": "include",
                                    "Inlcudes": [],
                                    "Line": 67
                                },
                                {
                                    "Args": [
                                        "/etc/nginx/sites-enabled/*"
                                    ],
                                    "Directive": "include",
                                    "Inlcudes": [],
                                    "Line": 68
                                }
                            ],
                            "Directive": "http",
                            "Line": 16
                        },
                        {
                            "Args": [
                                "html",
                                "htm",
                                "shtml"
                            ],
                            "Directive": "text/html",
                            "Line": 74
                        },
                        {
                            "Args": [
                                "css"
                            ],
                            "Directive": "text/css",
                            "Line": 75
                        },
                        {
                            "Args": [
                                "xml"
                            ],
                            "Directive": "text/xml",
                            "Line": 76
                        },
                        {
                            "Args": [
                                "gif"
                            ],
                            "Directive": "image/gif",
                            "Line": 77
                        },
                        {
                            "Args": [
                                "jpeg",
                                "jpg"
                            ],
                            "Directive": "image/jpeg",
                            "Line": 78
                        },
                        {
                            "Args": [
                                "js"
                            ],
                            "Directive": "application/javascript",
                            "Line": 79
                        },
                        {
                            "Args": [
                                "atom"
                            ],
                            "Directive": "application/atom+xml",
                            "Line": 80
                        },
                        {
                            "Args": [
                                "rss"
                            ],
                            "Directive": "application/rss+xml",
                            "Line": 81
                        },
                        {
                            "Args": [
                                "mml"
                            ],
                            "Directive": "text/mathml",
                            "Line": 82
                        },
                        {
                            "Args": [
                                "txt"
                            ],
                            "Directive": "text/plain",
                            "Line": 83
                        },
                        {
                            "Args": [
                                "jad"
                            ],
                            "Directive": "text/vnd.sun.j2me.app-descriptor",
                            "Line": 84
                        },
                        {
                            "Args": [
                                "wml"
                            ],
                            "Directive": "text/vnd.wap.wml",
                            "Line": 85
                        },
                        {
                            "Args": [
                                "htc"
                            ],
                            "Directive": "text/x-component",
                            "Line": 86
                        },
                        {
                            "Args": [
                                "png"
                            ],
                            "Directive": "image/png",
                            "Line": 87
                        },
                        {
                            "Args": [
                                "tif",
                                "tiff"
                            ],
                            "Directive": "image/tiff",
                            "Line": 88
                        },
                        {
                            "Args": [
                                "wbmp"
                            ],
                            "Directive": "image/vnd.wap.wbmp",
                            "Line": 89
                        },
                        {
                            "Args": [
                                "ico"
                            ],
                            "Directive": "image/x-icon",
                            "Line": 90
                        },
                        {
                            "Args": [
                                "jng"
                            ],
                            "Directive": "image/x-jng",
                            "Line": 91
                        },
                        {
                            "Args": [
                                "bmp"
                            ],
                            "Directive": "image/x-ms-bmp",
                            "Line": 92
                        },
                        {
                            "Args": [
                                "svg",
                                "svgz"
                            ],
                            "Directive": "image/svg+xml",
                            "Line": 93
                        },
                        {
                            "Args": [
                                "webp"
                            ],
                            "Directive": "image/webp",
                            "Line": 94
                        },
                        {
                            "Args": [
                                "woff"
                            ],
                            "Directive": "application/font-woff",
                            "Line": 95
                        },
                        {
                            "Args": [
                                "jar",
                                "war",
                                "ear"
                            ],
                            "Directive": "application/java-archive",
                            "Line": 96
                        },
                        {
                            "Args": [
                                "json"
                            ],
                            "Directive": "application/json",
                            "Line": 97
                        },
                        {
                            "Args": [
                                "hqx"
                            ],
                            "Directive": "application/mac-binhex40",
                            "Line": 98
                        },
                        {
                            "Args": [
                                "doc"
                            ],
                            "Directive": "application/msword",
                            "Line": 99
                        },
                        {
                            "Args": [
                                "pdf"
                            ],
                            "Directive": "application/pdf",
                            "Line": 100
                        },
                        {
                            "Args": [
                                "ps",
                                "eps",
                                "ai"
                            ],
                            "Directive": "application/postscript",
                            "Line": 101
                        },
                        {
                            "Args": [
                                "rtf"
                            ],
                            "Directive": "application/rtf",
                            "Line": 102
                        },
                        {
                            "Args": [
                                "m3u8"
                            ],
                            "Directive": "application/vnd.apple.mpegurl",
                            "Line": 103
                        },
                        {
                            "Args": [
                                "xls"
                            ],
                            "Directive": "application/vnd.ms-excel",
                            "Line": 104
                        },
                        {
                            "Args": [
                                "eot"
                            ],
                            "Directive": "application/vnd.ms-fontobject",
                            "Line": 105
                        },
                        {
                            "Args": [
                                "ppt"
                            ],
                            "Directive": "application/vnd.ms-powerpoint",
                            "Line": 106
                        },
                        {
                            "Args": [
                                "wmlc"
                            ],
                            "Directive": "application/vnd.wap.wmlc",
                            "Line": 107
                        },
                        {
                            "Args": [
                                "kml"
                            ],
                            "Directive": "application/vnd.google-earth.kml+xml",
                            "Line": 108
                        },
                        {
                            "Args": [
                                "kmz"
                            ],
                            "Directive": "application/vnd.google-earth.kmz",
                            "Line": 109
                        },
                        {
                            "Args": [
                                "7z"
                            ],
                            "Directive": "application/x-7z-compressed",
                            "Line": 110
                        },
                        {
                            "Args": [
                                "cco"
                            ],
                            "Directive": "application/x-cocoa",
                            "Line": 111
                        },
                        {
                            "Args": [
                                "jardiff"
                            ],
                            "Directive": "application/x-java-archive-diff",
                            "Line": 112
                        },
                        {
                            "Args": [
                                "jnlp"
                            ],
                            "Directive": "application/x-java-jnlp-file",
                            "Line": 113
                        },
                        {
                            "Args": [
                                "run"
                            ],
                            "Directive": "application/x-makeself",
                            "Line": 114
                        },
                        {
                            "Args": [
                                "pl",
                                "pm"
                            ],
                            "Directive": "application/x-perl",
                            "Line": 115
                        },
                        {
                            "Args": [
                                "prc",
                                "pdb"
                            ],
                            "Directive": "application/x-pilot",
                            "Line": 116
                        },
                        {
                            "Args": [
                                "rar"
                            ],
                            "Directive": "application/x-rar-compressed",
                            "Line": 117
                        },
                        {
                            "Args": [
                                "rpm"
                            ],
                            "Directive": "application/x-redhat-package-manager",
                            "Line": 118
                        },
                        {
                            "Args": [
                                "sea"
                            ],
                            "Directive": "application/x-sea",
                            "Line": 119
                        },
                        {
                            "Args": [
                                "swf"
                            ],
                            "Directive": "application/x-shockwave-flash",
                            "Line": 120
                        },
                        {
                            "Args": [
                                "sit"
                            ],
                            "Directive": "application/x-stuffit",
                            "Line": 121
                        },
                        {
                            "Args": [
                                "tcl",
                                "tk"
                            ],
                            "Directive": "application/x-tcl",
                            "Line": 122
                        },
                        {
                            "Args": [
                                "der",
                                "pem",
                                "crt"
                            ],
                            "Directive": "application/x-x509-ca-cert",
                            "Line": 123
                        },
                        {
                            "Args": [
                                "xpi"
                            ],
                            "Directive": "application/x-xpinstall",
                            "Line": 124
                        },
                        {
                            "Args": [
                                "xhtml"
                            ],
                            "Directive": "application/xhtml+xml",
                            "Line": 125
                        },
                        {
                            "Args": [
                                "xspf"
                            ],
                            "Directive": "application/xspf+xml",
                            "Line": 126
                        },
                        {
                            "Args": [
                                "zip"
                            ],
                            "Directive": "application/zip",
                            "Line": 127
                        },
                        {
                            "Args": [
                                "bin",
                                "exe",
                                "dll",
                                "deb",
                                "dmg",
                                "iso",
                                "img",
                                "msi",
                                "msp",
                                "msm"
                            ],
                            "Directive": "application/octet-stream",
                            "Line": 128
                        },
                        {
                            "Args": [
                                "docx"
                            ],
                            "Directive": "application/vnd.openxmlformats-officedocument.wordprocessingml.document",
                            "Line": 129
                        },
                        {
                            "Args": [
                                "xlsx"
                            ],
                            "Directive": "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet",
                            "Line": 130
                        },
                        {
                            "Args": [
                                "pptx"
                            ],
                            "Directive": "application/vnd.openxmlformats-officedocument.presentationml.presentation",
                            "Line": 131
                        },
                        {
                            "Args": [
                                "mid",
                                "midi",
                                "kar"
                            ],
                            "Directive": "audio/midi",
                            "Line": 132
                        },
                        {
                            "Args": [
                                "mp3"
                            ],
                            "Directive": "audio/mpeg",
                            "Line": 133
                        },
                        {
                            "Args": [
                                "ogg"
                            ],
                            "Directive": "audio/ogg",
                            "Line": 134
                        },
                        {
                            "Args": [
                                "m4a"
                            ],
                            "Directive": "audio/x-m4a",
                            "Line": 135
                        },
                        {
                            "Args": [
                                "ra"
                            ],
                            "Directive": "audio/x-realaudio",
                            "Line": 136
                        },
                        {
                            "Args": [
                                "3gpp",
                                "3gp"
                            ],
                            "Directive": "video/3gpp",
                            "Line": 137
                        },
                        {
                            "Args": [
                                "ts"
                            ],
                            "Directive": "video/mp2t",
                            "Line": 138
                        },
                        {
                            "Args": [
                                "mp4"
                            ],
                            "Directive": "video/mp4",
                            "Line": 139
                        },
                        {
                            "Args": [
                                "mpeg",
                                "mpg"
                            ],
                            "Directive": "video/mpeg",
                            "Line": 140
                        },
                        {
                            "Args": [
                                "mov"
                            ],
                            "Directive": "video/quicktime",
                            "Line": 141
                        },
                        {
                            "Args": [
                                "webm"
                            ],
                            "Directive": "video/webm",
                            "Line": 142
                        },
                        {
                            "Args": [
                                "flv"
                            ],
                            "Directive": "video/x-flv",
                            "Line": 143
                        },
                        {
                            "Args": [
                                "m4v"
                            ],
                            "Directive": "video/x-m4v",
                            "Line": 144
                        },
                        {
                            "Args": [
                                "mng"
                            ],
                            "Directive": "video/x-mng",
                            "Line": 145
                        },
                        {
                            "Args": [
                                "asx",
                                "asf"
                            ],
                            "Directive": "video/x-ms-asf",
                            "Line": 146
                        },
                        {
                            "Args": [
                                "wmv"
                            ],
                            "Directive": "video/x-ms-wmv",
                            "Line": 147
                        },
                        {
                            "Args": [
                                "avi"
                            ],
                            "Directive": "video/x-msvideo",
                            "Line": 148
                        }
                    ],
                    "Status": "failed"
                }
            ],
            "configArgs": [
                "--prefix=/etc/nginx",
                "--sbin-path=/usr/sbin/nginx",
                "--modules-path=/usr/lib/nginx/modules",
                "--conf-path=/etc/nginx/nginx.conf",
                "--error-log-path=/var/log/nginx/error.log",
                "--http-log-path=/var/log/nginx/access.log",
                "--pid-path=/var/run/nginx.pid",
                "--lock-path=/var/run/nginx.lock",
                "--http-client-body-temp-path=/var/cache/nginx/client_temp",
                "--http-proxy-temp-path=/var/cache/nginx/proxy_temp",
                "--http-fastcgi-temp-path=/var/cache/nginx/fastcgi_temp",
                "--http-uwsgi-temp-path=/var/cache/nginx/uwsgi_temp",
                "--http-scgi-temp-path=/var/cache/nginx/scgi_temp",
                "--user=nginx",
                "--group=nginx",
                "--with-compat",
                "--with-file-aio",
                "--with-threads",
                "--with-http_addition_module",
                "--with-http_auth_request_module",
                "--with-http_dav_module",
                "--with-http_flv_module",
                "--with-http_gunzip_module",
                "--with-http_gzip_static_module",
                "--with-http_mp4_module",
                "--with-http_random_index_module",
                "--with-http_realip_module",
                "--with-http_secure_link_module",
                "--with-http_slice_module",
                "--with-http_ssl_module",
                "--with-http_stub_status_module",
                "--with-http_sub_module",
                "--with-http_v2_module",
                "--with-mail",
                "--with-mail_ssl_module",
                "--with-stream",
                "--with-stream_realip_module",
                "--with-stream_ssl_module",
                "--with-stream_ssl_preread_module",
                "--with-cc-opt='-g -O2 -fdebug-prefix-map=/data/builder/debuild/nginx-1.20.2/debian/debuild-base/nginx-1.20.2=. -fstack-protector-strong -Wformat -Werror=format-security -Wp,-D_FORTIFY_SOURCE=2 -fPIC'",
                "--with-ld-opt='-Wl,-Bsymbolic-functions -Wl,-z,relro -Wl,-z,now -Wl,--as-needed -pie'"
            ],
            "configHash": "Gm5Qn58Et9MYlqt9UCR7AJ--NFM=",
            "openSSL": "1.1.1f",
            "version": "nginx/1.20.2"
        }
    },
    "context": "test"
}
```

## Config

All configuration options can be provided via the command line or as environment variables:

| Environment Variable | CLI Flag | Automatic | Description |
|----------------------|----------|-----------|-------------|
| `CONFIG`| `--config` | ✅ | Config file location. Can be used instead of the CLI or environment variables if needed |
| `LOG`| `--log` | ✅ | Set the log level. Valid values: panic, fatal, error, warn, info, debug, trace |
| `NATS_SERVERS`| `--nats-servers` | ✅ | A list of NATS servers to connect to |
| `NATS_NAME_PREFIX`| `--nats-name-prefix` | ✅ | A name label prefix. Sources should append a dot and their hostname .{hostname} to this, then set this is the NATS connection name which will be sent to the server on CONNECT to identify the client |
| `NATS_JWT` | `--nats-jwt` | ✅ | The JWT token that should be used to authenticate to NATS, provided in raw format e.g. `eyJ0eXAiOiJKV1Q{...}` |
| `NATS_NKEY_SEED` | `--nats-nkey-seed` | ✅ | The NKey seed which corresponds to the NATS JWT e.g. `SUAFK6QUC{...}` |
| `MAX-PARALLEL`| `--max-parallel` | ✅ | Max number of requests to run in parallel |

### `srcman` config

When running in srcman, all of the above parameters marked with a checkbox are provided automatically, any additional parameters must be provided under the `config` key. These key-value pairs will become files in the `/etc/srcman/config` directory within the container.

```yaml
apiVersion: srcman.example.com/v0
kind: Source
metadata:
  name: nginx-source
spec:
  image: ghcr.io/overmindtech/nginx-source:latest
  replicas: 2
  manager: manager-sample


```

**NOTE:** Remove the above boilerplate once you know what configuration will be required.

### Health Check

The source hosts a health check on `:8080/healthz` which will return an error if NATS is not connected. An example Kubernetes readiness probe is:

```yaml
readinessProbe:
  httpGet:
    path: /healthz
    port: 8080
```

## Development

### Dev Container

Due to the fact that this library currently shells out to Python (specially [crossplane](https://github.com/nginxinc/crossplane)) it requires that Python and crossplane also be installed to be abel to develop and run the tests locally. Additionaly many tests rely on hvaing an active NATS connection. Please use the provided dev container to develop this project, you can build and connect to this container using VSCode's [Remote - Containers](https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.remote-containers) plugin, which VSCode should suggest for you automatically when you load this repo. This handles by Python and NATS for you.

### Testing

Tests in this package can be run using:

```shell
go test ./...
```

### Packaging

Docker images can be created manually using `docker build`, but GitHub actions also exist that are able to create, tag and push images. Images will be build for the `main` branch, and also for any commits tagged with a version such as `v1.2.0`
