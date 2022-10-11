## jenkins

```
{
  "StreamConfig": {
    
  },
  "State": {
    "Running": false,
    "Paused": false,
    "Restarting": false,
    "OOMKilled": false,
    "RemovalInProgress": false,
    "Dead": false,
    "Pid": 0,
    "ExitCode": 255,
    "Error": "",
    "StartedAt": "2022-01-06T06:17:32.022076154Z",
    "FinishedAt": "2022-10-10T19:10:11.565285358+08:00",
    "Health": null
  },
  "ID": "55d3066daffb1126896dc086c902917b883cce0d002c1f20063ddb597e5712df",
  "Created": "2021-08-20T00:29:14.849735026Z",
  "Managed": false,
  "Path": "/sbin/tini",
  "Args": [
    "--",
    "/usr/local/bin/jenkins.sh"
  ],
  "Config": {
    "Hostname": "55d3066daffb",
    "Domainname": "",
    "User": "jenkins",
    "AttachStdin": false,
    "AttachStdout": false,
    "AttachStderr": false,
    "ExposedPorts": {
      "50000/tcp": {
        
      },
      "8080/tcp": {
        
      }
    },
    "Tty": true,
    "OpenStdin": true,
    "StdinOnce": false,
    "Env": [
      "PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin",
      "LANG=C.UTF-8",
      "JAVA_HOME=/docker-java-home",
      "JAVA_VERSION=8u212",
      "JAVA_DEBIAN_VERSION=8u212-b01-1~deb9u1",
      "JENKINS_HOME=/var/jenkins_home",
      "JENKINS_SLAVE_AGENT_PORT=50000",
      "JENKINS_VERSION=2.291",
      "JENKINS_UC=https://updates.jenkins.io",
      "JENKINS_UC_EXPERIMENTAL=https://updates.jenkins.io/experimental",
      "JENKINS_INCREMENTALS_REPO_MIRROR=https://repo.jenkins-ci.org/incrementals",
      "COPY_REFERENCE_FILE_LOG=/var/jenkins_home/copy_reference_file.log"
    ],
    "Cmd": null,
    "Image": "jenkins4eval/jenkins:2.291",
    "Volumes": {
      "/var/jenkins_home": {
        
      }
    },
    "WorkingDir": "",
    "Entrypoint": [
      "/sbin/tini",
      "--",
      "/usr/local/bin/jenkins.sh"
    ],
    "OnBuild": null,
    "Labels": {
      
    }
  },
  "Image": "sha256:e26481bc83986ecc73a3426382f7939bbda252237096aea8cad7304e85d4d82a",
  "NetworkSettings": {
    "Bridge": "",
    "SandboxID": "076c3dd6d51f935528eec8df44c989eb951ce9aac9446b69c4d15d9b98a30d48",
    "HairpinMode": false,
    "LinkLocalIPv6Address": "",
    "LinkLocalIPv6PrefixLen": 0,
    "Networks": {
      "bridge": {
        "IPAMConfig": null,
        "Links": null,
        "Aliases": null,
        "NetworkID": "980f98611226baea59dd672a3f1864b54696e87704db7276f8e7a9c5a1de0254",
        "EndpointID": "640c2d196449c79f60c40957d8053b9468e7476bef1ade02c535d12ee6cedb11",
        "Gateway": "172.17.0.1",
        "IPAddress": "172.17.0.4",
        "IPPrefixLen": 16,
        "IPv6Gateway": "",
        "GlobalIPv6Address": "",
        "GlobalIPv6PrefixLen": 0,
        "MacAddress": "02:42:ac:11:00:04",
        "DriverOpts": null,
        "IPAMOperational": false
      }
    },
    "Service": null,
    "Ports": {
      "50000/tcp": [
        {
          "HostIp": "0.0.0.0",
          "HostPort": "50000"
        }
      ],
      "8080/tcp": [
        {
          "HostIp": "0.0.0.0",
          "HostPort": "58080"
        }
      ]
    },
    "SandboxKey": "/var/run/docker/netns/076c3dd6d51f",
    "SecondaryIPAddresses": null,
    "SecondaryIPv6Addresses": null,
    "IsAnonymousEndpoint": false,
    "HasSwarmEndpoint": false
  },
  "LogPath": "/home/docker/lib/containers/55d3066daffb1126896dc086c902917b883cce0d002c1f20063ddb597e5712df/55d3066daffb1126896dc086c902917b883cce0d002c1f20063ddb597e5712df-json.log",
  "Name": "/jenkins",
  "Driver": "overlay2",
  "OS": "linux",
  "MountLabel": "",
  "ProcessLabel": "",
  "RestartCount": 0,
  "HasBeenStartedBefore": true,
  "HasBeenManuallyStopped": false,
  "MountPoints": {
    "/etc/localtime": {
      "Source": "/etc/localtime",
      "Destination": "/etc/localtime",
      "RW": true,
      "Name": "",
      "Driver": "",
      "Type": "bind",
      "Propagation": "rprivate",
      "Spec": {
        "Type": "bind",
        "Source": "/etc/localtime",
        "Target": "/etc/localtime"
      },
      "SkipMountpointCreation": false
    },
    "/var/jenkins_home": {
      "Source": "/home/docker/v/jenkins",
      "Destination": "/var/jenkins_home",
      "RW": true,
      "Name": "",
      "Driver": "",
      "Type": "bind",
      "Propagation": "rprivate",
      "Spec": {
        "Type": "bind",
        "Source": "/home/docker/v/jenkins",
        "Target": "/var/jenkins_home"
      },
      "SkipMountpointCreation": false
    }
  },
  "SecretReferences": null,
  "ConfigReferences": null,
  "AppArmorProfile": "",
  "HostnamePath": "/home/docker/lib/containers/55d3066daffb1126896dc086c902917b883cce0d002c1f20063ddb597e5712df/hostname",
  "HostsPath": "/home/docker/lib/containers/55d3066daffb1126896dc086c902917b883cce0d002c1f20063ddb597e5712df/hosts",
  "ShmPath": "",
  "ResolvConfPath": "/home/docker/lib/containers/55d3066daffb1126896dc086c902917b883cce0d002c1f20063ddb597e5712df/resolv.conf",
  "SeccompProfile": "",
  "NoNewPrivileges": false
}


docker run -dti \
  -p 58080:8080 -p 50000:50000 \
  -v /etc/localtime:/etc/localtime \
  -v /home/docker/v/jenkins:/var/jenkins_home \
  --restart always \
  --name jenkins \
  --privileged \
  jenkins4eval/jenkins:2.291
```

# mysql

