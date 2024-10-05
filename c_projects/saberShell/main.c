#include <stdio.h>
#include <string.h>
#include <stdlib.h>
#include <unistd.h>   // For fork(), execvp()
#include <sys/wait.h> // For wait()
#include <fcntl.h>  // For open()

#define MAX_LENGTH 1024
#define MAX_ARGS 64

int main()
{
    char input[MAX_LENGTH];
    char *args[MAX_ARGS];
    char *token;
    int in_redirect = 0, out_redirect = 0;
    char *in_file;
    char *out_file;
    while (1)
    {
        printf("SaberShell> ");
        if (fgets(input, MAX_LENGTH, stdin) == NULL)
        {
            // Handle Ctrl+D to exit (EOF)
            printf("Bye Bye\n");

            break;
        }
        // Remove the newline char from the input
        input[strcspn(input, "\n")] = '\0';

        // If exit is typed close the shell
        if (strcmp(input, "exit") == 0)
        {
            printf("Bye Bye\n");
            break;
        }

        int i = 0;
        // strtok returns the words of a string separated by the deliminator, so here it will returns words separatod by space.
        // Each time it is called will return the next word.
        token = strtok(input, " ");
        while (token != NULL && i < MAX_ARGS - 1)
        {
            if (strcmp(token, "<") == 0)
            {
                in_redirect = 1;
                token = strtok(NULL, " "); // Get the input file
                if (token != NULL)
                {
                    in_file = token;
                }
            }
            else if (strcmp(token, ">") == 0)
            {
                out_redirect = 1;
                token = strtok(NULL, " "); // Get the output file
                if (token != NULL)
                {
                    out_file = token;
                }
            }
            else
            {
                args[i++] = token;
            }
            token = strtok(NULL, " ");
        }
        args[i] = NULL; // Null-terminate the array

        // Built-in command: cd
        if (strcmp(args[0], "cd") == 0)
        {
            if (args[1] == NULL)
            {
                fprintf(stderr, "saberShell: expected argument to \"cd\"\n");
            }
            else
            {
                if (chdir(args[1]) != 0)
                {
                    perror("saberShell");
                }
            }
            continue;
        }
        
        int pid = fork();

        if (pid == -1)
        {
            perror("Fork Failed");
        }
        else if (pid == 0)
        {
            if (in_redirect && in_file)
            {
                int fd_in = open(in_file, O_RDONLY);
                if (fd_in < 0) {
                    perror("saberShell: cannot open input file");
                    exit(EXIT_FAILURE);
                }
                dup2(fd_in, STDIN_FILENO); // Redirect stdin to file
                close(fd_in);
            }

            if (out_redirect && out_file)
            {
                int fd_out = open(out_file, O_WRONLY | O_CREAT | O_TRUNC, 0644);
                if (fd_out < 0) {
                    perror("saberShell: cannot open output file");
                    exit(EXIT_FAILURE);
                }
                dup2(fd_out, STDOUT_FILENO); // Redirect stdout to file
                close(fd_out);
            }

            if (execvp(args[0], args) == -1)
            {
                perror("exec failed");
            }
            exit(EXIT_FAILURE); // Exit if exec fails
        }
        else
        {
            // Parent process: Wait for the child to finish
            wait(NULL);
        }
        in_redirect = 0;
        out_redirect = 0;
        in_file = NULL;
        out_file = NULL;
    }
    return 0;
}
