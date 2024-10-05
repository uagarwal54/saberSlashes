#include <string.h>
#include <stdio.h>

int main(){
    char name[] = "udbhav is a platform engineer";
    char *token = strtok(name, " ");
    while(token != NULL){
        printf("%s\n", token);
        token = strtok(NULL, " ");
    }
    return 0;
}

