#ifndef GRAPH_NODE_H_
#define GRAPH_NODE_H_

/*
* This file has declarations for classes used to represent the Graph
*/

#include <vector>
#include <stack>
#include <string>
#include <sstream>
#include <unordered_set>

using namespace std;

template <typename T>
class Graph
{
    public:

        //An object of this class holds a vertex of the Graph
        class Node
        {
            private:
                T value;
                unsigned int degree;

            public:
                Node (T id)
                {
                    value = id;
                    degree = 0;
                }

                const T &getVal() const
                {
                    return value;
                }

                const unsigned int getDegree() const
                {
                    return degree;
                }

                const void incrDegree()
                {
                    degree++;
                }
        };

        //An object of this class represents an edge in the Graph.
        class Edge
        {
            private:
                Node *orgNode;//the originating vertex
                Node *dstNode;//the destination vertex

            public:
                Edge (Node *firstNode, Node *secNode)
                {
                    orgNode = firstNode;
                    dstNode = secNode;
                }

                const Node *getDstNode() const
                {
                    return dstNode;
                }

                const Node *getOrgNode() const
                {
                    return orgNode;
                }

                const bool isCoveredBy (const Node *node) const
                {
                    return node == orgNode || node == dstNode;
                }
        };

        ~Graph()
        {
            //free mem allocated to vertices
            for (node_it it = nodes.begin(),
                 end = nodes.end();
                 it != end;
                 ++it) {
                delete (*it);
            }

            nodes.clear();

            for (edge_it it = edges.begin(),
                 end = edges.end();
                 it != end;
                 ++it) {
                delete (*it);
            }

            edges.clear();
        }

        const unsigned int size() const
        {
            return nodes.size();
        }

        const void display(ostream &output)
        {
            output << "NODES" << endl;

            for (node_it it = nodes.begin(),
                 end = nodes.end();
                 it != end;
                 ++it) {
                output << (*it)->getVal()
                     << " ";
            }

            output << endl << "EDGES" << endl;
            for (edge_it it = edges.begin(),
                 end = edges.end();
                 it != end;
                 ++it) {
                output << (*it)->getOrgNode()->getVal()
                     << " -> "
                     << (*it)->getDstNode()->getVal()
                     << endl ;
            }
        }

        void insert (T val)
        {
            nodes.push_back (new Node (val));
        }

        void connect (const int origin, const int destination)
        {
            edges.insert (new Edge (nodes.at (origin), nodes.at (destination)));
        }

        Node *getNode (int index) const
        {
            return nodes.at (index);
        }

    private:
        typedef typename vector<Node *>::iterator node_it;
        typedef typename unordered_set<Edge *>::iterator edge_it;

        vector<Node *> nodes;
        unordered_set<Edge *> edges;
        // const bool isVertexCover(Graph supset) {
        //     if(size() == 0) return false;

        // }

        // Graph vertexCover (int k)
        // {
        //     Graph result;
        //     bool coverFound = false;

        //     do {
        //         bool currentNodeProcessed = false;
        //         for (int i = 0, l = this->size(); i < l && !currentNodeProcessed; ++i) {
        //             Node *cur = nodes.at (i);

        //             for (int j = 0; j < l && !currentNodeProcessed; ++j) {
        //                 if (j == i) {
        //                     continue;
        //                 }

        //                 Node *toCheck = nodes.at(j);
        //                 vector<Edge> &edges = toCheck->getEdges();
        //                 for(int k = 0, m = edges.size(); k < m  && !currentNodeProcessed; ++k) {
        //                     if(edges.at(i).isCoveredBy(cur)) {
        //                         result.addNewNode(cur);
        //                         currentNodeProcessed = true;
        //                     }
        //                 }
        //             }

        //             if(currentNodeProcessed) {

        //             }
        //         }
        //     } while (!coverFound && result.size() < k);

        //     if(!coverFound) {
        //         throw "Cover not found.";
        //     }

        //     return result;
        // }


        // Node *findNodeByName (string name)
        // {
        //     for (node_it it = nodes.begin(),
        //          end = nodes.end();
        //          it != end;
        //          ++it) {
        //         if ( (*it)->getVal() == name) {
        //             return (*it);
        //         }
        //     }

        //     return NULL;
        // }
};
#endif