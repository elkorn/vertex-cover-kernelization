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
        typedef typename node_s::const_iterator node_it;

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
                const string toString() const;
                bool equals (Arc);
                const bool isCoveredBy (const node_p) const;
        };

        class Node
        {
            friend class Graph;

            public:
                typedef enum visited {
                    YES,
                    NO
                } visited;

                Node();
                Node (T);
                const T& getName() const;
                void setName (T);
                vector<Arc> getAdjacent();
                node_p copy();
                const unsigned int getDegree() const;
                void setAdjacent (vector<Arc>);
                void addArc (node_p, double);
                void removeArc (node_p);
                bool equals (node_p/*, binary_function<T,T,bool>*/);
                const string toString () const;
                const bool& getIsExternal () const;

            protected:
                T name;
                arc_s adjacent;
                visited state;

            private:
                Node(T, bool);
                bool isExternal;
        };


        Graph (/*binary_function<T,T,bool>*/);
        ~Graph();
        const int& getNuNodes() const;
        const int& getNuArcs() const;
        int getVersion();
        vector<node_p> getNodeMap();
        node_p getNode (T);
        void addNode (node_p);
        void addNode (T);
        void removeNode (node_p);
        void removeNode (T);
        const node_s& getNodes() const;
        void addArc (node_p , node_p , double);
        void addArc (T, T, double);
        void addArc (node_p , node_p);
        void addArc (T, T);
        void removeArc (node_p, node_p);
        void removeArc (T, T);
        // void removeRef (T);s
        void resetArcs();
        bool contains (T);
        const bool isVertexCover(const Graph &) const;
        node_p makeNode(T);

    protected:
        node_s nodeMap;
        int nuNodes;
        int nuArcs;
        int version;

        void clearVisited();
        void makeVisited(node_p);
};

#endif // GRAPH_H_INCLUDED
