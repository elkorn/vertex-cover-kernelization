#ifndef GRAPH_H_INCLUDED
#include "headers/graph.h"
#endif

using namespace std;

template <typename T>
Graph<T>::Node::Node()
{
}

template <typename T>
Graph<T>::Node::Node (T n)
{
    name = n;
}

template <typename T>
T Graph<T>::Node::getName()
{
    return name;
}

template <typename T>
void Graph<T>::Node::setName (T n)
{
    name = n;
}

template <typename T>
vector<typename Graph<T>::Arc> Graph<T>::Node::getAdjacent()
{
    return adjacent;
}

template <typename T>
unsigned int  Graph<T>::Node::getDegree()
{
    return adjacent.size();
}

template <typename T>
void Graph<T>::Node::setAdjacent (vector<Arc> arcs)
{
    adjacent = arcs;
}

template <typename T>
typename Graph<T>::node_p Graph<T>::Node::copy()
{
    node_p n = new Node();
    n->setName (getName());
    n->setAdjacent (getAdjacent());
    return n;
}

template <typename T>
void Graph<T>::Node::addArc (node_p head, double weight)
{
    Arc a (head, this, weight);
    adjacent.push_back (a);
}

template <typename T>
void Graph<T>::Node::removeArc (node_p head)
{
    arc_it begin = adjacent.begin();
    for (int i = 0, l = getDegree(); i < l; ++i) {
        if (adjacent.at (i).getHead() == head) {
            adjacent.erase (begin + i);
        }
    }
}

template <typename T>
bool Graph<T>::Node::equals (node_p n/*, binary_function<T,T,bool> comparer*/)
{
    if (getName() == n->getName()) {
        if (getAdjacent().size() == n->getAdjacent().size()) {
            for (unsigned int i = 0; i < getAdjacent().size(); i++) {
                Arc n1 = getAdjacent().at (i);
                Arc c1 = n->getAdjacent().at (i);

                if (! (n1.equals (c1))) {
                    return false;
                }
            }

            return true;

        } else {
            return false;
        }
    }

    return false;
}

template <typename T>
string Graph<T>::Node::toString ()
{
    stringstream ss ("");
    ss << getName();

    for (int i = 0, l = adjacent.size(); i < l; ++i) {
        ss << std::endl << "\t" << ARC << adjacent.at (i).toString();
    }

    ss << std::endl;
    return ss.str();
}

