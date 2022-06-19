#include <iostream>
#include <fstream>
#include <algorithm>
#include <map>

using namespace std;

char N[15] = {'_', 'A', '2', '3', '4', '5', '6', '7', '8', '9', 'T', 'J', 'Q', 'K', 'A'};
ofstream file;
int value;
std::map<int, std::string> thresholds;

string key(int a, int b, int c, int d, int e){
    char cards[5] = {N[a], N[b], N[c], N[d], N[e]};
    sort(cards, cards+5);
    string key = "";
    for(int card : cards)
        key += card;
    
    return key;
}

void highCard(){
    // High Card
    // No duplicates, no straights, sorted from 12346 to 8TJQK
    for(int a=6;a<=14;a++)
        for(int b=2;b<a;b++)
            for(int c=b+1;c<a;c++)
                for(int d=c+1;d<a;d++)
                    for(int e=d+1;e<a;e++)
                        if((c != b+1) || (d != c+1) || (e != d+1) || (a != e+1))
                            file << key(a,b,c,d,e) << ' ' << value++ << ' ';
    thresholds[value] = "High Card";
}

void onePair(){
    // Single Pair
    // Only a single pair, sorted by kicker cards
    for(int p=2;p<=14;p++){
        for(int x=4;x<=14;x++){
            for(int y=3;y<x;y++){
                for(int z=2;z<y;z++){
                    if(x!=p && y!=p && z != p){
                        file << key(p, p, x, y, z) << ' ' << value++ << ' ';
                    }
                }
            }
        }
    }
    thresholds[value] = "High Card";
}

void twoPair(){
    // Two Pair
    // Two pairs and sorted by value, then value of kicker card
    for(int p1=3;p1<=14;p1++){
        for(int p2=2;p2<p1;p2++){
            for(int k=2;k<=14;k++){
                if(k != p1 && k != p2){
                    file << key(p1, p1, p2, p2, k) << ' ' << value++ << ' ';
                }
            }
        }
    }
    thresholds[value] = "Two Pair";
}

void threes(){
    // Three of a Kind
    // Sorted by triple's value, then by kicker cards (cannot be full house)
    for(int t=2;t<=14;t++){
        //generate 2 card combos
        for(int a=2;a<=14;a++){
            for(int b=2;b<a;b++){
                if(a != t && b != t)
                    file << key(a, b, t, t, t) << ' ' << value++ << ' ';
            }
        }
    }
    thresholds[value] = "Three of Kind";
}

void straight(){
    for(int i=1;i<=10;i++){
        file << key(i, i+1, i+2, i+3, i+4) << ' ' << value++ << ' ';
    }
    thresholds[value] = "Straight";
}

void flush(){
    // Flush
    // No duplicates, no straights, sorted from 12346 to 8TJQK
    int start = value;
    for(int a=6;a<=14;a++)
        for(int b=2;b<a;b++)
            for(int c=b+1;c<a;c++)
                for(int d=c+1;d<a;d++)
                    for(int e=d+1;e<a;e++)
                        if((c != b+1) || (d != c+1) || (e != d+1) || (a != e+1))
                            file << key(a,b,c,d,e) << "f " << value++ << ' ';
    thresholds[value] = "Flush";
}

void fullHouse(){
    for(int a=2;a<=14;a++){
        for(int b=2;b<=14;b++){
            if(a != b)
                file << key(b, b, a, a, a) << ' ' << value++ << ' ';
        }
    }
    thresholds[value] = "Full House";
}

void fours(){
    // Four of a Kind
    // Sorted by fours value, then by kicker card
    for(int f=2;f<=14;f++){
        for(int k=2;k<=14;k++){
            if(f != k)
                file << key(f, f, f, f, k) << ' ' << value++ << ' ';
        }
    }
    thresholds[value] = "Four of a Kind";
}

void straightFlush(){
    for(int i=1;i<=10;i++){
        file << key(i, i+1, i+2, i+3, i+4) << "f " << value++ << ' ';
    }
    thresholds[value] = "Straight Flush";
}

int main(){
    file.open("store.txt");
    
    value = 1;

    highCard();
    onePair();
    twoPair();
    threes();
    straight();
    flush();
    fullHouse();
    fours();
    straightFlush();

    file << std::endl;

    for(auto & p : thresholds) {
      file << p.first << " " << p.second << std::endl;      
    }

    file.close();
}
