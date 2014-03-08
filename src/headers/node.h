#ifndef GRAPH_H_INCLUDED
#include "graph.h"
#endif

#ifndef NODE_H_INCLUDED
#define NODE_H_INCLUDED

template<typename T>
class Graph<T>::Node
{
    protected:
        T name;
        arc_s adjacent;

    public:
        Node();
        Node (T);
        T getName();
        void setName (T);
        vector<Arc> getAdjacent();
        node_p copy();
        unsigned int getDegree();
        void setAdjacent (arc_s);
        void addArc (node_p, double);
        void removeArc (node_p);
        bool equals (node_p/*, binary_function<T,T,bool>*/);
        string toString ();
};

#endif