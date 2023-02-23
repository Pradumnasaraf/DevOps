This repository contains ansible playbook which is created using roles to Install and Configure Nginx in different linux distro's including CentOS and Ubuntu.

This Nginx repository mainly contains install-nginx.yml and nginx-role which is created using ansible-galaxy.
The nginx-role structure is given below, one can modify each section based on the specific requirement.

<pre>
├── defaults
│  └── main.yml
├── handlers
│   └── main.yml
├── meta
│  └── main.yml
├── tasks
│   └── main.yml
├── templates
│   ├── style.css.j2
│   ├── script.js.j2
│   ├── nginx_ssl.conf.j2
│   ├── nginx.conf.j2
│   ├── index.html.j2
├── tests
│   └── test.yml
└── vars
    └── main.yml
</pre>

The output of the nginx server landing index page look something like this below. You can play around the codes and customize it based on the specific requirements.

![ansible-nginx-output](https://user-images.githubusercontent.com/37767537/220902070-34a46a39-b25f-4934-994a-a7ccc4123970.png)
