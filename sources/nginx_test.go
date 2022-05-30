package sources

import (
	"context"
	"encoding/json"
	"regexp"
	"testing"

	"github.com/overmindtech/discovery"
	"github.com/overmindtech/multiconn"
	"github.com/overmindtech/nginx-source/triggers"
	"github.com/overmindtech/sdp-go"
)

type TestNginxCommandSource struct {
	// The item and error that should be returned when the nginx -V command is
	// run
	VersionItem  *sdp.Item
	VersionError error

	// The item and error that should be returned when the nginx -T command is
	// run
	ConfigItem  *sdp.Item
	ConfigError error
}

func (s *TestNginxCommandSource) Type() string {
	return "command"
}

func (s *TestNginxCommandSource) Name() string {
	return "TestNginxCommandSource"
}

func (s *TestNginxCommandSource) Contexts() []string {
	return []string{"test"}
}

func (s *TestNginxCommandSource) Hidden() bool {
	return false
}

func (s *TestNginxCommandSource) Get(ctx context.Context, itemContext string, query string) (*sdp.Item, error) {
	if matched, _ := regexp.MatchString(`nginx -V`, query); matched {
		return s.VersionItem, s.VersionError
	}
	if matched, _ := regexp.MatchString(`nginx -T`, query); matched {
		return s.ConfigItem, s.ConfigError
	}

	return nil, nil
}

func (s *TestNginxCommandSource) Find(ctx context.Context, itemContext string) ([]*sdp.Item, error) {
	return []*sdp.Item{}, nil
}

func (s *TestNginxCommandSource) Weight() int {
	return 10
}

// This file contains tests for the ColourNameSource source. It is a good idea
// to write as many exhaustive tests as possible at this level to ensure that
// your source responds correctly to certain requests.
func TestGet(t *testing.T) {
	tests := []SourceTest{
		{
			Name:        "get should fail",
			ItemContext: "something.specific",
			Query:       "irrelevant",
			Method:      sdp.RequestMethod_GET,
			ExpectedError: &ExpectedError{
				Type:             sdp.ItemRequestError_OTHER,
				ErrorStringRegex: regexp.MustCompile(`Get is not supported`),
				Context:          "something.specific",
			},
		},
	}

	RunSourceTests(t, tests, &NginxSource{})
}

func TestFind(t *testing.T) {
	tests := []SourceTest{
		{
			Name:          "find returns no items",
			ItemContext:   "something.specific",
			Method:        sdp.RequestMethod_FIND,
			ExpectedError: nil,
			ExpectedItems: &ExpectedItems{
				NumItems: 0,
			},
		},
	}

	RunSourceTests(t, tests, &NginxSource{})
}