```
{
  "StreamConfig": {
    
  },
  "State": {
    "Running": false,
    "Paused": false,
    "Restarting": false,
    "OOMKilled": false,
    "RemovalInProgress": false,
    "Dead": false,
    "Pid": 0,
    "ExitCode": 255,
    "Error": "",
    "StartedAt": "2022-01-06T06:17:31.783579377Z",
    "FinishedAt": "2022-10-10T19:10:11.565686287+08:00",
    "Health": {
      "Status": "healthy",
      "FailingStreak": 0,
      "Log": [
        {
          "Start": "2022-09-11T13:40:56.829798272+08:00",
          "End": "2022-09-11T13:40:57.078332352+08:00",
          "ExitCode": 0,
          "Output": "\u0007mysqladmin: connect to server at 'localhost' failed\nerror: 'Access denied for user 'healthchecker'@'localhost' (using password: YES)'\n"
        },
        {
          "Start": "2022-09-11T13:41:27.277552538+08:00",
          "End": "2022-09-11T13:41:37.079900608+08:00",
          "ExitCode": 0,
          "Output": "\u0007mysqladmin: connect to server at 'localhost' failed\nerror: 'Access denied for user 'healthchecker'@'localhost' (using password: YES)'\n"
        },
        {
          "Start": "2022-09-11T13:42:07.260331732+08:00",
          "End": "2022-09-11T13:42:07.515588149+08:00",
          "ExitCode": 0,
          "Output": "\u0007mysqladmin: connect to server at 'localhost' failed\nerror: 'Access denied for user 'healthchecker'@'localhost' (using password: YES)'\n"
        },
        {
          "Start": "2022-09-11T13:42:39.710993557+08:00",
          "End": "2022-09-11T13:42:39.970789081+08:00",
          "ExitCode": 0,
          "Output": "\u0007mysqladmin: connect to server at 'localhost' failed\nerror: 'Access denied for user 'healthchecker'@'localhost' (using password: YES)'\n"
        },
        {
          "Start": "2022-09-11T13:43:10.150005803+08:00",
          "End": "2022-09-11T13:43:10.438154775+08:00",
          "ExitCode": 0,
          "Output": "\u0007mysqladmin: connect to server at 'localhost' failed\nerror: 'Access denied for user 'healthchecker'@'localhost' (using password: YES)'\n"
        }
      ]
    }
  },
  "ID": "f6bdc3de22214bec90f6809abc271437d49d764b4e2b5d577af0c7aaff2932b3",
  "Created": "2021-08-21T02:47:00.919397771Z",
  "Managed": false,
  "Path": "/entrypoint.sh",
  "Args": [
    "mysqld"
  ],
  "Config": {
    "Hostname": "f6bdc3de2221",
    "Domainname": "",
    "User": "",
    "AttachStdin": false,
    "AttachStdout": false,
    "AttachStderr": false,
    "ExposedPorts": {
      "22/tcp": {
        
      },
      "3306/tcp": {
        
      },
      "33060/tcp": {
        
      }
    },
    "Tty": true,
    "OpenStdin": true,
    "StdinOnce": false,
    "Env": [
      "PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin"
    ],
    "Cmd": [
      "mysqld"
    ],
    "Healthcheck": {
      "Test": [
        "CMD-SHELL",
        "/healthcheck.sh"
      ]
    },
    "Image": "ibex/debian-mysql-server-5.7:5.7.26",
    "Volumes": {
      "/var/lib/mysql": {
        
      }
    },
    "WorkingDir": "",
    "Entrypoint": [
      "/entrypoint.sh"
    ],
    "OnBuild": null,
    "Labels": {
      
    }
  },
  "Image": "sha256:a53b3adb3cdce5449d8ae0a3e2aefe6d098e82862ad9b34ddee436cd5e360bc3",
  "NetworkSettings": {
    "Bridge": "",
    "SandboxID": "3a9e66762b1090e6ed1203d6412be564c56eb48dc3f52111384e54a45d06af73",
    "HairpinMode": false,
    "LinkLocalIPv6Address": "",
    "LinkLocalIPv6PrefixLen": 0,
    "Networks": {
      "bridge": {
        "IPAMConfig": null,
        "Links": null,
        "Aliases": null,
        "NetworkID": "980f98611226baea59dd672a3f1864b54696e87704db7276f8e7a9c5a1de0254",
        "EndpointID": "b780ed9307a9d9b8304e010a1f5d97dedbd6f22b4139cfe93b7b623dbc777a3c",
        "Gateway": "172.17.0.1",
        "IPAddress": "172.17.0.2",
        "IPPrefixLen": 16,
        "IPv6Gateway": "",
        "GlobalIPv6Address": "",
        "GlobalIPv6PrefixLen": 0,
        "MacAddress": "02:42:ac:11:00:02",
        "DriverOpts": null,
        "IPAMOperational": false
      }
    },
    "Service": null,
    "Ports": {
      "22/tcp": [
        {
          "HostIp": "0.0.0.0",
          "HostPort": "3322"
        }
      ],
      "3306/tcp": [
        {
          "HostIp": "0.0.0.0",
          "HostPort": "3306"
        }
      ],
      "33060/tcp": null
    },
    "SandboxKey": "/var/run/docker/netns/3a9e66762b10",
    "SecondaryIPAddresses": null,
    "SecondaryIPv6Addresses": null,
    "IsAnonymousEndpoint": false,
    "HasSwarmEndpoint": false
  },
  "LogPath": "/home/docker/lib/containers/f6bdc3de22214bec90f6809abc271437d49d764b4e2b5d577af0c7aaff2932b3/f6bdc3de22214bec90f6809abc271437d49d764b4e2b5d577af0c7aaff2932b3-json.log",
  "Name": "/mysql",
  "Driver": "overlay2",
  "OS": "linux",
  "MountLabel": "",
  "ProcessLabel": "",
  "RestartCount": 0,
  "HasBeenStartedBefore": true,
  "HasBeenManuallyStopped": false,
  "MountPoints": {
    "/etc/localtime": {
      "Source": "/etc/localtime",
      "Destination": "/etc/localtime",
      "RW": true,
      "Name": "",
      "Driver": "",
      "Type": "bind",
      "Propagation": "rprivate",
      "Spec": {
        "Type": "bind",
        "Source": "/etc/localtime",
        "Target": "/etc/localtime"
      },
      "SkipMountpointCreation": false
    },
    "/var/lib/mysql": {
      "Source": "/home/docker/v/mysql",
      "Destination": "/var/lib/mysql",
      "RW": true,
      "Name": "",
      "Driver": "",
      "Type": "bind",
      "Propagation": "rprivate",
      "Spec": {
        "Type": "bind",
        "Source": "/home/docker/v/mysql",
        "Target": "/var/lib/mysql"
      },
      "SkipMountpointCreation": false
    }
  },
  "SecretReferences": null,
  "ConfigReferences": null,
  "AppArmorProfile": "",
  "HostnamePath": "/home/docker/lib/containers/f6bdc3de22214bec90f6809abc271437d49d764b4e2b5d577af0c7aaff2932b3/hostname",
  "HostsPath": "/home/docker/lib/containers/f6bdc3de22214bec90f6809abc271437d49d764b4e2b5d577af0c7aaff2932b3/hosts",
  "ShmPath": "",
  "ResolvConfPath": "/home/docker/lib/containers/f6bdc3de22214bec90f6809abc271437d49d764b4e2b5d577af0c7aaff2932b3/resolv.conf",
  "SeccompProfile": "",
  "NoNewPrivileges": false
}



docker run -dti --name=mysql \
  --restart always \
  -p 3306:3306 -p 3322:22 \
  -v /etc/localtime:/etc/localtime \
  -v /home/docker/v/mysql:/var/lib/mysql \
  ibex/debian-mysql-server-5.7:5.7.26
```


## Node Exporter

