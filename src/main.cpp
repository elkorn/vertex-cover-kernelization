#include <iostream>
#include <functional>
#include <string>
#include <vector>
#include "graph.cpp"

using namespace std;

void printNodes (vector<typename Graph<int>::node_p> nodeMap);

class IntComparer: public binary_function<int, int, bool>
{
    public:
        bool operator() (int a, int b) const
        {
            cout << "test" << endl;
            return a == b;
        }
};

int main()
{
    string name;
    int num = 10;
    IntComparer sc;
    cout << sc(1,2) << " " << sc(1,1) << endl;
    Graph<int> g;
    cout << "Just a simple example... " << endl;

    for (int i = 0; i < num; i++) {
        g.addNode (num-i);
    }

    g.addArc(5, 4);
    g.addArc(2, 7);

    cout << endl << endl << "Nodes before removal..." << endl;
    printNodes (g.getNodes());
    g.removeNode (5);
    cout << endl << endl << "Nodes after removal..." << endl;
    printNodes (g.getNodes());
    system ("PAUSE");
}

void printNodes (vector<typename Graph<int>::node_p> map)
{
    for (unsigned int i = 0; i < map.size(); i++) {
        cout << map.at (i)->toString() << endl;
    }
}
