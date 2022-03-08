#include <stdio.h>
#include <stdlib.h>

typedef struct Tree {
    int tree_array[100];
    int last_index;
} Tree;

void swap(int *a, int *b)
{
    printf("swap : \t%d\t%d\n", *a, *b);
    int temp = *b;
    *b = *a;
    *a = temp;
}

void Initialize(Tree *tree) {
    (*tree) = (Tree){
        .tree_array = {},
        .last_index = 0
    };
}

void Setdata(Tree *tree, int arr[], int len){
    int i;
    for (i = 0; i < len; i++)
        (*tree).tree_array[i] = arr[i];
    (*tree).last_index = len;
    for (i = 0; i < (*tree).last_index; i++)
        printf("%d ", (*tree).tree_array[i]);
    printf("\n");
}

void Showdata(Tree *tree){
    int i;
    for (i = 0; i < (*tree).last_index; i++)
        printf("%d ", (*tree).tree_array[i]);
    printf("\n");
}

void Adddata(Tree *tree, int new){
    printf("Add New data :\t%d\n", new);
    (*tree).tree_array[(*tree).last_index++] = new;
}

void Deletedata(Tree *tree, int del){
    printf("Delete data :\t%d\n", del);
    int i;
    for (i = 0; i < (*tree).last_index; i++){
        if ((*tree).tree_array[i] == del){
            swap(&(*tree).tree_array[i], &(*tree).tree_array[--(*tree).last_index]);
            (*tree).tree_array[(*tree).last_index] = 0;
        }
    }
}

void Max_heapify(int arr[], int start, int end)
{
    // Create data node and son node
    int dad = start;
    int son = dad * 2 + 1;
    while (son <= end)
    { // If son node in tree range
        // Compare two child node, if left is greater than right, choose left
        if (son + 1 <= end && arr[son] < arr[son + 1])
            son++;
        if (arr[dad] > arr[son]) // IF dad node greate son node adjusted, end func
            return;
        else
        { // Else swap dad and son, and compare grandson
            swap(&arr[dad], &arr[son]);
            dad = son;
            son = dad * 2 + 1;
        }
    }
}

void Min_heapify(int arr[], int start, int end)
{
    // Create data node and son node
    int dad = start;
    int son = dad * 2 + 1;
    while (son <= end)
    { // If son node in tree range
        // Compare two child node, if left is greater than right, choose left
        if (son + 1 <= end && arr[son] > arr[son + 1])
            son++;
        if (arr[dad] < arr[son]) // IF dad node greate son node adjusted, end func
            return;
        else
        { // Else swap dad and son, and compare grandson
            swap(&arr[dad], &arr[son]);
            dad = son;
            son = dad * 2 + 1;
        }
    }
}

int Checklevel(int index){
    int count = 0;
    while(index != 1){
        index /= 2;
        count++;
    }
    return count + 1;
}

// void Pushdown_Min(int arr[], int index, int end){
//     int dad = index;
//     int son = dad * 2 + 1;
//     while (son <= end)
//     { // If son node in tree range
//         // Compare two child node, if left is greater than right, choose left
//         if (son + 1 <= end && arr[son] > arr[son + 1])
//             son++;
//         if (arr[dad] < arr[son]) // IF dad node greate son node adjusted, end func
//             return;
//         else
//         { // Else swap dad and son, and compare grandson
//             swap(&arr[dad], &arr[son]);
//             dad = son;
//             son = dad * 2 + 1;
//         }
//     }
// }

void Pushdown_Min(int arr[], int index, int end){
    int dad = index;
    int son = dad * 2 + 1;
    while (son <= end)
    { // If son node in tree range
        // Compare two child node, if left is greater than right, choose left
        int gson = son * 2;
        if (son * 2 <= end){
            if (arr[gson] < arr[dad]){
                swap(&arr[gson], &arr[dad]);
                if (arr[gson] > arr[son]){
                    swap(&arr[gson], &arr[son]);
                }
                Pushdown_Min(arr, gson, end);
            }
        } else if(arr[son] < arr[index]){
            swap(&arr[son], &arr[dad]);   
        }
    }
}

void Pushdown_Max(int arr[], int index, int end){
    int dad = index;
    int son = dad * 2 + 1;
    while (son <= end)
    { // If son node in tree range
        // Compare two child node, if left is greater than right, choose left
        int gson = son * 2;
        if (son * 2 <= end){
            if (arr[gson] > arr[dad]){
                swap(&arr[gson], &arr[dad]);
                if (arr[gson] < arr[son]){
                    swap(&arr[gson], &arr[son]);
                }
                Pushdown_Max(arr, gson, end);
            }
        } else if(arr[son] > arr[index]){
            swap(&arr[son], &arr[dad]);   
        }
    }
}

// void Pushdown_Max(Tree *tree, int index){
//     if ((index)*2 <= (*tree).last_index + 1){
//     }
// }

void Pushdown(Tree *tree, int index){
    printf("Level %d\n", Checklevel(index));
    if ((Checklevel(index) % 2) != 0) { //Min levle
        Pushdown_Min((*tree).tree_array, index, (*tree).last_index);
    } else {
        Pushdown_Max((*tree).tree_array, index, (*tree).last_index); 
    }
}

void Maxheap(Tree *tree)
{
    int i;
    for (i = (*tree).last_index / 2 - 1; i >= 0; i--)
        Max_heapify((*tree).tree_array, i, (*tree).last_index - 1);
        Showdata(tree);
}

void Minheap(Tree *tree)
{
    int i;
    for (i = (*tree).last_index / 2 - 1; i >= 0; i--)
        Min_heapify((*tree).tree_array, i, (*tree).last_index - 1);
        Showdata(tree);
}

void Min_max_heap(Tree *tree){
    int i;
    for (i = (*tree).last_index / 2 - 1; i >= 0; i--){
        printf("Donwback %d\t%d\n", (*tree).tree_array[i], i);
        Pushdown(tree, i + 1);
    }
    Showdata(tree);
}

int main()
{
    Tree data;
    Initialize(&data);

    int ex1data[] = {27, 7, 80, 5, 67, 18, 62, 24, 58, 25};
    int ex2data[] = {40, 23, 10, 15, 8};
    int ex3data[] = {20, 30, 10, 50, 60, 40, 45, 5, 15, 25};
    int ex4data[] = {15, 10, 8, 22, 12, 6, 18};
    int ex4_2data[] = {20, 8, 28, 10, 4, 5, 40, 55};

    //Ex1 bulid maxheap
    int len = sizeof ex1data / sizeof *ex1data;
    Initialize(&data);

    printf("Ch8.5 Ex1. Maxheapify\n");
    Setdata(&data, ex1data, len); //Set data in tree
    Maxheap(&data);

    //Ex2 adddata in maxheap
    len = sizeof ex2data / sizeof *ex2data;
    Initialize(&data);
    
    printf("\nCh8.5 Ex2. Add and delete\n");
    Setdata(&data, ex2data, len);

    //(a) Add 60 and 20
    Adddata(&data, 60);
    Maxheap(&data);
    Adddata(&data, 20);
    Maxheap(&data);
    //(b) Delete 60 and 23
    Deletedata(&data, 60);
    Maxheap(&data);
    Deletedata(&data, 23);
    Maxheap(&data);

    //Ex3 bulid minheap
    printf("\nCh8.5 Ex3. Build Minheapify\n");
    len = sizeof ex3data / sizeof *ex3data;
    Initialize(&data);
    Setdata(&data, ex3data, len);
    Minheap(&data);

    //Ex5 min-max-heapify
    printf("\nCh8.5 Ex4. Build Min-maxheap\n");
    len = sizeof ex4_2data / sizeof *ex4_2data;
    Initialize(&data);
    Setdata(&data, ex4_2data, len);

    Min_max_heap(&data);
    return;
}