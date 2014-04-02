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

        bool areConnected(const int a, const int b) {
            return a != b && (graph[a][b] == 1 || graph[b][a] == 1);
        }

        const int& size() {
            return n;
        }

        void print() {
            Printer::printMatrix(graph);
        }

    private:
        int n;

        vector<vector<int>> graph;
        vector<set<int>> neighbors;

        Graph (vector<vector<int>> theGraph) :
            n (theGraph.size()),
            graph (theGraph)
        {
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

        vector<vector<int>> removeVertex (vector<vector<int>> graph, int vertex)
        {
            vector<vector<int>> result = graph;

            for (int i = 0; i < result.size(); ++i)
            {
                // result[vertex].erase(result[i].begin() + i);
                // result[i].erase(result[vertex].begin() + i);
                result[i][vertex] = 0;
                result[vertex][i] = 0;
            }

            return result;
        }
};












