#include <iostream>
#include <functional>
#include <string>
#include <vector>
#include <cassert>
#include "../src/vertex-cover-solver-buss.cpp"

using namespace std;

int main()
{
    Graph g;
    vector<int> shouldCover;
    vector<int> shouldNotCover;
    shouldCover.push_back (0);
    shouldCover.push_back (1);
    shouldNotCover.push_back (0);
    g.load ("../data/simple_3x3.txt");
    assert (g.isVertexCover (shouldCover) == true);
    assert (g.isVertexCover (shouldNotCover) == false);
    return 0;
}