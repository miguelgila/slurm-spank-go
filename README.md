# slurm-spank-go

This is a repository that I use to test whether it is possible to have a Slurm SPANK plugin written in C and Go. The part that links with Slurm is written in C, and the logic of the plugin is or should be written in Go.

NOTE: Take this as a personal development repo, possibly not working and most likely not worth replicating or trying to use in production. I am not a proficient developer in any language, and definitely not in C or Go, so there is a big chance that some of the things I try here are either incorrect, or impossible to do. But why not? :nerd_face:

## Pre-requisites

* Slurm 20.02.2
* Golang 1.11

## How to build and run/test

I use a bunch of CentOS-based Docker containers to do this development. You can find them in [miguelgila/docker-centos-slurm](https://github.com/miguelgila/docker-centos-slurm). Once you are in the container, just run

```
cd /shared
git clone git@github.com:miguelgila/slurm-spank-go.git
cd slurm-spank-go
make
make install # assuming /usr
mkdir 
echo 'optional /usr/lib64/bar.so' >> /etc/slurm/plugstack.conf
srun hostname
grep bar /var/log/slurm/*.log # make sure Slurm is logging at level 4 or above 
```

There are no unit tests, so testing is mostly try-and error :grimacing:


