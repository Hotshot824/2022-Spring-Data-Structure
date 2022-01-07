#include<stdlib.h>
#include<stdio.h>

typedef struct node {
	int data;
	struct node *next;
	struct node *prev;
} Node;

Node* createNode(int value){
	Node *new_node = (Node*)malloc(sizeof(Node));
	new_node -> data = value;
	new_node -> next = NULL;
	new_node -> prev = NULL;
	return new_node;
}

void addNode(Node **start, int value) {
	Node* new_node;
	new_node = createNode(value);
	
	if(*start == NULL) {
		*start = new_node;
		return;
	}	else {
		Node *current;
		current = *start;
		while(current -> next != NULL) {
			current = current -> next;
		}
		current -> next = new_node;
		new_node -> prev = current;
	}
}

void insert_Node(Node **start, int Index, int Value) {
	Node *current;
	Node *temp;
	Node *new_node = (Node*)malloc(sizeof(Node));
	new_node -> data = Value;
	new_node -> next = NULL;
	new_node -> prev = NULL;
	if(Index == 0) {
		temp = *start;
		temp -> prev = new_node;
		new_node -> next = temp;
		*start = new_node;
		return;
	}
	int i = 0;
	current = *start;
	while(current != NULL) {
		if(i == Index - 1) {
			temp = current -> next;
			temp -> prev = new_node;
			current -> next = new_node;
			new_node -> next = temp;
			new_node -> prev = current;
			return;
		}
		current = current -> next;
		++i;
	}
	printf("Error!\n");
}

void insert_value_Node(Node **start, int Value) {
	Node *current;
	Node *temp;
	Node *new_node = (Node*)malloc(sizeof(Node));
	new_node -> data = Value;
	new_node -> next = NULL;
	new_node -> prev = NULL;

	current = *start;
	if(Value < current -> data) {
		current -> prev = new_node;
		*start = new_node;
		new_node -> next = current;
		return;
	}
	while(current != NULL) {
		if(Value < current ->  data) {
			current -> prev -> next = new_node;
			new_node -> prev = current -> prev;
			current -> prev = new_node;
			new_node -> next = current;

			/*
			temp = current -> next;
			temp -> prev = new_node;
			current -> next = new_node;
			new_node -> next = temp;
			new_node -> prev = current;*/

			return;
		}
		if(current -> next == NULL) { //Last node
			current -> next = new_node;
			new_node -> prev = current;
			return;
		}
		current = current -> next;
	}
}

void Printlist(Node *node) {
	while(node != NULL) {
		printf("%d ", node -> data);
		node = node -> next;
	}
	printf("\n");
}

void Inv_Printlist(Node *node) {
	while(node -> next != NULL) {
		node = node -> next;
	}
	while(node != NULL) {
		printf("%d ", node -> data);
		node = node -> prev;
	}
	printf("\n");
}

void DeleteValue(Node **start, int Value) {
	Node *current;
	Node *temp;
	current = *start;
	if(*start == NULL) {
		printf("Delete Error! List is Null.\n");
	}
	if(current -> data == Value) {
		*start = current -> next;
		temp = *start;
		temp -> prev = NULL;
		free(current);
		return;
	}
	while(current != NULL) {
		if(current -> next -> data == Value) {
			temp = current -> next;
			current -> next = current -> next -> next;
			current -> next -> prev = current;
			free(temp);
			return;
		}
		current = current -> next;
	}
}

int main() {
	Node *head;
	int chioce, data, Round = 1;
	printf("1.Input data.\n2.Print data\n3.Print Inverse data\n4.DeleteValue.\n5.Exit\n6.Insert by Index\n7.Insert by Value\n");
	while(true) {
		printf("chioce? \n");
		scanf("%d", &chioce);
		switch(chioce) {
			case 1:
				printf("Input data, ctrl+z Exit.\n");
				while(printf("Round %d Input:", Round) && scanf("%d", &data) != EOF) {
					addNode(&head, data);
					Round++;
				}
				break;
			case 2:
				printf("Print all data: \n");
				Printlist(head);
				break;
			case 3:
				printf("Print Inverse all data: \n");
				Inv_Printlist(head);
				break;
			case 4:
				printf("Please Input want Delete Value!\n");
				scanf("%d", &data);
				DeleteValue(&head, data);
				break;
			case 5:
				printf("Bye!\n");
				exit(0);
			case 6:
				printf("Insert by Index : ");
				scanf("%d", &chioce); //index
				scanf("%d", &data);
				insert_Node(&head, chioce, data);
			case 7:
				printf("Insert by value : ");
				scanf("%d", &data);
				insert_value_Node(&head, data);
			default:
				break;
		}
	}
	return 0;
}
