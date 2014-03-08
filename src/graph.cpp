#include <algorithm>
#include "headers/graph.h"
#include "node.cpp"
#include "arc.cpp"

using namespace std;

template <typename T>
Graph<T>::Graph (/*binary_function<T,T,bool> theComparer*/)
{
    nuNodes = 0;
    nuArcs  = 0;
    version = 0;
    // comparer = theComparer;
}

template <typename T>
int Graph<T>::getNuArcs()
{
    return nuArcs;
}

template <typename T>
int Graph<T>::getNuNodes()
{
    return nuNodes;
}

template <typename T>
int Graph<T>::getVersion()
{
    return version;
}

template <typename T>
void Graph<T>::addNode (node_p n)
{
    nodeMap.push_back (n);
    nuNodes++;
    version++;
}

template <typename T>
void Graph<T>::addNode (T n)
{
    if (!contains (n)) {
        addNode (new Node (n));
    }
}

template <typename T>
void Graph<T>::removeNode (node_p n)
{
    // removeRef (name);
    for (unsigned int i = 0; i < nodeMap.size(); i++) {
        node_p p = nodeMap.at (i);

        if (n == p) {
            vector<Arc> adjacent = p->getAdjacent();

            for (arc_it it = adjacent.begin(),
                 end = adjacent.end();
                 it != end;
                 ++it) {
                it->getHead()->removeArc(n);
            }

            delete p;
            p = 0;
            n = 0;
            nodeMap.erase (nodeMap.begin() + i);
            nuNodes--;
            version++;
        }
    }
}

template <typename T>
void Graph<T>::removeNode (T name)
{
    return removeNode (getNode (name));
}

template <typename T>
vector<typename Graph<T>::node_p> Graph<T>::getNodes()
{
    return nodeMap;
}

// template <typename T>
// bool Graph<T>::addArc (node_p tail, node_p head, double weight)
// {
//     string h = head.getName();
//     string t = tail.getName();
//     return addArc (t, h, weight);
// }

// template <typename T>
// bool Graph<T>::addArc (Node tail, Node head)
// {
//     return addArc (tail, head, 1);
// }

template <typename T>
void Graph<T>::addArc (T tail, T head, double weight)
{
    node_p n_tail = getNode (tail),
           n_head = getNode (head);

    if (n_tail != NULL  && n_tail != NULL) {
        n_tail->addArc (n_head, weight);
        n_head->addArc (n_tail, weight);
        nuArcs++;
        version++;
    }
}

template <typename T>
void Graph<T>::addArc (T tail, T head)
{
    return addArc (tail, head, 1);
}

template <typename T>
void Graph<T>::removeArc (node_p head, node_p tail)
{
    head->removeArc (tail);
    tail->removeArc (head);
    nuArcs--;
    version++;
}

template <typename T>
void Graph<T>::removeArc (T head, T tail)
{
    return removeArc (getNode (head), getNode (tail));
}

// template <typename T>
// void Graph<T>::removeRef (T name)
// {
//     for (unsigned int i = 0; i < nodeMap.size(); i++) {
//         Node n = nodeMap.at (i);
//         vector<Arc> adjacent = n.getAdjacent();

//         for (unsigned int j = 0; j < adjacent.size(); j++) {
//             Arc arc = adjacent.at (j);

//             if (arc.getHead().compare (name) == 0) {
//                 adjacent.erase (adjacent.begin() + j);
//                 nuArcs--;
//             }
//         }

//         n.setAdjacent (adjacent);
//     }
// }

template <typename T>
void Graph<T>::resetArcs()
{
    for (unsigned int i = 0; i < nodeMap.size(); i++) {
        node_p n = nodeMap.at (i);
        vector<Arc> arcs = n->getAdjacent();
        arcs.clear();
        n->setAdjacent (arcs);
    }

    nuArcs = 0;
}

template <typename T>
bool Graph<T>::contains (T label)
{
    for (unsigned int i = 0; i < nodeMap.size(); i++) {
        node_p n = nodeMap.at (i);

        if (n->getName() == label) {
            // if (comparer((T)n->getName(), (T)label)) {
            return true;
        }
    }

    return false;
}

template <typename T>
typename Graph<T>::node_p Graph<T>::getNode (T label)
{
    for (unsigned int i = 0; i < nodeMap.size(); i++) {
        if (nodeMap.at (i)->getName() == label) {
            return nodeMap.at (i);
        }
    }

    return NULL;
}

template <typename T>
vector<typename Graph<T>::node_p> Graph<T>::getNodeMap()
{
    return nodeMap;
}

template <typename T>
Graph<T>::~Graph()
{
    for (int i = 0, l = nodeMap.size(); i < l; ++i) {
        delete nodeMap.at (i);
    }

    nodeMap.clear();
}

