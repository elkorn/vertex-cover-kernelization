#include <iostream>
#include <cassert>
#include "../src/Graph.cpp"

using namespace std;

void title(const string title) {
    cout << " *** " << title << " *** " << endl;
}

int main(int argc, char const *argv[])
{
    title(string(argv[0]));
    title("BASIC API");
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
    g2.connect(2, 1);

    g1.display(cout);
    g2.display(cout);

    title("VERTEX COVER");

    Graph<int>::Node* sharedNode = new Graph<int>::Node(13);
    g1.insertNode(sharedNode);
    Graph<int> g3;
    g3.insertNode(sharedNode);

    assert(g3.isVertexCover(g1) == true);
    delete sharedNode;
    title("All tests ok!");
    return 0;
}