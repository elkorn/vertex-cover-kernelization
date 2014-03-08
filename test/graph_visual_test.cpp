#include <iostream>
#include <functional>
#include <string>
#include <vector>
#include "../src/graph.cpp"
#include "../src/test.cpp"
#include "../src/graphPrinter.cpp"

using namespace std;

int main()
{
    string name;
    int num = 10;
    float test = 66.6;
    Graph<float> g, otherG;
    Graph<float>::node_p sharedNode = g.makeNode (test);
    cout << "Just a simple example... " << endl;

    for (int i = 0; i < num; i++) {
        g.addNode ( (float) num - i + .01);
    }

    g.addArc (5.01, 4.01);
    g.addArc (2.01, 7.01);
    g.addArc (3.01, 7.01);
    cout << sharedNode->getName() << endl;
    g.addNode (sharedNode);
    cout << "Added shared node" << endl;
    GraphPrinter::printGraph (g, "g");
    otherG.addNode (sharedNode);
    cout << "Added shared node to otherG" << endl;
    GraphPrinter::printGraph (otherG, "otherG");
    cout << endl << endl << "Nodes before removal..." << endl;
    g.removeNode (5);
    cout << endl << endl << "Nodes after removal..." << endl;
    GraphPrinter::printGraph (g, "g");

    delete sharedNode;
    sharedNode = NULL;
}
