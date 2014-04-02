#include <iostream>
#include <fstream>
#include <string>
#include <vector>
#include <set>
#include <utility>
#include "printer.cpp"
#include "graph.cpp"

using namespace std;

class SolverBuss
{
    public:
        SolverBuss (Graph g) : graph (g) {
            cout << "Buss solver for graph:" << endl;
            graph.print();
        }

        bool isVertexCover (vector<int> inputCoverCandidate)
        {
            vector<int> coverCandidate =
                convertCoverToContainmentList (inputCoverCandidate);
            cout << "Cover candidate:" << endl;
            Printer::printVector (coverCandidate);
            int n = graph.size();

            for (int row = 0; row < n; ++row)
            {
                for (int col = row + 1; col < n; ++col)
                {
                    if (graph.areConnected (row, col) &&
                            coverCandidate[row] == 0 &&
                            coverCandidate[col] == 0)
                    {
                        cout << "Is NOT a vertex cover." << endl;
                        return false;
                    }
                }
            }

            cout << "Is a vertex cover." << endl;
            return true;
        }

    private:
        Graph graph;

        vector<int> convertCoverToContainmentList (vector<int> cover)
        {
            vector<int> result;
            result.resize (graph.size(), 0);

            for (int i = 0; i < cover.size(); ++i)
            {
                result[cover[i]] = 1;
            }

            return result;
        }
};












