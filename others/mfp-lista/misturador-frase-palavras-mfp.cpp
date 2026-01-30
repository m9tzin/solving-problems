/**
* Dada uma string inicial, devemos retornar as duas combinações de palavras que a compõe:
* Exemplo:
* input: hweolrllod  output: hello world
* A ideia é: vamos percorrer o input salvando as letras em posições ímpares => primeira palavra
* e em posições pares => segunda palavra
*/

#include <iostream>
using namespace std;

int main() {
    string s;
    cin >> s;
    string a, b;

    /* push_back: Adds a new element at the end of the vector, after its current last element. The content of val is copied (or moved) to the new element. */
    for(int i=0; i < s.size(); i+=2){
        a.push_back(s[i]);
        b.push_back(s[i+1]);
    }
    cout << a << endl << b << endl;
    return 0;
}