/* construire une pyramide:
 * dans une boucle qui prend en compte la hauteur de la pyramide
 * partir d'un nombre de blancs donnes (5)
 * ajouter une etoile
 * ajouter un \n
 * decrementer les blancs (4)
 * incrementer l'etoile (ce qui fait 2)
 */

int pyramide (int hauteur)
{
    int i, j, blancs, stars;

    blancs = hauteur;
    stars = 1;

    for (i = 0; i < hauteur; i++)
        {
            for (j=0; j < blancs; j++)
                putchar (' ');
            for (j=0; j < stars; j++)
                putchar ('*');
        blancs --;
        stars +=2;
        putchar ('\n');
        }
}

int main ()
{
    pyramide (5);
}
