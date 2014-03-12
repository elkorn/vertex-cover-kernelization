#include <algorithm>
#include <iostream>
#include "headers/graph.h"
#include "node.cpp"
#include "arc.cpp"

using namespace std;

#define SHOULD_NO_EDGES_ALLOW_VERTEX_COVER false

template <typename T>
Graph<T>::Graph (/*binary_function<T,T,bool> theComparer*/)
{
    nuNodes = 0;
    nuArcs  = 0;
    version = 0;
    // comparer = theComparer;
}

template <typename T>
const int& Graph<T>::getNuArcs() const
{
    return nuArcs;
}

template <typename T>
const int &Graph<T>::getNuNodes() const
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
    n->isExternal = true;
    nodeMap.push_back (n);
    nuNodes++;
    version++;
}

template <typename T>
void Graph<T>::addNode (T n)
{
    if (!contains (n))
    {
        nodeMap.push_back (new Node (n, false));
        nuNodes++;
        version++;
    }
}

template <typename T>
void Graph<T>::removeNode (node_p n)
{
    // removeRef (name);
    for (unsigned int i = 0; i < nodeMap.size(); i++)
    {
        node_p p = nodeMap.at (i);

        if (n == p)
        {
            vector<Arc> adjacent = p->getAdjacent();

            for (arc_it it = adjacent.begin(),
                    end = adjacent.end();
                    it != end;
                    ++it)
            {
                it->getHead()->removeArc (n);
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
const typename Graph<T>::node_s &Graph<T>::getNodes() const
{
    return nodeMap;
}

template <typename T>
void Graph<T>::addArc (T tail, T head, double weight)
{
    node_p n_tail = getNode (tail),
           n_head = getNode (head);

    if (n_tail != NULL  && n_tail != NULL)
    {
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
    for (unsigned int i = 0; i < nodeMap.size(); i++)
    {
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
    for (unsigned int i = 0; i < nodeMap.size(); i++)
    {
        node_p n = nodeMap.at (i);

        if (n->getName() == label)
        {
            // if (comparer((T)n->getName(), (T)label)) {
            return true;
        }
    }

    return false;
}

template <typename T>
typename Graph<T>::node_p Graph<T>::getNode (T label)
{
    for (unsigned int i = 0; i < nodeMap.size(); i++)
    {
        if (nodeMap.at (i)->getName() == label)
        {
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
    for (int i = 0, l = nodeMap.size(); i < l; ++i)
    {
        node_p p = nodeMap.at (i);

        if (p != NULL && !p->getIsExternal())
        {
#ifdef DEBUG
            cout << "Deleting " << p->getName() << endl;
#endif
            delete p;
            nodeMap[i] = NULL;
        }

#ifdef DEBUG

        else
        {
            cout << p->getName() << " is external, not deleting" << endl;
        }

#endif
    }

    nodeMap.clear();
}

template <typename T>
void Graph<T>::clearVisited()
{
    for (node_it it = nodeMap.begin(), end = nodeMap.end(); it != end; ++it)
    {
        (*it)->visited = Node::visited.NO;
    }
}

template <typename T>
void Graph<T>::makeVisited (node_p node)
{
    node->visited = Node::visited.YES;
}

template <typename T>
typename Graph<T>::node_p Graph<T>::makeNode (T val)
{
    return new Node (val, true);
}


template <typename T>
const bool  Graph<T>::isVertexCover (const Graph<T> &supset) const
{
    // In the context of this method, `this` is a subset of `supset`
    // Formally:
    // `this` = $S(V_s, E_s); |E_s|=k$
    // `supset` = $G(V,E); |E|=n$
    if (getNuNodes() == 0)
    {
        return false;
    }

    if (supset.getNuArcs() == 0)
    {
        return SHOULD_NO_EDGES_ALLOW_VERTEX_COVER;
    }

    // O(knm)
    for (node_it nit = getNodes().begin(),
            nend = getNodes().end();
            nit != nend;
            ++nit)
    {
        node_s nodes = supset.getNodes();

        // O(n)
        for (node_it snit = nodes.begin(),
                snend = nodes.end();
                snit != snend;
                ++snit)
        {
            arc_s adjacent = (*snit)->getAdjacent();

            // O(m), where m is adjacent.size
            for (arc_it sait = adjacent.begin(),
                    saend = adjacent.end();
                    sait != saend;
                    ++sait)
            {
                if (!sait->isCoveredBy (*nit))
                {
                    return false;
                }
            }
        }
    }

    return true;
}


