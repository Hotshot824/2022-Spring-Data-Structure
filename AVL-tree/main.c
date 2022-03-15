#include<stdio.h>
#include<stdlib.h>
#include<string.h>

struct student {
    char name[20];
    int score;
    int bf;
    struct student *llink, *rlink;
};

struct student *ptr, *root, *this_n, *prev, *pivot, *pivot_prev;
int nodecount = 0;

void flushBuffer(void){
    while(getchar() != '\n'){
        continue;
    }
}

void insert_f(void){
    char name_t[20];
    int score_t;
    printf("\n Please enter name : ");
    scanf("%s", name_t);
    flushBuffer();
    printf(" Please enter score : ");
    scanf("%d", score_t);
    flushBuffer();
    nodecount++;
    sort_f(name_t, score_t);
}

void sort_f(char name_t[], int score_t){
    int op;
    this_n = root;
    while((this_n != NULL) && (strcmp(name_t, this_n -> name) != 0)){
        if (strcmp(name_t, this_n -> name) < 0){
            prev = this_n;
            this_n = this_n -> llink;
        }
        else{
            prev = this_n;
            this_n = this_n -> rlink;
        }
    }

    if(this_n == NULL || strcmp(name_t, this_n -> name) != 0){
        ptr = (struct student *)malloc(sizeof(struct student));
        strcpy(ptr -> name, name_t);
        ptr -> score = score_t;
        ptr -> llink = NULL;
        ptr -> rlink = NULL;
        if (root == NULL)
            root = ptr;
        if (prev != NULL){
            if(strcmp(ptr -> name, prev -> name) < 0)
                prev -> llink = ptr;
            else
                prev -> rlink = ptr;
        }
        bf_count(root);
        pivot_find();

        if (pivot != NULL) {
            op = type_find();
            switch (op){
            case 11:
                type_ll();
                break;
            case 22:
                type_rr();
                break;
            case 12:
                type_lr();
                break;
            case 21:
                type_rl();
                break;  
            }
        }
        bf_count(root);
    }
    else {
        printf(" %s No exist!\n", name_t);
    }
}

int main(){
    char option;
    while (1)
    {
        printf("    <1> insert\n");
        printf("    <2> delete\n");
        printf("    <3> modify\n");
        printf("    <4> list\n");
        printf("    <5> exit\n");
        printf("**********************\n");
        printf("    Enter choice : \n");
        option = getchar();
        flushBuffer();
        switch(option) {
            case '1':
                insert_f();
                break;
            case '2':
                delete_f();
                break;
            case '3':
                modify_f();
                break;
            case '4':
                list_f();
                break;
            case '5':
                exit(0);
        }
    }

    return;
}