#include <iostream>
#include <functional>
#include <string>
#include <vector>
#include <cassert>
#include "../src/vertex-cover-solver-buss.cpp"

using namespace std;

int main()
{
    Graph g = Graph::fromFile ("../data/simple_3x3.txt");
    SolverBuss solver (g);
    vector<int> shouldCover;
    vector<int> shouldNotCover;
    shouldCover.push_back (0);
    shouldCover.push_back (1);
    shouldNotCover.push_back (0);
    assert (solver.isVertexCover (shouldCover) == true);
    assert (solver.isVertexCover (shouldNotCover) == false);
    // Can detect when a graph has no cover.
    assert (solver.findCover (1).size() == 0);
    return 0;
}