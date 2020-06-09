# slurm-spank-go

This is a repository that I use to test whether it is possible to have a Slurm SPANK plugin written in C and Go. The part that links with Slurm is written in C, and the logic of the plugin is or should be written in Go.

NOTE: Take this as a personal development repo, possibly not working and most likely not worth replicating or trying to use in production. I am not a proficient developer in any language, and definitely not in C or Go, so there is a big chance that some of the things I try here are either incorrect, or impossible to do. But why not? :nerd_face:

## Pre-requisites

* Slurm 20.02.2
* Golang 1.11

## How to build and run/test

I use a bunch of CentOS-based Docker containers to do this development. You can find them in [miguelgila/docker-centos-slurm](https://github.com/miguelgila/docker-centos-slurm). Once you are in the container, just run

```
git clone git@github.com:miguelgila/slurm-spank-go.git
cd slurm-spank-go
make
make install # assuming /usr
echo 'optional /usr/lib64/bar.so' >> /etc/slurm/plugstack.conf
srun hostname
tail -n100 /var/log/slurm/slurmd.log |grep -E 'c:|go:' # make sure Slurm is logging at level 4 or above 
```

This is an example of the output it produces right now:

```
[root@c10 slurm-spank-go]# make clean && make && make install && srun hostname 
rm -f foo.a foo.h bar core foo foo.so foo.h bar.so foo.o bar.o
go build -buildmode=c-shared -o foo.o foo.go
In file included from bar.c:8:0:
foo.go:8:13: warning: ‘my_slurm_verbose’ defined but not used [-Wunused-function]
 static void my_slurm_verbose(char* s) {
             ^
install bar.so /usr/lib64
install foo.o /usr/lib64
c10
[root@c10 slurm-spank-go]# tail -n49 /var/log/slurm/slurmd.log |grep -E 'c:|go:'
[2020-06-09T09:29:54.582] c: Starting slurm_spank_job_prolog
[2020-06-09T09:29:54.583]  go: Starting [Spank_job_prolog]
[2020-06-09T09:29:54.583]  go: Finishing [Spank_job_prolog]
[2020-06-09T09:29:54.583] c: Finishing slurm_spank_task_init
[2020-06-09T09:29:54.624] [31.0] c: Starting slurm_spank_init
[2020-06-09T09:29:54.625] [31.0]  go: Starting [Spank_init]
[2020-06-09T09:29:54.625] [31.0]  go: Will return to C the value v=9
[2020-06-09T09:29:54.626] [31.0]  go: Finishing [Spank_init]
[2020-06-09T09:29:54.626] [31.0] c: Go Spank_init() function returned 9
[2020-06-09T09:29:54.626] [31.0] c: Finishing slurm_spank_init
[2020-06-09T09:29:54.626] [31.0] c: Starting slurm_spank_init_post_opt
[2020-06-09T09:29:54.626] [31.0]  go: Starting [Spank_init_post_opt]
[2020-06-09T09:29:54.626] [31.0]  go: Finishing [Spank_init_post_opt]
[2020-06-09T09:29:54.626] [31.0] c: Finishing slurm_spank_task_init
[2020-06-09T09:29:54.638] [31.0] c: Starting slurm_spank_user_init
[2020-06-09T09:29:54.638] [31.0]  go: Starting [Spank_user_init]
[2020-06-09T09:29:54.638] [31.0]  go: Finishing [Spank_user_init]
[2020-06-09T09:29:54.639] [31.0] c: Finishing slurm_spank_task_init
[2020-06-09T09:29:54.639] [31.0] c: Starting slurm_spank_task_post_fork
[2020-06-09T09:29:54.639] [31.0] c: Finishing slurm_spank_task_exit
[2020-06-09T09:29:54.640] [31.0] c: Starting slurm_spank_task_init_privileged
[2020-06-09T09:29:54.640] [31.0]  go: Starting [Spank_task_init_privileged]
[2020-06-09T09:29:54.640] [31.0]  go: Finishing [Spank_task_init_privileged]
[2020-06-09T09:29:54.640] [31.0] c: Finishing slurm_spank_task_init
[2020-06-09T09:29:54.641] [31.0] c: Starting slurm_spank_task_init
[2020-06-09T09:29:54.642] [31.0]  go: Starting [Spank_task_init]
[2020-06-09T09:29:54.642] [31.0]  go: Finishing [Spank_task_init]
[2020-06-09T09:29:54.642] [31.0] c: Finishing slurm_spank_task_init
[2020-06-09T09:29:54.647] [31.0] c: Starting slurm_spank_task_exit
[2020-06-09T09:29:54.648] [31.0]  go: Starting [Spank_task_exit]
[2020-06-09T09:29:54.648] [31.0]  go: Finishing [Spank_task_exit]
[2020-06-09T09:29:54.648] [31.0] c: Finishing slurm_spank_task_exit
[2020-06-09T09:29:54.650] [31.0] c: Starting slurm_spank_exit
[2020-06-09T09:29:54.650] [31.0]  go: Starting [Spank_exit]
[2020-06-09T09:29:54.650] [31.0]  go: Finishing [Spank_exit]
[2020-06-09T09:29:54.650] [31.0] c: Finishing slurm_spank_exit
[2020-06-09T09:29:54.673] c: Starting slurm_spank_job_epilog
[2020-06-09T09:29:54.675]  go: Starting [Spank_job_epilog]
[2020-06-09T09:29:54.675]  go: Finishing [Spank_job_epilog]
[2020-06-09T09:29:54.675] c: Finishing slurm_spank_job_epilog]
```

There are no unit tests, so testing is mostly try-and error :grimacing:

END
