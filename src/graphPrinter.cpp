#include <iostream>
#include "headers/graph.h"

class GraphPrinter
{
    public:
        template <typename T>
        static void printGraph (const Graph<T> &graph, string name)
        {
            typename Graph<T>::node_s map = graph.getNodes();
            cout << endl << "=====" << endl;
            cout << "  " << name << endl;
            cout << "-----" << endl;

            for (typename Graph<T>::node_it it = map.begin(),
                 end = map.end();
                 it != end;
                 ++it) {
                cout << (*it)->toString() << endl;
            }

            cout << "=====" << endl << endl;
        }
};