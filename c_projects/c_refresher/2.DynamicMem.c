#include <stdio.h>
#include <string.h>
#include <stdlib.h>
/*
    malloc(size): Allocates memory of 'size' bytes and returns a pointer to it
    calloc(num, size): Allocate memory for an array of 'num' elems with each elem being of 'size' bytes, initializes the mem allocs to 0
    realloc(*ptr, size): Resizes previously allocated mem block to 'size' bytes
    free(*ptr): Frees mem previously allocated by malloc, calloc or realloc
 */
int main()
{
     char *str = (char *)malloc(50 * sizeof(char));
     if (str == NULL){
        printf("Memory Allocation Failed!!!!");
        return -1;
     }
     strcpy(str, "Hello Dynamic Memory!!");
     printf("%s\n", str);
     free(str);
}