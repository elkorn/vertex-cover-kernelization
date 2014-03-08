#ifndef GRAPH_H_INCLUDED
#include "graph.h"
#endif

#ifndef ARC_H_INCLUDED
#define ARC_H_INCLUDED

class Graph<T>::Arc
{
    protected:
        node_p head;
        node_p tail;
        double weight;

    public:
        Arc();
        Arc (node_p, node_p, double);
        Arc (node_p, node_p);
        node_p getHead();
        node_p getTail();
        double getWeight();
        void setHead (node_p);
        void setTail (node_p);
        void setWeight (double);
        Arc copy();
        string toString();
        bool equals (Arc);
        bool isCoveredBy (node_p);
};

#endif