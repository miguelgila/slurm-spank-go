
#include <stdio.h>
#include <stdlib.h>

#include <string.h>
#include <slurm/spank.h>

#include "foo.h"

SPANK_PLUGIN(basic, 1);

int slurm_spank_init(spank_t sp, int ac, char **av)
{
    slurm_verbose("bar: slurm_spank_init() - Start");
    slurm_verbose("bar: Calling go function: Slurm_spank_init()");
    int v = Spank_init();
    slurm_verbose("bar: Go Spank_init() function returned %d", v);
    // Uncommenting the 2 lines above produces this error in slurmd log:
    // [2020-06-04T17:15:17.343] [1034195.extern] bar: slurm_spank_init() - Start
    // [2020-06-04T17:15:17.343] [1034195.extern] bar: Calling go function: Slurm_spank_init()
    // [2020-06-04T17:15:17.344] error: _forkexec_slurmstepd: slurmstepd failed to send return code got 0: No such process
    // [2020-06-04T17:15:17.344] debug3: _spawn_prolog_stepd: return from _forkexec_slurmstepd -1
    // [2020-06-04T17:15:17.345] Could not launch job 1034195 and not able to requeue it, cancelling job
    slurm_verbose("bar: slurm_spank_init() - Done");
    return (0);
}


int main(int argc, char **argv) {
    // This calls a Go function and passes to it a value
    PrintInt(42);

    // This calls a Go function and gets a value from it
    int v = Spank_init();
    printf("%d\n",v);
    return 0;
}
