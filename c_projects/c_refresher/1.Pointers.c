#include <stdio.h>

/*
Pointers in C hold addresses of some memory allocated to store some data.
There are 2 symbols assosiated with them
    1. & := The 'address of' operator fetches the address of the variable
    2. * := The 'defrence operator fewtches the actual value stored in some address stored in a pointer '
 */
int main()
{
    int x = 10;
    int *ptr1 = &x;
    printf("%d\n", *ptr1);
    printf("==========================\n");
    /*
    Strings in C are array of chars with the name of the array being the pointer to the first elem itself.
     */

    char name[] = "Udbhav Agarwal";
    char *ptr = name; // Now ptr has the addres of U, the first char of the array name
    printf("%c\n", *ptr); // This will print the Char U
    printf("==========================\n");
    // This is how pointer arithematic works, we can just increment the pointer and get the next elem
    // in case of arrays
    while (*ptr != '\0'){
        printf("%c\n",*ptr);
        ptr++;
    }

    
}