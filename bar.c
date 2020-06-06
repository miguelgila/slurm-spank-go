
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