```
{
  "StreamConfig": {
    
  },
  "State": {
    "Running": false,
    "Paused": false,
    "Restarting": false,
    "OOMKilled": false,
    "RemovalInProgress": false,
    "Dead": false,
    "Pid": 0,
    "ExitCode": 255,
    "Error": "",
    "StartedAt": "2022-01-06T06:17:31.149332813Z",
    "FinishedAt": "2022-10-10T19:10:11.566365774+08:00",
    "Health": null
  },
  "ID": "a7c70c370760811bf265e543e3d2a75275a05b54dc0d9ecc99ef157cab7517fe",
  "Created": "2020-08-09T09:31:25.335470789Z",
  "Managed": false,
  "Path": "./netreporter",
  "Args": [
    
  ],
  "Config": {
    "Hostname": "pi8",
    "Domainname": "",
    "User": "",
    "AttachStdin": false,
    "AttachStdout": false,
    "AttachStderr": false,
    "Tty": true,
    "OpenStdin": true,
    "StdinOnce": false,
    "Env": [
      "PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin"
    ],
    "Cmd": [
      "./netreporter"
    ],
    "Image": "netreporter:v1",
    "Volumes": null,
    "WorkingDir": "/",
    "Entrypoint": null,
    "OnBuild": null,
    "Labels": {
      
    }
  },
  "Image": "sha256:8e1c590ceb03fe9e9fc21f4c0372c3892c546bb910a5ec456ac133805bc74cdf",
  "NetworkSettings": {
    "Bridge": "",
    "SandboxID": "d3d0f5e03c159c9d6450b4e2a4c6d5d85c3278a193febceeaf893b4c69862c2e",
    "HairpinMode": false,
    "LinkLocalIPv6Address": "",
    "LinkLocalIPv6PrefixLen": 0,
    "Networks": {
      "host": {
        "IPAMConfig": null,
        "Links": null,
        "Aliases": null,
        "NetworkID": "6f465fbd2e13a1f238245c30f2a229ebd64e2700e25c44bcece4d3830f192ac1",
        "EndpointID": "47fc942a7e3251461788028c27254cce6d85617baf2f2e612a7438d15b822169",
        "Gateway": "",
        "IPAddress": "",
        "IPPrefixLen": 0,
        "IPv6Gateway": "",
        "GlobalIPv6Address": "",
        "GlobalIPv6PrefixLen": 0,
        "MacAddress": "",
        "DriverOpts": null,
        "IPAMOperational": false
      }
    },
    "Service": null,
    "Ports": {
      
    },
    "SandboxKey": "/var/run/docker/netns/default",
    "SecondaryIPAddresses": null,
    "SecondaryIPv6Addresses": null,
    "IsAnonymousEndpoint": false,
    "HasSwarmEndpoint": false
  },
  "LogPath": "/home/docker/lib/containers/a7c70c370760811bf265e543e3d2a75275a05b54dc0d9ecc99ef157cab7517fe/a7c70c370760811bf265e543e3d2a75275a05b54dc0d9ecc99ef157cab7517fe-json.log",
  "Name": "/netreporter-1",
  "Driver": "overlay2",
  "OS": "linux",
  "MountLabel": "",
  "ProcessLabel": "",
  "RestartCount": 0,
  "HasBeenStartedBefore": true,
  "HasBeenManuallyStopped": false,
  "MountPoints": {
    "/etc/localtime": {
      "Source": "/etc/localtime",
      "Destination": "/etc/localtime",
      "RW": false,
      "Name": "",
      "Driver": "",
      "Type": "bind",
      "Relabel": "ro",
      "Propagation": "rprivate",
      "Spec": {
        "Type": "bind",
        "Source": "/etc/localtime",
        "Target": "/etc/localtime",
        "ReadOnly": true
      },
      "SkipMountpointCreation": false
    },
    "/home/workspace/logs": {
      "Source": "/home/workspace/logs",
      "Destination": "/home/workspace/logs",
      "RW": true,
      "Name": "",
      "Driver": "",
      "Type": "bind",
      "Propagation": "rprivate",
      "Spec": {
        "Type": "bind",
        "Source": "/home/workspace/logs",
        "Target": "/home/workspace/logs/"
      },
      "SkipMountpointCreation": false
    }
  },
  "SecretReferences": null,
  "ConfigReferences": null,
  "AppArmorProfile": "",
  "HostnamePath": "/home/docker/lib/containers/a7c70c370760811bf265e543e3d2a75275a05b54dc0d9ecc99ef157cab7517fe/hostname",
  "HostsPath": "/home/docker/lib/containers/a7c70c370760811bf265e543e3d2a75275a05b54dc0d9ecc99ef157cab7517fe/hosts",
  "ShmPath": "",
  "ResolvConfPath": "/home/docker/lib/containers/a7c70c370760811bf265e543e3d2a75275a05b54dc0d9ecc99ef157cab7517fe/resolv.conf",
  "SeccompProfile": "",
  "NoNewPrivileges": false
}




```

## Gitlab

