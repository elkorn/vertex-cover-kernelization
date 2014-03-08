#include <iostream>
#include <functional>
#include <string>
#include <vector>
#include "../src/graph.cpp"
#include "../src/test.cpp"

using namespace std;

int main()
{
    Test::title ("Vertex Cover");
    string name;
    int num = 10;
    float test = 66.6;
    Graph<float> g, otherG;
    Graph<float>::node_p sharedNode = g.makeNode (test);

    for (int i = 0; i < num; i++)
    {
        g.addNode ( (float) num - i + .01);
    }

    g.addArc (5.01, 4.01);
    g.addArc (2.01, 7.01);
    g.addArc (3.01, 7.01);
    g.addNode (sharedNode);
    otherG.addNode (sharedNode);

    Test::test (!otherG.isVertexCover (g),
                "Independent sets cannot be vertex covers of each other.");
    Test::test (!g.isVertexCover (otherG),
                "Independent sets cannot be vertex covers of each other.");

    
    delete sharedNode;
    sharedNode = NULL;
}