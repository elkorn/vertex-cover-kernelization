#ifndef GRAPH_H_INCLUDED
#define GRAPH_H_INCLUDED

#include <vector>
#include <string>
#include <functional>
#include <sstream>

#define ARC "-> "

using namespace std;

template <typename T>
class Graph
{
    public:
        class Node;
        class Arc;

        typedef Node *node_p;
        typedef vector<node_p> node_s;
        typedef vector<Arc> arc_s;
        typedef typename arc_s::iterator arc_it;
        typedef typename node_s::iterator node_it;

        class Arc
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

        class Node
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
                void setAdjacent (vector<Arc>);
                void addArc (node_p, double);
                void removeArc (node_p);
                bool equals (node_p/*, binary_function<T,T,bool>*/);
                string toString ();
        };


        Graph(/*binary_function<T,T,bool>*/);
        ~Graph();
        int getNuNodes();
        int getNuArcs();
        int getVersion();
        vector<node_p> getNodeMap();
        node_p getNode (T);
        void addNode (node_p);
        void addNode (T);
        void removeNode (node_p);
        void removeNode (T);
        vector<node_p> getNodes();
        void addArc (node_p , node_p , double);
        void addArc (T, T, double);
        void addArc (node_p , node_p);
        void addArc (T, T);
        void removeArc (node_p, node_p);
        void removeArc (T, T);
        // void removeRef (T);
        void resetArcs();
        bool contains (T);

    protected:
        node_s nodeMap;
        int nuNodes;
        int nuArcs;
        int version;
        binary_function<T,T,bool> comparer;
};

#endif // GRAPH_H_INCLUDED
