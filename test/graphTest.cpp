#include <iostream>
#include <cassert>
#include "../src/Graph.cpp"

using namespace std;

int main(int argc, char const *argv[])
{

    Graph<int> g1;
    g1.insert(1);
    assert(g1.size() == 1);
    g1.insert(1);
    g1.insert(1);
    assert(g1.size() == 3);

    Graph<string> g2;
    g2.insert("one");
    assert(g2.size() == 1);
    g2.insert("two");
    g2.insert("three");
    assert(g2.size() == 3);

    g1.display(cout);
    g2.display(cout);

    cout << "[" << argv[0] << "] All tests ok!" << endl;
    return 0;
}