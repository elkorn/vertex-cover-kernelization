#ifndef GRAPH_H_INCLUDED
#include "headers/graph.h"
#endif

using namespace std;

template <typename T>
Graph<T>::Node::Node()
{
    state = NO;
}

template <typename T>
Graph<T>::Node::Node (T n) :
    name (n),
    isExternal (true)
{
    state = NO;
}

template <typename T>
Graph<T>::Node::Node (T n, bool external) :
    name (n),
    isExternal (external) {
        state = NO;
}

template <typename T>
const T &Graph<T>::Node::getName() const
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
const unsigned int  Graph<T>::Node::getDegree() const
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
const string Graph<T>::Node::toString () const
{
    stringstream ss ("");
    ss << getName();

    for (int i = 0, l = getDegree(); i < l; ++i) {
        ss << "\t" << ARC << adjacent.at (i).toString();

        if (i < l - 1) {
            ss << std::endl;
        }
    }

    // ss << std::endl;
    return ss.str();
}

template <typename T>
const bool& Graph<T>::Node::getIsExternal () const {
    return isExternal;
}

