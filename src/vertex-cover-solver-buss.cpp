#include <iostream>
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
        void load (string filePath)
        {
            ifstream infile (filePath);
            infile >> n;

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

            collectNeighbors();
            Printer::printMatrix(graph);
        }

        vector<int> convertCoverToContainmentList(vector<int> cover) {
            vector<int> result;
            result.resize(n, 0);

            for(int i = 0; i < cover.size(); ++i) {
                result[cover[i]] = 1;
            }

            return result;
        }

        bool isVertexCover (vector<int> inputCoverCandidate)
        {
            vector<int> coverCandidate = convertCoverToContainmentList(inputCoverCandidate);
            cout << "Cover candidate:" << endl;
            Printer::printVector(coverCandidate);

            for(int row = 0; row < n; ++row) {
                for(int col = row + 1; col < n; ++col) {
                    if(graph[row][col] == 1 &&
                       coverCandidate[row] == 0 &&
                       coverCandidate[col] == 0) {
                        return false;
                    }
                }
            }

            return true;
        }

    private:
        int n;
        vector<vector<int>> graph;
        vector<set<int>> neighbors;

        void collectNeighbors()
        {
            for (int i = 0; i < graph.size(); i++)
            {
                set<int> neighbor;

                for (int j = 0; j < graph[i].size(); j++)
                    if (graph[i][j] == 1)
                    {
                        cout << "Inserting neighbor " << j << " for " << i << endl;
                        neighbor.insert (j);
                    }

                neighbors.push_back (neighbor);
            }
        }

        /*
            Determine whether a vertex is removable from a cover.
            A vertex can be removed from a cover if all of its neighbors belong
            to the same cover.

            Complexity: O(n)
         */
        bool removable (vector<int> neighborsOfVertex, vector<int> cover)
        {
            // The vertex can be removed if all of its neighbors belong to a cover.
            // O(n)
            for (int i = 0; i < neighborsOfVertex.size(); i++)
            {
                if (cover[neighborsOfVertex[i]] == 0)
                {
                    return false;
                }
            }

            return true;
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

        int cover_size (vector<int> cover)
        {
            int count = 0;

            for (int i = 0; i < cover.size(); i++)
                if (cover[i] == 1)
                {
                    count++;
                }

            return count;
        }
};












