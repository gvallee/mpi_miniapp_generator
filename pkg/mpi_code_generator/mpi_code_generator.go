//
// Copyright (c) 2022, NVIDIA CORPORATION. All rights reserved.
//
// See LICENSE.txt for license information
//

package mpi_code_generator

import "os"

func AddIncludes(fd *os.File) {
	fd.WriteString("#include <assert.h>\n")
	fd.WriteString("#include <stdlib.h>\n")
	fd.WriteString("#include <unistd.h>\n")
	fd.WriteString("#include <string.h>\n")
	fd.WriteString("#include <stdio.h>\n")
	fd.WriteString("#include \"mpi.h\"\n\n")
}

func AddGenericMain(fd *os.File) {
	fd.WriteString("int main(int argc, char **argv) {\n")
	fd.WriteString("\tint rank, size;\n")
}

func AddGenericInit(fd *os.File) {
	fd.WriteString("\tMPI_Init(&argc, &argv);\n")
	fd.WriteString("\tMPI_Comm_rank(MPI_COMM_WORLD, &rank);\n")
	fd.WriteString("\tMPI_Comm_size(MPI_COMM_WORLD, &size);\n")
}

func AddGenericFinalize(fd *os.File) {
	fd.WriteString("\tMPI_Finalize();\n")
}
