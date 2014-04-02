#include <fstream>
#include <string>
#include <vector>
#include <set>
#include <utility>
#include "printer.cpp"

using namespace std;

class Graph
{
    public:
        Graph (const Graph &other) :
            n (other.size()),
            displayedSize (n)
        {
            graph.resize(other.graph.size());
            for(int i = 0; i < other.graph.size(); ++i) {
                graph[i] = vector<int>(other.graph.at(i));
            }

            vertexRemoved.resize(n);
            collectNeighbors();
        }

        static Graph fromFile (string filePath)
        {
            int n;
            ifstream infile (filePath);
            infile >> n;
            vector<vector<int>> graph;

            for (int i = 0; i < n; i++)
            {
                vector<int> row;
                int isConnected;

                for (int j = 0; j < n; j++)
                {
                    infile >> isConnected;
                    row.push_back (isConnected);
                }

                graph.push_back (row);
            }

            return Graph (graph);
        }

        vector<int> convertCoverToContainmentList (vector<int> cover)
        {
            vector<int> result;
            result.resize (n, 0);

            for (int i = 0; i < cover.size(); ++i)
            {
                result[cover[i]] = 1;
            }

            return result;
        }

        bool areConnected (const int a, const int b)
        {
            return a != b && (graph[a][b] == 1 || graph[b][a] == 1);
        }

        const int &size() const
        {
            return displayedSize;
        }

        const int degree (int vertex)
        {
            return neighbors[vertex].size();
        }

        const int numberOfEdges()
        {
            int result = 0;

            for (int i = 0; i < neighbors.size(); ++i)
            {
                result += neighbors[i].size();
            }

            return result;
        }

        void removeVertex (int vertex)
        {
            if (vertex >= 0 && vertex < n)
            {
                if (!vertexRemoved.at (vertex))
                {
                    vertexRemoved[vertex] = true;
                    // displayed size is variable, n is constant.
                    --displayedSize;

                    for (int i = 0; i < neighbors.size(); ++i)
                    {
                        if (i == vertex)
                        {
                            neighbors[i].clear();
                        }

                        else
                        {
                            neighbors[i].erase (vertex);
                        }
                    }
                }
            }
        }

        void print()
        {
            Printer::printMatrix (graph);
        }

    private:
        const int n;
        int displayedSize;

        vector<vector<int>> graph;
        vector<set<int>> neighbors;
        vector<bool> vertexRemoved;

        Graph (vector<vector<int>> theGraph) :
            n (theGraph.size()),
            displayedSize (n)
        {
            graph.resize(theGraph.size());
            for(int i = 0; i < theGraph.size(); ++i) {
                graph[i] = vector<int>(theGraph.at(i));
            }

            vertexRemoved.resize(n);
            collectNeighbors();
        }

        void collectNeighbors()
        {
            for (int i = 0; i < graph.size(); i++)
            {
                set<int> neighbor;

                for (int j = 0; j < graph[i].size(); j++)
                    if (graph[i][j] == 1)
                    {
                        neighbor.insert (j);
                    }

                neighbors.push_back (neighbor);
            }
        }
};