func TestSearch(t *testing.T) {
	var configAttributes *sdp.ItemAttributes
	var versionAttributes *sdp.ItemAttributes
	var err error
	var queryBytes []byte

	configAttributes, err = sdp.ToAttributes(map[string]interface{}{
		"exitCode": 0,
		"name":     "/usr/sbin/nginx -Tq",
		"stderr":   "",
		"stdout":   "# configuration file /etc/nginx/nginx.conf:\n# MANAGED BY PUPPET\nuser www-data;\nworker_processes auto;\nworker_rlimit_nofile 1024;\n\npid        /var/run/nginx.pid;\ninclude /etc/nginx/modules-enabled/*.conf;\n\nevents {\n  accept_mutex on;\n  accept_mutex_delay 500ms;\n  worker_connections 1024;\n}\n\nhttp {\n\n  include       mime.types;\n  default_type  application/octet-stream;\n\n  access_log /var/log/nginx/access.log;\n  error_log /var/log/nginx/error.log error;\n\n\n  sendfile on;\n  server_tokens on;\n\n  types_hash_max_size 1024;\n  types_hash_bucket_size 512;\n\n  server_names_hash_bucket_size 64;\n  server_names_hash_max_size 512;\n\n  keepalive_timeout   65s;\n  keepalive_requests  100;\n  client_body_timeout 60s;\n  send_timeout        60s;\n  lingering_timeout   5s;\n  tcp_nodelay         on;\n\n\n  client_body_temp_path   /run/nginx/client_body_temp;\n  client_max_body_size    10m;\n  client_body_buffer_size 128k;\n  proxy_temp_path         /run/nginx/proxy_temp;\n  proxy_connect_timeout   90s;\n  proxy_send_timeout      90s;\n  proxy_read_timeout      90s;\n  proxy_buffers           32 4k;\n  proxy_buffer_size       8k;\n  proxy_set_header        Host $host;\n  proxy_set_header        X-Real-IP $remote_addr;\n  proxy_set_header        X-Forwarded-For $proxy_add_x_forwarded_for;\n  proxy_set_header        X-Forwarded-Proto $scheme;\n  proxy_set_header        Proxy \"\";\n  proxy_headers_hash_bucket_size 64;\n\n  ssl_session_cache         shared:SSL:10m;\n  ssl_session_timeout       5m;\n  ssl_protocols             TLSv1 TLSv1.1 TLSv1.2;\n  ssl_ciphers               ECDHE-ECDSA-CHACHA20-POLY1305:ECDHE-RSA-CHACHA20-POLY1305:ECDHE-ECDSA-AES128-GCM-SHA256:ECDHE-RSA-AES128-GCM-SHA256:ECDHE-ECDSA-AES256-GCM-SHA384:ECDHE-RSA-AES256-GCM-SHA384:DHE-RSA-AES128-GCM-SHA256:DHE-RSA-AES256-GCM-SHA384:ECDHE-ECDSA-AES128-SHA256:ECDHE-RSA-AES128-SHA256:ECDHE-ECDSA-AES128-SHA:ECDHE-RSA-AES256-SHA384:ECDHE-RSA-AES128-SHA:ECDHE-ECDSA-AES256-SHA384:ECDHE-ECDSA-AES256-SHA:ECDHE-RSA-AES256-SHA:DHE-RSA-AES128-SHA256:DHE-RSA-AES128-SHA:DHE-RSA-AES256-SHA256:DHE-RSA-AES256-SHA:ECDHE-ECDSA-DES-CBC3-SHA:ECDHE-RSA-DES-CBC3-SHA:EDH-RSA-DES-CBC3-SHA:AES128-GCM-SHA256:AES256-GCM-SHA384:AES128-SHA256:AES256-SHA256:AES128-SHA:AES256-SHA:DES-CBC3-SHA:!DSS;\n  ssl_prefer_server_ciphers on;\n  ssl_stapling              off;\n  ssl_stapling_verify       off;\n\n\n  include /etc/nginx/conf.d/*.conf;\n  include /etc/nginx/sites-enabled/*;\n}\n\n# configuration file /etc/nginx/mime.types:\n# MANAGED BY PUPPET\ntypes {\n    text/html html htm shtml;\n    text/css css;\n    text/xml xml;\n    image/gif gif;\n    image/jpeg jpeg jpg;\n    application/javascript js;\n    application/atom+xml atom;\n    application/rss+xml rss;\n    text/mathml mml;\n    text/plain txt;\n    text/vnd.sun.j2me.app-descriptor jad;\n    text/vnd.wap.wml wml;\n    text/x-component htc;\n    image/png png;\n    image/tiff tif tiff;\n    image/vnd.wap.wbmp wbmp;\n    image/x-icon ico;\n    image/x-jng jng;\n    image/x-ms-bmp bmp;\n    image/svg+xml svg svgz;\n    image/webp webp;\n    application/font-woff woff;\n    application/java-archive jar war ear;\n    application/json json;\n    application/mac-binhex40 hqx;\n    application/msword doc;\n    application/pdf pdf;\n    application/postscript ps eps ai;\n    application/rtf rtf;\n    application/vnd.apple.mpegurl m3u8;\n    application/vnd.ms-excel xls;\n    application/vnd.ms-fontobject eot;\n    application/vnd.ms-powerpoint ppt;\n    application/vnd.wap.wmlc wmlc;\n    application/vnd.google-earth.kml+xml kml;\n    application/vnd.google-earth.kmz kmz;\n    application/x-7z-compressed 7z;\n    application/x-cocoa cco;\n    application/x-java-archive-diff jardiff;\n    application/x-java-jnlp-file jnlp;\n    application/x-makeself run;\n    application/x-perl pl pm;\n    application/x-pilot prc pdb;\n    application/x-rar-compressed rar;\n    application/x-redhat-package-manager rpm;\n    application/x-sea sea;\n    application/x-shockwave-flash swf;\n    application/x-stuffit sit;\n    application/x-tcl tcl tk;\n    application/x-x509-ca-cert der pem crt;\n    application/x-xpinstall xpi;\n    application/xhtml+xml xhtml;\n    application/xspf+xml xspf;\n    application/zip zip;\n    application/octet-stream bin exe dll deb dmg iso img msi msp msm;\n    application/vnd.openxmlformats-officedocument.wordprocessingml.document docx;\n    application/vnd.openxmlformats-officedocument.spreadsheetml.sheet xlsx;\n    application/vnd.openxmlformats-officedocument.presentationml.presentation pptx;\n    audio/midi mid midi kar;\n    audio/mpeg mp3;\n    audio/ogg ogg;\n    audio/x-m4a m4a;\n    audio/x-realaudio ra;\n    video/3gpp 3gpp 3gp;\n    video/mp2t ts;\n    video/mp4 mp4;\n    video/mpeg mpeg mpg;\n    video/quicktime mov;\n    video/webm webm;\n    video/x-flv flv;\n    video/x-m4v m4v;\n    video/x-mng mng;\n    video/x-ms-asf asx asf;\n    video/x-ms-wmv wmv;\n    video/x-msvideo avi;\n}\n\n# configuration file /etc/nginx/conf.d/default.conf:\nserver {\n    listen       80;\n    server_name  localhost;\n\n    #access_log  /var/log/nginx/host.access.log  main;\n\n    location / {\n        root   /usr/share/nginx/html;\n        index  index.html index.htm;\n    }\n\n    #error_page  404              /404.html;\n\n    # redirect server error pages to the static page /50x.html\n    #\n    error_page   500 502 503 504  /50x.html;\n    location = /50x.html {\n        root   /usr/share/nginx/html;\n    }\n\n    # proxy the PHP scripts to Apache listening on 127.0.0.1:80\n    #\n    #location ~ \\.php$ {\n    #    proxy_pass   http://127.0.0.1;\n    #}\n\n    # pass the PHP scripts to FastCGI server listening on 127.0.0.1:9000\n    #\n    #location ~ \\.php$ {\n    #    root           html;\n    #    fastcgi_pass   127.0.0.1:9000;\n    #    fastcgi_index  index.php;\n    #    fastcgi_param  SCRIPT_FILENAME  /scripts$fastcgi_script_name;\n    #    include        fastcgi_params;\n    #}\n\n    # deny access to .htaccess files, if Apache's document root\n    # concurs with nginx's one\n    #\n    #location ~ /\\.ht {\n    #    deny  all;\n    #}\n}\n\n",
	})

	if err != nil {
		t.Fatal(err)
	}

	versionAttributes, err = sdp.ToAttributes(map[string]interface{}{
		"exitCode": 0,
		"name":     "/usr/sbin/nginx -V -c /etc/nginx/nginx.conf",
		"stderr":   "nginx version: nginx/1.20.2\nbuilt by gcc 9.3.0 (Ubuntu 9.3.0-10ubuntu2) \nbuilt with OpenSSL 1.1.1f  31 Mar 2020\nTLS SNI support enabled\nconfigure arguments: --prefix=/etc/nginx --sbin-path=/usr/sbin/nginx --modules-path=/usr/lib/nginx/modules --conf-path=/etc/nginx/nginx.conf --error-log-path=/var/log/nginx/error.log --http-log-path=/var/log/nginx/access.log --pid-path=/var/run/nginx.pid --lock-path=/var/run/nginx.lock --http-client-body-temp-path=/var/cache/nginx/client_temp --http-proxy-temp-path=/var/cache/nginx/proxy_temp --http-fastcgi-temp-path=/var/cache/nginx/fastcgi_temp --http-uwsgi-temp-path=/var/cache/nginx/uwsgi_temp --http-scgi-temp-path=/var/cache/nginx/scgi_temp --user=nginx --group=nginx --with-compat --with-file-aio --with-threads --with-http_addition_module --with-http_auth_request_module --with-http_dav_module --with-http_flv_module --with-http_gunzip_module --with-http_gzip_static_module --with-http_mp4_module --with-http_random_index_module --with-http_realip_module --with-http_secure_link_module --with-http_slice_module --with-http_ssl_module --with-http_stub_status_module --with-http_sub_module --with-http_v2_module --with-mail --with-mail_ssl_module --with-stream --with-stream_realip_module --with-stream_ssl_module --with-stream_ssl_preread_module --with-cc-opt='-g -O2 -fdebug-prefix-map=/data/builder/debuild/nginx-1.20.2/debian/debuild-base/nginx-1.20.2=. -fstack-protector-strong -Wformat -Werror=format-security -Wp,-D_FORTIFY_SOURCE=2 -fPIC' --with-ld-opt='-Wl,-Bsymbolic-functions -Wl,-z,relro -Wl,-z,now -Wl,--as-needed -pie'",
		"stdout":   "",
	})

	if err != nil {
		t.Fatal(err)
	}

	trigger := triggers.TriggerData{
		TriggerType: triggers.SERVICE,
		TriggerItemRef: &sdp.Reference{
			Type:                 "service",
			UniqueAttributeValue: "nginx.service",
			Context:              "test",
		},
		ServiceData: &triggers.ServiceData{
			Binary: "/usr/sbin/nginx",
			Args:   []string{"-c", "/etc/nginx/nginx.conf"},
		},
	}

	queryBytes, err = json.Marshal(trigger)

	if err != nil {
		t.Fatal(err)
	}

	responderEngine := discovery.Engine{
		Name:                  "test-responder",
		MaxParallelExecutions: 1,
		NATSOptions: &multiconn.NATSConnectionOptions{
			Servers: []string{
				"nats://nats:4222",
				"nats://localhost:4222",
			},
		},
	}

	responderEngine.AddSources(&TestNginxCommandSource{
		VersionItem: &sdp.Item{
			Type:            "command",
			UniqueAttribute: "name",
			Attributes:      versionAttributes,
			Context:         "test",
		},
		VersionError: nil,
		ConfigItem: &sdp.Item{
			Type:            "command",
			UniqueAttribute: "name",
			Attributes:      configAttributes,
			Context:         "test",
		},
		ConfigError: nil,
	})

	// Start the responder engine so that it can reply to requests
	err = responderEngine.Start()

	if err != nil {
		t.Fatal(err)
	}

	defer responderEngine.Stop()

	tests := []SourceTest{
		{
			Name:        "with no query",
			ItemContext: "something.specific",
			Method:      sdp.RequestMethod_SEARCH,
			ExpectedError: &ExpectedError{
				Type:    sdp.ItemRequestError_OTHER,
				Context: "something.specific",
			},
		},
		{
			Name:        "with a working return",
			ItemContext: "test",
			Method:      sdp.RequestMethod_SEARCH,
			Query:       string(queryBytes),
			ExpectedItems: &ExpectedItems{
				NumItems: 1,
				ExpectedAttributes: []map[string]interface{}{
					{
						"version": "nginx/1.20.2",
						"builtBy": "gcc 9.3.0 (Ubuntu 9.3.0-10ubuntu2) ",
						"openSSL": "1.1.1f",
						"configArgs": []interface{}{
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
							"--with-ld-opt='-Wl,-Bsymbolic-functions -Wl,-z,relro -Wl,-z,now -Wl,--as-needed -pie'",
						},
					},
				},
			},
		},
	}

	// Run another engine for the nginx source so that it actually has to
	// communicate over NATS
	sourceEngine := discovery.Engine{
		Name:                  "nginx-source",
		MaxParallelExecutions: 1,
		NATSOptions: &multiconn.NATSConnectionOptions{
			Servers: []string{
				"nats://nats:4222",
				"nats://localhost:4222",
			},
		},
	}

	err = sourceEngine.Start()

	if err != nil {
		t.Fatal(err)
	}

	defer sourceEngine.Stop()

	RunSourceTests(t, tests, &NginxSource{
		Engine: &sourceEngine,
	})
}