```
{
  "StreamConfig": {
    
  },
  "State": {
    "Running": false,
    "Paused": false,
    "Restarting": false,
    "OOMKilled": false,
    "RemovalInProgress": false,
    "Dead": false,
    "Pid": 0,
    "ExitCode": 255,
    "Error": "",
    "StartedAt": "2022-01-06T06:17:32.00700346Z",
    "FinishedAt": "2022-10-10T19:10:11.565285488+08:00",
    "Health": {
      "Status": "healthy",
      "FailingStreak": 0,
      "Log": [
        {
          "Start": "2022-09-11T13:38:30.049991715+08:00",
          "End": "2022-09-11T13:38:31.025421912+08:00",
          "ExitCode": 0,
          "Output": "  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current\n                                 Dload  Upload   Total   Spent    Left  Speed\n\r  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0\r  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0\u003c!DOCTYPE html\u003e\n\u003chtml class=\"\" lang=\"en\"\u003e\n\u003chead prefix=\"og: http://ogp.me/ns#\"\u003e\n\u003cmeta charset=\"utf-8\"\u003e\n\u003clink as=\"style\" href=\"http://localhost/assets/application-5cd37ee959b5338b5fb48eafc6c7290ca1fa60e653292304102cc19a16cc25e4.css\" rel=\"preload\"\u003e\n\u003clink as=\"style\" href=\"http://localhost/assets/highlight/themes/white-31667f75379ab8ff09116fbd1a32c1345c330a14a2f91f085fe7c60467c92131.css\" rel=\"preload\"\u003e\n\n\u003cmeta content=\"IE=edge\" http-equiv=\"X-UA-Compatible\"\u003e\n\n\u003cmeta content=\"object\" property=\"og:type\"\u003e\n\u003cmeta content=\"GitLab\" property=\"og:site_name\"\u003e\n\u003cmeta content=\"Help\" property=\"og:title\"\u003e\n\u003cmeta content=\"As they sow , so let them reap\" property=\"og:description\"\u003e\n\u003cmeta content=\"http://localhost/assets/gitlab_logo-7ae504fe4f68fdebb3c2034e36621930cd36ea87924c11ff65dbcb8ed50dca58.png\" property=\"og:image\"\u003e\n\u003cmeta content=\"64\" property=\"og:image:width\"\u003e\n\u003cmeta content=\"64\" property=\"og:image:height\"\u003e\n\u003cmeta content=\"http://localhost/help\" property=\"og:url\"\u003e\n\u003cmeta content=\"summary\" property=\"twitter:card\"\u003e\n\u003cmeta content=\"Help\" property=\"twitter:title\"\u003e\n\u003cmeta content=\"As they sow , so let them reap\" property=\"twitter:description\"\u003e\n\u003cmeta content=\"http://localhost/assets/gitlab_logo-7ae504fe4f68fdebb3c2034e36621930cd36ea87924c11ff65dbcb8ed50dca58.png\" property=\"twitter:image\"\u003e\n\n\u003ctitle\u003eHelp · GitLab\u003c/title\u003e\n\u003cmeta content=\"As they sow , so let them reap\" name=\"description\"\u003e\n\n\u003clink rel=\"shortcut icon\" type=\"image/png\" href=\"/uploads/-/system/appearance/favicon/1/favicon.ico\" id=\"favicon\" data-original-href=\"/uploads/-/system/appearance/favicon/1/favicon.ico\" /\u003e\n\u003cstyle\u003e\n@keyframes blinking-dot{0%{opacity:1}25%{opacity:0.4}75%{opacity:0.4}100%{opacity:1}}@keyframes blinking-scroll-button{0%{opacity:0.2}50%{opacity:1}100%{opacity:0.2}}@keyframes gl-spinner-rotate{0%{transform:rotate(0)}100%{transform:rotate(360deg)}}body.ui-indigo .navbar-gitlab{background-color:#292961}body.ui-indigo .navbar-gitlab .navbar-collapse{color:#d1d1f0}body.ui-indigo .navbar-gitlab .container-fluid .navbar-toggler{border-left:1px solid #6868b9;color:#d1d1f0}body.ui-indigo .navbar-gitlab .navbar-sub-nav\u003eli\u003ea:hover,body.ui-indigo .navbar-gitlab .navbar-sub-nav\u003eli\u003ea:focus,body.ui-indigo .navbar-gitlab .navbar-sub-nav\u003eli\u003ebutton:hover,body.ui-indigo .navbar-gitlab .navbar-sub-nav\u003eli\u003ebutton:focus,body.ui-indigo .navbar-gitlab .navbar-nav\u003eli\u003ea:hover,body.ui-indigo .navbar-gitlab .navbar-nav\u003eli\u003ea:focus,body.ui-indigo .navbar-gitlab .navbar-nav\u003eli\u003ebutton:hover,body.ui-indigo .navbar-gitlab .navbar-nav\u003eli\u003ebutton:focus{background-color:rgba(209,209,240,0.2)}body.ui-indigo .navbar-gitlab .navbar-sub-nav\u003eli.active\u003ea,body.ui-indigo .navbar-gitlab .navbar-sub-nav\u003eli.active\u003ebutton,body.ui-indigo .navbar-gitlab .navbar-sub-nav\u003eli.dropdown.show\u003ea,body.ui-indigo .navbar-gitlab .navbar-sub-nav\u003eli.dropdown.show\u003ebutton,body.ui-indigo .navbar-gitlab .navbar-nav\u003eli.active\u003ea,body.ui-indigo .navbar-gitlab .navbar-nav\u003eli.active\u003ebutton,body.ui-indigo .navbar-gitlab .navbar-nav\u003eli.dropdown.show\u003ea,body.ui-indigo .navbar-gitlab .navbar-nav\u003eli.dropdown.show\u003ebutton{color:#292961;background-color:#fff}body.ui-indigo .navbar-gitlab .navbar-sub-nav\u003eli.line-separator,body.ui-indigo .navbar-gitlab .navbar-nav\u003eli.line-separator{border-left:1px solid rgba(209,209,240,0.2)}body.ui-indigo .navbar-gitlab .navbar-sub-nav{color:#d1d1f0}body.ui-indigo .navbar-gitlab .nav\u003eli{color:#d1d1f0}body.ui-indigo .navbar-gitlab .nav\u003eli\u003ea .notification-dot{border:2px solid #292961}body.ui-indigo .navbar-gitlab .nav\u003eli\u003ea.header-help-dropdown-toggle .notification-dot{background-color:#d1d1f0}body.ui-indigo .navbar-gitlab .nav\u003eli\u003ea.header-user-dropdown-toggle .header-user-avatar{border-color:#d1d1f0}@media (min-width: 576px){body.ui-indigo .navbar-gitlab .nav\u003eli\u003ea:hover,body.ui-indigo .navbar-gitlab .nav\u003eli\u003ea:focus{b..."
        },
        {
          "Start": "2022-09-11T13:39:31.76932713+08:00",
          "End": "2022-09-11T13:39:32.62681426+08:00",
          "ExitCode": 0,
          "Output": "  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current\n                                 Dload  Upload   Total   Spent    Left  Speed\n\r  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0\u003c!DOCTYPE html\u003e\n\u003chtml class=\"\" lang=\"en\"\u003e\n\u003chead prefix=\"og: http://ogp.me/ns#\"\u003e\n\u003cmeta charset=\"utf-8\"\u003e\n\u003clink as=\"style\" href=\"http://localhost/assets/application-5cd37ee959b5338b5fb48eafc6c7290ca1fa60e653292304102cc19a16cc25e4.css\" rel=\"preload\"\u003e\n\u003clink as=\"style\" href=\"http://localhost/assets/highlight/themes/white-31667f75379ab8ff09116fbd1a32c1345c330a14a2f91f085fe7c60467c92131.css\" rel=\"preload\"\u003e\n\n\u003cmeta content=\"IE=edge\" http-equiv=\"X-UA-Compatible\"\u003e\n\n\u003cmeta content=\"object\" property=\"og:type\"\u003e\n\u003cmeta content=\"GitLab\" property=\"og:site_name\"\u003e\n\u003cmeta content=\"Help\" property=\"og:title\"\u003e\n\u003cmeta content=\"As they sow , so let them reap\" property=\"og:description\"\u003e\n\u003cmeta content=\"http://localhost/assets/gitlab_logo-7ae504fe4f68fdebb3c2034e36621930cd36ea87924c11ff65dbcb8ed50dca58.png\" property=\"og:image\"\u003e\n\u003cmeta content=\"64\" property=\"og:image:width\"\u003e\n\u003cmeta content=\"64\" property=\"og:image:height\"\u003e\n\u003cmeta content=\"http://localhost/help\" property=\"og:url\"\u003e\n\u003cmeta content=\"summary\" property=\"twitter:card\"\u003e\n\u003cmeta content=\"Help\" property=\"twitter:title\"\u003e\n\u003cmeta content=\"As they sow , so let them reap\" property=\"twitter:description\"\u003e\n\u003cmeta content=\"http://localhost/assets/gitlab_logo-7ae504fe4f68fdebb3c2034e36621930cd36ea87924c11ff65dbcb8ed50dca58.png\" property=\"twitter:image\"\u003e\n\n\u003ctitle\u003eHelp · GitLab\u003c/title\u003e\n\u003cmeta content=\"As they sow , so let them reap\" name=\"description\"\u003e\n\n\u003clink rel=\"shortcut icon\" type=\"image/png\" href=\"/uploads/-/system/appearance/favicon/1/favicon.ico\" id=\"favicon\" data-original-href=\"/uploads/-/system/appearance/favicon/1/favicon.ico\" /\u003e\n\u003cstyle\u003e\n@keyframes blinking-dot{0%{opacity:1}25%{opacity:0.4}75%{opacity:0.4}100%{opacity:1}}@keyframes blinking-scroll-button{0%{opacity:0.2}50%{opacity:1}100%{opacity:0.2}}@keyframes gl-spinner-rotate{0%{transform:rotate(0)}100%{transform:rotate(360deg)}}body.ui-indigo .navbar-gitlab{background-color:#292961}body.ui-indigo .navbar-gitlab .navbar-collapse{color:#d1d1f0}body.ui-indigo .navbar-gitlab .container-fluid .navbar-toggler{border-left:1px solid #6868b9;color:#d1d1f0}body.ui-indigo .navbar-gitlab .navbar-sub-nav\u003eli\u003ea:hover,body.ui-indigo .navbar-gitlab .navbar-sub-nav\u003eli\u003ea:focus,body.ui-indigo .navbar-gitlab .navbar-sub-nav\u003eli\u003ebutton:hover,body.ui-indigo .navbar-gitlab .navbar-sub-nav\u003eli\u003ebutton:focus,body.ui-indigo .navbar-gitlab .navbar-nav\u003eli\u003ea:hover,body.ui-indigo .navbar-gitlab .navbar-nav\u003eli\u003ea:focus,body.ui-indigo .navbar-gitlab .navbar-nav\u003eli\u003ebutton:hover,body.ui-indigo .navbar-gitlab .navbar-nav\u003eli\u003ebutton:focus{background-color:rgba(209,209,240,0.2)}body.ui-indigo .navbar-gitlab .navbar-sub-nav\u003eli.active\u003ea,body.ui-indigo .navbar-gitlab .navbar-sub-nav\u003eli.active\u003ebutton,body.ui-indigo .navbar-gitlab .navbar-sub-nav\u003eli.dropdown.show\u003ea,body.ui-indigo .navbar-gitlab .navbar-sub-nav\u003eli.dropdown.show\u003ebutton,body.ui-indigo .navbar-gitlab .navbar-nav\u003eli.active\u003ea,body.ui-indigo .navbar-gitlab .navbar-nav\u003eli.active\u003ebutton,body.ui-indigo .navbar-gitlab .navbar-nav\u003eli.dropdown.show\u003ea,body.ui-indigo .navbar-gitlab .navbar-nav\u003eli.dropdown.show\u003ebutton{color:#292961;background-color:#fff}body.ui-indigo .navbar-gitlab .navbar-sub-nav\u003eli.line-separator,body.ui-indigo .navbar-gitlab .navbar-nav\u003eli.line-separator{border-left:1px solid rgba(209,209,240,0.2)}body.ui-indigo .navbar-gitlab .navbar-sub-nav{color:#d1d1f0}body.ui-indigo .navbar-gitlab .nav\u003eli{color:#d1d1f0}body.ui-indigo .navbar-gitlab .nav\u003eli\u003ea .notification-dot{border:2px solid #292961}body.ui-indigo .navbar-gitlab .nav\u003eli\u003ea.header-help-dropdown-toggle .notification-dot{background-color:#d1d1f0}body.ui-indigo .navbar-gitlab .nav\u003eli\u003ea.header-user-dropdown-toggle .header-user-avatar{border-color:#d1d1f0}@media (min-width: 576px){body.ui-indigo .navbar-gitlab .nav\u003eli\u003ea:hover,body.ui-indigo .navbar-gitlab .nav\u003eli\u003ea:focus{background-color:rgba(209,209,240,0.2)}}body.ui-indigo .navbar-gitlab .nav\u003eli\u003ea:..."
        },
        {
          "Start": "2022-09-11T13:40:32.817220832+08:00",
          "End": "2022-09-11T13:40:33.711121719+08:00",
          "ExitCode": 0,
          "Output": "  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current\n                                 Dload  Upload   Total   Spent    Left  Speed\n\r  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0\u003c!DOCTYPE html\u003e\n\u003chtml class=\"\" lang=\"en\"\u003e\n\u003chead prefix=\"og: http://ogp.me/ns#\"\u003e\n\u003cmeta charset=\"utf-8\"\u003e\n\u003clink as=\"style\" href=\"http://localhost/assets/application-5cd37ee959b5338b5fb48eafc6c7290ca1fa60e653292304102cc19a16cc25e4.css\" rel=\"preload\"\u003e\n\u003clink as=\"style\" href=\"http://localhost/assets/highlight/themes/white-31667f75379ab8ff09116fbd1a32c1345c330a14a2f91f085fe7c60467c92131.css\" rel=\"preload\"\u003e\n\n\u003cmeta content=\"IE=edge\" http-equiv=\"X-UA-Compatible\"\u003e\n\n\u003cmeta content=\"object\" property=\"og:type\"\u003e\n\u003cmeta content=\"GitLab\" property=\"og:site_name\"\u003e\n\u003cmeta content=\"Help\" property=\"og:title\"\u003e\n\u003cmeta content=\"As they sow , so let them reap\" property=\"og:description\"\u003e\n\u003cmeta content=\"http://localhost/assets/gitlab_logo-7ae504fe4f68fdebb3c2034e36621930cd36ea87924c11ff65dbcb8ed50dca58.png\" property=\"og:image\"\u003e\n\u003cmeta content=\"64\" property=\"og:image:width\"\u003e\n\u003cmeta content=\"64\" property=\"og:image:height\"\u003e\n\u003cmeta content=\"http://localhost/help\" property=\"og:url\"\u003e\n\u003cmeta content=\"summary\" property=\"twitter:card\"\u003e\n\u003cmeta content=\"Help\" property=\"twitter:title\"\u003e\n\u003cmeta content=\"As they sow , so let them reap\" property=\"twitter:description\"\u003e\n\u003cmeta content=\"http://localhost/assets/gitlab_logo-7ae504fe4f68fdebb3c2034e36621930cd36ea87924c11ff65dbcb8ed50dca58.png\" property=\"twitter:image\"\u003e\n\n\u003ctitle\u003eHelp · GitLab\u003c/title\u003e\n\u003cmeta content=\"As they sow , so let them reap\" name=\"description\"\u003e\n\n\u003clink rel=\"shortcut icon\" type=\"image/png\" href=\"/uploads/-/system/appearance/favicon/1/favicon.ico\" id=\"favicon\" data-original-href=\"/uploads/-/system/appearance/favicon/1/favicon.ico\" /\u003e\n\u003cstyle\u003e\n@keyframes blinking-dot{0%{opacity:1}25%{opacity:0.4}75%{opacity:0.4}100%{opacity:1}}@keyframes blinking-scroll-button{0%{opacity:0.2}50%{opacity:1}100%{opacity:0.2}}@keyframes gl-spinner-rotate{0%{transform:rotate(0)}100%{transform:rotate(360deg)}}body.ui-indigo .navbar-gitlab{background-color:#292961}body.ui-indigo .navbar-gitlab .navbar-collapse{color:#d1d1f0}body.ui-indigo .navbar-gitlab .container-fluid .navbar-toggler{border-left:1px solid #6868b9;color:#d1d1f0}body.ui-indigo .navbar-gitlab .navbar-sub-nav\u003eli\u003ea:hover,body.ui-indigo .navbar-gitlab .navbar-sub-nav\u003eli\u003ea:focus,body.ui-indigo .navbar-gitlab .navbar-sub-nav\u003eli\u003ebutton:hover,body.ui-indigo .navbar-gitlab .navbar-sub-nav\u003eli\u003ebutton:focus,body.ui-indigo .navbar-gitlab .navbar-nav\u003eli\u003ea:hover,body.ui-indigo .navbar-gitlab .navbar-nav\u003eli\u003ea:focus,body.ui-indigo .navbar-gitlab .navbar-nav\u003eli\u003ebutton:hover,body.ui-indigo .navbar-gitlab .navbar-nav\u003eli\u003ebutton:focus{background-color:rgba(209,209,240,0.2)}body.ui-indigo .navbar-gitlab .navbar-sub-nav\u003eli.active\u003ea,body.ui-indigo .navbar-gitlab .navbar-sub-nav\u003eli.active\u003ebutton,body.ui-indigo .navbar-gitlab .navbar-sub-nav\u003eli.dropdown.show\u003ea,body.ui-indigo .navbar-gitlab .navbar-sub-nav\u003eli.dropdown.show\u003ebutton,body.ui-indigo .navbar-gitlab .navbar-nav\u003eli.active\u003ea,body.ui-indigo .navbar-gitlab .navbar-nav\u003eli.active\u003ebutton,body.ui-indigo .navbar-gitlab .navbar-nav\u003eli.dropdown.show\u003ea,body.ui-indigo .navbar-gitlab .navbar-nav\u003eli.dropdown.show\u003ebutton{color:#292961;background-color:#fff}body.ui-indigo .navbar-gitlab .navbar-sub-nav\u003eli.line-separator,body.ui-indigo .navbar-gitlab .navbar-nav\u003eli.line-separator{border-left:1px solid rgba(209,209,240,0.2)}body.ui-indigo .navbar-gitlab .navbar-sub-nav{color:#d1d1f0}body.ui-indigo .navbar-gitlab .nav\u003eli{color:#d1d1f0}body.ui-indigo .navbar-gitlab .nav\u003eli\u003ea .notification-dot{border:2px solid #292961}body.ui-indigo .navbar-gitlab .nav\u003eli\u003ea.header-help-dropdown-toggle .notification-dot{background-color:#d1d1f0}body.ui-indigo .navbar-gitlab .nav\u003eli\u003ea.header-user-dropdown-toggle .header-user-avatar{border-color:#d1d1f0}@media (min-width: 576px){body.ui-indigo .navbar-gitlab .nav\u003eli\u003ea:hover,body.ui-indigo .navbar-gitlab .nav\u003eli\u003ea:focus{background-color:rgba(209,209,240,0.2)}}body.ui-indigo .navbar-gitlab .nav\u003eli\u003ea:..."
        },
        {
          "Start": "2022-09-11T13:41:33.748829471+08:00",
          "End": "2022-09-11T13:41:37.422876342+08:00",
          "ExitCode": 0,
          "Output": "  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current\n                                 Dload  Upload   Total   Spent    Left  Speed\n\r  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0\u003c!DOCTYPE html\u003e\n\u003chtml class=\"\" lang=\"en\"\u003e\n\u003chead prefix=\"og: http://ogp.me/ns#\"\u003e\n\u003cmeta charset=\"utf-8\"\u003e\n\u003clink as=\"style\" href=\"http://localhost/assets/application-5cd37ee959b5338b5fb48eafc6c7290ca1fa60e653292304102cc19a16cc25e4.css\" rel=\"preload\"\u003e\n\u003clink as=\"style\" href=\"http://localhost/assets/highlight/themes/white-31667f75379ab8ff09116fbd1a32c1345c330a14a2f91f085fe7c60467c92131.css\" rel=\"preload\"\u003e\n\n\u003cmeta content=\"IE=edge\" http-equiv=\"X-UA-Compatible\"\u003e\n\n\u003cmeta content=\"object\" property=\"og:type\"\u003e\n\u003cmeta content=\"GitLab\" property=\"og:site_name\"\u003e\n\u003cmeta content=\"Help\" property=\"og:title\"\u003e\n\u003cmeta content=\"As they sow , so let them reap\" property=\"og:description\"\u003e\n\u003cmeta content=\"http://localhost/assets/gitlab_logo-7ae504fe4f68fdebb3c2034e36621930cd36ea87924c11ff65dbcb8ed50dca58.png\" property=\"og:image\"\u003e\n\u003cmeta content=\"64\" property=\"og:image:width\"\u003e\n\u003cmeta content=\"64\" property=\"og:image:height\"\u003e\n\u003cmeta content=\"http://localhost/help\" property=\"og:url\"\u003e\n\u003cmeta content=\"summary\" property=\"twitter:card\"\u003e\n\u003cmeta content=\"Help\" property=\"twitter:title\"\u003e\n\u003cmeta content=\"As they sow , so let them reap\" property=\"twitter:description\"\u003e\n\u003cmeta content=\"http://localhost/assets/gitlab_logo-7ae504fe4f68fdebb3c2034e36621930cd36ea87924c11ff65dbcb8ed50dca58.png\" property=\"twitter:image\"\u003e\n\n\u003ctitle\u003eHelp · GitLab\u003c/title\u003e\n\u003cmeta content=\"As they sow , so let them reap\" name=\"description\"\u003e\n\n\u003clink rel=\"shortcut icon\" type=\"image/png\" href=\"/uploads/-/system/appearance/favicon/1/favicon.ico\" id=\"favicon\" data-original-href=\"/uploads/-/system/appearance/favicon/1/favicon.ico\" /\u003e\n\u003cstyle\u003e\n@keyframes blinking-dot{0%{opacity:1}25%{opacity:0.4}75%{opacity:0.4}100%{opacity:1}}@keyframes blinking-scroll-button{0%{opacity:0.2}50%{opacity:1}100%{opacity:0.2}}@keyframes gl-spinner-rotate{0%{transform:rotate(0)}100%{transform:rotate(360deg)}}body.ui-indigo .navbar-gitlab{background-color:#292961}body.ui-indigo .navbar-gitlab .navbar-collapse{color:#d1d1f0}body.ui-indigo .navbar-gitlab .container-fluid .navbar-toggler{border-left:1px solid #6868b9;color:#d1d1f0}body.ui-indigo .navbar-gitlab .navbar-sub-nav\u003eli\u003ea:hover,body.ui-indigo .navbar-gitlab .navbar-sub-nav\u003eli\u003ea:focus,body.ui-indigo .navbar-gitlab .navbar-sub-nav\u003eli\u003ebutton:hover,body.ui-indigo .navbar-gitlab .navbar-sub-nav\u003eli\u003ebutton:focus,body.ui-indigo .navbar-gitlab .navbar-nav\u003eli\u003ea:hover,body.ui-indigo .navbar-gitlab .navbar-nav\u003eli\u003ea:focus,body.ui-indigo .navbar-gitlab .navbar-nav\u003eli\u003ebutton:hover,body.ui-indigo .navbar-gitlab .navbar-nav\u003eli\u003ebutton:focus{background-color:rgba(209,209,240,0.2)}body.ui-indigo .navbar-gitlab .navbar-sub-nav\u003eli.active\u003ea,body.ui-indigo .navbar-gitlab .navbar-sub-nav\u003eli.active\u003ebutton,body.ui-indigo .navbar-gitlab .navbar-sub-nav\u003eli.dropdown.show\u003ea,body.ui-indigo .navbar-gitlab .navbar-sub-nav\u003eli.dropdown.show\u003ebutton,body.ui-indigo .navbar-gitlab .navbar-nav\u003eli.active\u003ea,body.ui-indigo .navbar-gitlab .navbar-nav\u003eli.active\u003ebutton,body.ui-indigo .navbar-gitlab .navbar-nav\u003eli.dropdown.show\u003ea,body.ui-indigo .navbar-gitlab .navbar-nav\u003eli.dropdown.show\u003ebutton{color:#292961;background-color:#fff}body.ui-indigo .navbar-gitlab .navbar-sub-nav\u003eli.line-separator,body.ui-indigo .navbar-gitlab .navbar-nav\u003eli.line-separator{border-left:1px solid rgba(209,209,240,0.2)}body.ui-indigo .navbar-gitlab .navbar-sub-nav{color:#d1d1f0}body.ui-indigo .navbar-gitlab .nav\u003eli{color:#d1d1f0}body.ui-indigo .navbar-gitlab .nav\u003eli\u003ea .notification-dot{border:2px solid #292961}body.ui-indigo .navbar-gitlab .nav\u003eli\u003ea.header-help-dropdown-toggle .notification-dot{background-color:#d1d1f0}body.ui-indigo .navbar-gitlab .nav\u003eli\u003ea.header-user-dropdown-toggle .header-user-avatar{border-color:#d1d1f0}@media (min-width: 576px){body.ui-indigo .navbar-gitlab .nav\u003eli\u003ea:hover,body.ui-indigo .navbar-gitlab .nav\u003eli\u003ea:focus{background-color:rgba(209,209,240,0.2)}}body.ui-indigo .navbar-gitlab .nav\u003eli\u003ea:..."
        },
        {
          "Start": "2022-09-11T13:42:37.453146447+08:00",
          "End": "2022-09-11T13:42:38.173846729+08:00",
          "ExitCode": 0,
          "Output": "  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current\n                                 Dload  Upload   Total   Spent    Left  Speed\n\r  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0\u003c!DOCTYPE html\u003e\n\u003chtml class=\"\" lang=\"en\"\u003e\n\u003chead prefix=\"og: http://ogp.me/ns#\"\u003e\n\u003cmeta charset=\"utf-8\"\u003e\n\u003clink as=\"style\" href=\"http://localhost/assets/application-5cd37ee959b5338b5fb48eafc6c7290ca1fa60e653292304102cc19a16cc25e4.css\" rel=\"preload\"\u003e\n\u003clink as=\"style\" href=\"http://localhost/assets/highlight/themes/white-31667f75379ab8ff09116fbd1a32c1345c330a14a2f91f085fe7c60467c92131.css\" rel=\"preload\"\u003e\n\n\u003cmeta content=\"IE=edge\" http-equiv=\"X-UA-Compatible\"\u003e\n\n\u003cmeta content=\"object\" property=\"og:type\"\u003e\n\u003cmeta content=\"GitLab\" property=\"og:site_name\"\u003e\n\u003cmeta content=\"Help\" property=\"og:title\"\u003e\n\u003cmeta content=\"As they sow , so let them reap\" property=\"og:description\"\u003e\n\u003cmeta content=\"http://localhost/assets/gitlab_logo-7ae504fe4f68fdebb3c2034e36621930cd36ea87924c11ff65dbcb8ed50dca58.png\" property=\"og:image\"\u003e\n\u003cmeta content=\"64\" property=\"og:image:width\"\u003e\n\u003cmeta content=\"64\" property=\"og:image:height\"\u003e\n\u003cmeta content=\"http://localhost/help\" property=\"og:url\"\u003e\n\u003cmeta content=\"summary\" property=\"twitter:card\"\u003e\n\u003cmeta content=\"Help\" property=\"twitter:title\"\u003e\n\u003cmeta content=\"As they sow , so let them reap\" property=\"twitter:description\"\u003e\n\u003cmeta content=\"http://localhost/assets/gitlab_logo-7ae504fe4f68fdebb3c2034e36621930cd36ea87924c11ff65dbcb8ed50dca58.png\" property=\"twitter:image\"\u003e\n\n\u003ctitle\u003eHelp · GitLab\u003c/title\u003e\n\u003cmeta content=\"As they sow , so let them reap\" name=\"description\"\u003e\n\n\u003clink rel=\"shortcut icon\" type=\"image/png\" href=\"/uploads/-/system/appearance/favicon/1/favicon.ico\" id=\"favicon\" data-original-href=\"/uploads/-/system/appearance/favicon/1/favicon.ico\" /\u003e\n\u003cstyle\u003e\n@keyframes blinking-dot{0%{opacity:1}25%{opacity:0.4}75%{opacity:0.4}100%{opacity:1}}@keyframes blinking-scroll-button{0%{opacity:0.2}50%{opacity:1}100%{opacity:0.2}}@keyframes gl-spinner-rotate{0%{transform:rotate(0)}100%{transform:rotate(360deg)}}body.ui-indigo .navbar-gitlab{background-color:#292961}body.ui-indigo .navbar-gitlab .navbar-collapse{color:#d1d1f0}body.ui-indigo .navbar-gitlab .container-fluid .navbar-toggler{border-left:1px solid #6868b9;color:#d1d1f0}body.ui-indigo .navbar-gitlab .navbar-sub-nav\u003eli\u003ea:hover,body.ui-indigo .navbar-gitlab .navbar-sub-nav\u003eli\u003ea:focus,body.ui-indigo .navbar-gitlab .navbar-sub-nav\u003eli\u003ebutton:hover,body.ui-indigo .navbar-gitlab .navbar-sub-nav\u003eli\u003ebutton:focus,body.ui-indigo .navbar-gitlab .navbar-nav\u003eli\u003ea:hover,body.ui-indigo .navbar-gitlab .navbar-nav\u003eli\u003ea:focus,body.ui-indigo .navbar-gitlab .navbar-nav\u003eli\u003ebutton:hover,body.ui-indigo .navbar-gitlab .navbar-nav\u003eli\u003ebutton:focus{background-color:rgba(209,209,240,0.2)}body.ui-indigo .navbar-gitlab .navbar-sub-nav\u003eli.active\u003ea,body.ui-indigo .navbar-gitlab .navbar-sub-nav\u003eli.active\u003ebutton,body.ui-indigo .navbar-gitlab .navbar-sub-nav\u003eli.dropdown.show\u003ea,body.ui-indigo .navbar-gitlab .navbar-sub-nav\u003eli.dropdown.show\u003ebutton,body.ui-indigo .navbar-gitlab .navbar-nav\u003eli.active\u003ea,body.ui-indigo .navbar-gitlab .navbar-nav\u003eli.active\u003ebutton,body.ui-indigo .navbar-gitlab .navbar-nav\u003eli.dropdown.show\u003ea,body.ui-indigo .navbar-gitlab .navbar-nav\u003eli.dropdown.show\u003ebutton{color:#292961;background-color:#fff}body.ui-indigo .navbar-gitlab .navbar-sub-nav\u003eli.line-separator,body.ui-indigo .navbar-gitlab .navbar-nav\u003eli.line-separator{border-left:1px solid rgba(209,209,240,0.2)}body.ui-indigo .navbar-gitlab .navbar-sub-nav{color:#d1d1f0}body.ui-indigo .navbar-gitlab .nav\u003eli{color:#d1d1f0}body.ui-indigo .navbar-gitlab .nav\u003eli\u003ea .notification-dot{border:2px solid #292961}body.ui-indigo .navbar-gitlab .nav\u003eli\u003ea.header-help-dropdown-toggle .notification-dot{background-color:#d1d1f0}body.ui-indigo .navbar-gitlab .nav\u003eli\u003ea.header-user-dropdown-toggle .header-user-avatar{border-color:#d1d1f0}@media (min-width: 576px){body.ui-indigo .navbar-gitlab .nav\u003eli\u003ea:hover,body.ui-indigo .navbar-gitlab .nav\u003eli\u003ea:focus{background-color:rgba(209,209,240,0.2)}}body.ui-indigo .navbar-gitlab .nav\u003eli\u003ea:..."
        }
      ]
    }
  },
  "ID": "fc16a5baabfa2a1acffaa841402195ee947dfa43eb71465be9d98f63a206ce0e",
  "Created": "2021-08-19T11:56:00.790626413Z",
  "Managed": false,
  "Path": "/assets/wrapper",
  "Args": [
    
  ],
  "Config": {
    "Hostname": "git.anxpp.com",
    "Domainname": "",
    "User": "",
    "AttachStdin": false,
    "AttachStdout": false,
    "AttachStderr": false,
    "ExposedPorts": {
      "22/tcp": {
        
      },
      "443/tcp": {
        
      },
      "80/tcp": {
        
      }
    },
    "Tty": true,
    "OpenStdin": true,
    "StdinOnce": false,
    "Env": [
      "GITLAB_OMNIBUS_CONFIG=     nginx['redirect_http_to_https'] = false; ",
      "PATH=/opt/gitlab/embedded/bin:/opt/gitlab/bin:/assets:/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin",
      "LANG=C.UTF-8",
      "EDITOR=/bin/vi",
      "TERM=xterm"
    ],
    "Cmd": [
      "/assets/wrapper"
    ],
    "Healthcheck": {
      "Test": [
        "CMD-SHELL",
        "/opt/gitlab/bin/gitlab-healthcheck --fail --max-time 10"
      ],
      "Interval": 60000000000,
      "Timeout": 30000000000,
      "Retries": 5
    },
    "Image": "yrzr/gitlab-ce-arm64v8:14.1.3-ce.0",
    "Volumes": {
      "/etc/gitlab": {
        
      },
      "/var/log/gitlab": {
        
      },
      "/var/opt/gitlab": {
        
      }
    },
    "WorkingDir": "",
    "Entrypoint": null,
    "OnBuild": null,
    "Labels": {
      
    }
  },
  "Image": "sha256:4daf257602ba6c5ea8c7a9977ad986b53fe663db993d777b2be1bffb965ef42b",
  "NetworkSettings": {
    "Bridge": "",
    "SandboxID": "65ae5633a2b63f97de1d30c3a039225b458cf405867510dcaf8519d229f80c3b",
    "HairpinMode": false,
    "LinkLocalIPv6Address": "",
    "LinkLocalIPv6PrefixLen": 0,
    "Networks": {
      "bridge": {
        "IPAMConfig": null,
        "Links": null,
        "Aliases": null,
        "NetworkID": "980f98611226baea59dd672a3f1864b54696e87704db7276f8e7a9c5a1de0254",
        "EndpointID": "2dbc88f411e9acfd438332dfffbc3b2dc539b622d86fef5d0b93b05ac4943418",
        "Gateway": "172.17.0.1",
        "IPAddress": "172.17.0.3",
        "IPPrefixLen": 16,
        "IPv6Gateway": "",
        "GlobalIPv6Address": "",
        "GlobalIPv6PrefixLen": 0,
        "MacAddress": "02:42:ac:11:00:03",
        "DriverOpts": null,
        "IPAMOperational": false
      }
    },
    "Service": null,
    "Ports": {
      "22/tcp": [
        {
          "HostIp": "0.0.0.0",
          "HostPort": "9922"
        }
      ],
      "443/tcp": [
        {
          "HostIp": "0.0.0.0",
          "HostPort": "9443"
        }
      ],
      "80/tcp": [
        {
          "HostIp": "0.0.0.0",
          "HostPort": "9980"
        }
      ]
    },
    "SandboxKey": "/var/run/docker/netns/65ae5633a2b6",
    "SecondaryIPAddresses": null,
    "SecondaryIPv6Addresses": null,
    "IsAnonymousEndpoint": false,
    "HasSwarmEndpoint": false
  },
  "LogPath": "/home/docker/lib/containers/fc16a5baabfa2a1acffaa841402195ee947dfa43eb71465be9d98f63a206ce0e/fc16a5baabfa2a1acffaa841402195ee947dfa43eb71465be9d98f63a206ce0e-json.log",
  "Name": "/gitlab",
  "Driver": "overlay2",
  "OS": "linux",
  "MountLabel": "",
  "ProcessLabel": "",
  "RestartCount": 0,
  "HasBeenStartedBefore": true,
  "HasBeenManuallyStopped": false,
  "MountPoints": {
    "/etc/gitlab": {
      "Source": "/home/docker/v/gitlab/conf",
      "Destination": "/etc/gitlab",
      "RW": true,
      "Name": "",
      "Driver": "",
      "Type": "bind",
      "Relabel": "z",
      "Propagation": "rprivate",
      "Spec": {
        "Type": "bind",
        "Source": "/home/docker/v/gitlab/conf",
        "Target": "/etc/gitlab"
      },
      "SkipMountpointCreation": false
    },
    "/etc/localtime": {
      "Source": "/etc/localtime",
      "Destination": "/etc/localtime",
      "RW": true,
      "Name": "",
      "Driver": "",
      "Type": "bind",
      "Propagation": "rprivate",
      "Spec": {
        "Type": "bind",
        "Source": "/etc/localtime",
        "Target": "/etc/localtime"
      },
      "SkipMountpointCreation": false
    },
    "/var/log/gitlab": {
      "Source": "/home/docker/v/gitlab/logs",
      "Destination": "/var/log/gitlab",
      "RW": true,
      "Name": "",
      "Driver": "",
      "Type": "bind",
      "Relabel": "z",
      "Propagation": "rprivate",
      "Spec": {
        "Type": "bind",
        "Source": "/home/docker/v/gitlab/logs",
        "Target": "/var/log/gitlab"
      },
      "SkipMountpointCreation": false
    },
    "/var/opt/gitlab": {
      "Source": "/home/docker/v/gitlab/data",
      "Destination": "/var/opt/gitlab",
      "RW": true,
      "Name": "",
      "Driver": "",
      "Type": "bind",
      "Relabel": "z",
      "Propagation": "rprivate",
      "Spec": {
        "Type": "bind",
        "Source": "/home/docker/v/gitlab/data",
        "Target": "/var/opt/gitlab"
      },
      "SkipMountpointCreation": false
    }
  },
  "SecretReferences": null,
  "ConfigReferences": null,
  "AppArmorProfile": "",
  "HostnamePath": "/home/docker/lib/containers/fc16a5baabfa2a1acffaa841402195ee947dfa43eb71465be9d98f63a206ce0e/hostname",
  "HostsPath": "/home/docker/lib/containers/fc16a5baabfa2a1acffaa841402195ee947dfa43eb71465be9d98f63a206ce0e/hosts",
  "ShmPath": "",
  "ResolvConfPath": "/home/docker/lib/containers/fc16a5baabfa2a1acffaa841402195ee947dfa43eb71465be9d98f63a206ce0e/resolv.conf",
  "SeccompProfile": "",
  "NoNewPrivileges": false
}


docker run -ti \
  --detach \
  --restart always \
  --name gitlab \
  --privileged \
  --publish 9922:22 \
  --log-driver none \
  --publish 9980:80 \
  --publish 9443:443 \
  --hostname git.anxpp.com \
  --env GITLAB_OMNIBUS_CONFIG=" \
    nginx['redirect_http_to_https'] = false; "\
  -v /etc/localtime:/etc/localtime \
  --volume /home/docker/v/gitlab/conf:/etc/gitlab:z \
  --volume /home/docker/v/gitlab/logs:/var/log/gitlab:z \
  --volume /home/docker/v/gitlab/data:/var/opt/gitlab:z \
  yrzr/gitlab-ce-arm64v8:14.1.3-ce.0




docker run -ti \
  --detach \
  --restart always \
  --name gitlab \
  --privileged \
  --publish 9922:22 \
  --publish 9980:80 \
  --publish 9443:443 \
  --hostname git.anxpp.com \
  --env GITLAB_OMNIBUS_CONFIG=" \
    nginx['redirect_http_to_https'] = false; praefect['logging_level'] = 'error'; gitaly['logging_level'] = 'error'; "\
  --env GITLAB_LOG_LEVEL=error \
  -v /etc/localtime:/etc/localtime \
  --volume /home/docker/v/gitlab/conf:/etc/gitlab:z \
  --volume /home/docker/v/gitlab/logs:/var/log/gitlab:z \
  --volume /home/docker/v/gitlab/data:/var/opt/gitlab:z \
  yrzr/gitlab-ce-arm64v8:14.1.4-ce.0 && docker logs -f gitlab

```

```

{
    "registry-mirrors": ["https://f5klawc6.mirror.aliyuncs.com", "http://hub-mirror.c.163.com"],
    "data-root": "/home/docker/lib",
    "experimental": true,
    "log-driver":"json-file",
    "log-opts":{
        "max-size" :"5m","max-file":"1"
    }
}

```
