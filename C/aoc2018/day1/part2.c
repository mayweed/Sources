#include<stdio.h>
#include<stdlib.h>

//to test the link, how about ring array (index to 0 when max array etc...cant crush
//data ceci dit...)
typedef struct node{
    int value;
    struct node *next;
} node;

int main(void){

    FILE *input=fopen("input.txt","r");
    int total=0;
    int change=0;

    struct node *numbers=NULL;

    //read input file only once, if the redundant freq is not found
    //you read the freq in numbers value (you iterate the list and add the value at
    //the end) til you found the ONE
    if(input){
        while(fscanf(input,"%i",&change) == 1){
            total+=change;
            //fprintf(stderr,"%d\n",total);

            //init an elt
            node *n = (node*) malloc(sizeof(node));
            n -> value=total;
            n -> next=NULL;
            //fprintf(stderr,"%i\n",n->value);

            if(numbers){
                //add an elt to the list
                for (node *ptr=numbers; ptr!=NULL;ptr=ptr->next){
                    if (n->value == ptr -> value){
                        printf("find the magic word %d\n",ptr->value);
                        break;
                    }
                    if (!ptr->next){
                        ptr->next=n;
                        break;
                        
                    }
                    //fprintf(stderr,"%d %d\n",numbers->value,numbers->next->value);
                }
            }
            //list is empty, first element
            else{
                numbers=n;
            }
        }
        fclose(input);
    }

    free(numbers);
    printf("%d\n",total);
    return 0;
}
