This repository contains ansible playbook which is created using roles to Install and Configure Tomcat in different linux distro's including CentOS and Ubuntu.

This Tomcat repository mainly contains install-tomcat.yml and tomcat-role-centos and tomcat-role-ubuntu repo's which are created using ansible-galaxy. Along with it this repo contains one sample.war file which is a sample Hello World Application and tomcat-war-deploy.yml file for it's deployment. 
The tomcat-role-centos and tomcat-role-ubuntu structures are given below, one can modify each section based on the specific requirement.

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
│   ├── context.xml.j2
│   ├── tomcat-users.xml.j2
│   ├── tomcat.service.j2
├── tests
│   └── test.yml
└── vars
    └── main.yml
</pre>

The output of the tomcat server landing page look something like this below. You can play around the codes and customize it based on the specific requirements.

![ansible-tomcat output](https://user-images.githubusercontent.com/37767537/220903083-751c8950-d636-44e3-8b83-d194a3423293.png)
