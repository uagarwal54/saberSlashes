#include <stdio.h>
#include <string.h>
#include <stdlib.h>

int main(){
    int size = 10;
    int index = 0;
    char *str = (char *)malloc(size * sizeof(char));
    if (str == NULL){
        printf("Memory Allocation Failed!!!!");
        return -1;
    }
    char c = getchar();
    while(c != '\n'){
        str[index] = c;
        c = getchar();
        index++;
        if (index >= size){
            size = size * 2;
            realloc(str, size * sizeof(char));
            if (str == NULL) {
                printf("Memory reallocation failed\n");
                return -1;
            }
        }
    }
    printf("You entered: %s\n", str);
    
}