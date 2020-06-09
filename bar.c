
#include <stdio.h>
#include <stdlib.h>

#include <string.h>
#include <slurm/spank.h>

#include "foo.h"

SPANK_PLUGIN(basic, 1);

int slurm_spank_job_prolog(spank_t spank, int argc, char *argv[])
{
    slurm_verbose("c: Starting slurm_spank_job_prolog");
    Spank_job_prolog();
    slurm_verbose("c: Finishing slurm_spank_task_init");
    return (0);
}
int slurm_spank_init_post_opt(spank_t sp, int ac, char **av)
{
    slurm_verbose("c: Starting slurm_spank_init_post_opt");
    Spank_init_post_opt();
    slurm_verbose("c: Finishing slurm_spank_task_init");
    return (0);
}
int slurm_spank_local_user_init(spank_t sp, int ac, char **av)
{
    slurm_verbose("c: Starting slurm_spank_local_user_init");
    Spank_local_user_init();
    slurm_verbose("c: Finishing slurm_spank_task_init"); 
    return (0);
}
int slurm_spank_user_init(spank_t sp, int ac, char **av)
{
    slurm_verbose("c: Starting slurm_spank_user_init");
    Spank_user_init();
    slurm_verbose("c: Finishing slurm_spank_task_init");
    return (0);
}
int slurm_spank_task_init_privileged(spank_t sp, int ac, char **av)
{
    slurm_verbose("c: Starting slurm_spank_task_init_privileged");
    Spank_task_init_privileged();
    slurm_verbose("c: Finishing slurm_spank_task_init");
    return (0);
}
int slurm_spank_task_init(spank_t sp, int ac, char **av)
{
    slurm_verbose("c: Starting slurm_spank_task_init");
    Spank_task_init();
    slurm_verbose("c: Finishing slurm_spank_task_init");
    return (0);
}
int slurm_spank_task_post_fork(spank_t sp, int ac, char **av)
{
    slurm_verbose("c: Starting slurm_spank_task_post_fork");
    // This will NOT work, why?? FIXME
    //Spank_task_post_fork();
    slurm_verbose("c: Finishing slurm_spank_task_exit");
    return (0);
}
int slurm_spank_task_exit(spank_t sp, int ac, char **av)
{
    slurm_verbose("c: Starting slurm_spank_task_exit");
    Spank_task_exit();
    slurm_verbose("c: Finishing slurm_spank_task_exit");
    return (0);
}
int slurm_spank_exit(spank_t sp, int ac, char **av)
{
    slurm_verbose("c: Starting slurm_spank_exit");
    Spank_exit();
    slurm_verbose("c: Finishing slurm_spank_exit");
    return (0);
}
int slurm_spank_job_epilog(spank_t sp, int ac, char **av)
{
    slurm_verbose("c: Starting slurm_spank_job_epilog");
    Spank_job_epilog();
    slurm_verbose("c: Finishing slurm_spank_job_epilog");
    return (0);
}
int slurm_spank_slurmd_exit(spank_t sp, int ac, char **av)
{
    slurm_verbose("c: Starting slurm_spank_slurmd_exit");
    Spank_slurmd_exit();
    slurm_verbose("c: Finishing slurm_spank_slurmd_exit");
    return (0);
}
int slurm_spank_init(spank_t sp, int ac, char **av)
{
    slurm_verbose("c: Starting slurm_spank_init");
    int v = Spank_init();
    slurm_verbose("c: Go Spank_init() function returned %d", v);
    slurm_verbose("c: Finishing slurm_spank_init");

    return (0);
}


int main(int argc, char **argv) {
    int v = Spank_init();
    printf("%d\n",v);
    return 0;
}
