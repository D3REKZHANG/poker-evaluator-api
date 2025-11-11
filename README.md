# poker-evaluator-api

Blazing fast poker hand evaluator API written in Go and C++.

My take on a lookup table poker hand evaluation algorithm.

Can process 100,000 evaluations for 5-card poker hands in **0.019s**

Uses the Echo framework for the rest API.


## Implementation

My approach involves precomputing a hashtable where the keys are all the possible 5 card combinations from a standard deck of playing cards.

My twist is that instead of the 52C5 = **2,598,960** combinations, I only care about the possible card value combinations, regardless of suit. I then add the possible flush combinations, adding an 'f' to these keys. This reduces it down to a total of **7464** possible combinations (explained in math section).

This makes finding the strength of a given 5 card hand very simple, as all I have to do is quickly check if the suits are all equal (flush), and then sort the values to get the hashtable key to lookup the corresponding strength in average **O(1)** time.

## Math

To find the number of ways of picking 5 cards from the 13 card values where suit doesn't matter:

This problem is equivalent to picking 5 items from 13 piles of 4 indistinguishable items, so its value can be obtained from the following generating function:

![equation](https://latex.codecogs.com/png.latex?f%28x%29%20%3D%20%28x%5E0&plus;x%5E1&plus;x%5E2&plus;x%5E3&plus;x%5E4%29%5E%7B13%7D%20%3D%20%281&plus;x&plus;x%5E2&plus;x%5E3&plus;x%5E4%29%5E%7B13%7D)

This is because for each card value (1-13), you can pick 0 of them, 1 of them, ... up to 4 of them. This is represented by the exponents in the generating function.
Since there are 13 different piles (values), then we multiply this polynomial 13 times which will generate all the different possibilities to take the different items (cards).

After expanding this monstrous expression ([Wolfram](https://www.wolframalpha.com/input/?i=%281%2Bx%2Bx%5E2%2Bx%5E3%2Bx%5E4%29%5E13)), taking the coefficient before the `x^5` term gives us the total number of ways to pick 5 items (cards). In this case, the value is `6175`. We need to add 1 since `A` can also be treated as a 1 in straights.

For flushes and straight flushes, there can't be any duplicates, so the possible combinations for flushes is simply `13C5 = 1287`. We must also add 1 for the same reason as above.

Thus we have our total:

![equation](https://latex.codecogs.com/png.latex?6175&plus;1&plus;1287&plus;1%3D7464)

Turns out I ended up recalling something from AP Stats in high school lol. Thanks Mr. Segev!



