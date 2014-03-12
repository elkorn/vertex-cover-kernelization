#include <iostream>
#include <fstream>
#include <string>
#include <vector>
using namespace std;

// http://www.dharwadker.org/vertex_cover/


template <typename T>
void printVector (vector<T> vec)
{
    for (int col = 0, x_m = vec.size(); col < x_m; ++col)
    {
        cout << vec[col] << " ";
    }
}

template<typename T>
void printMatrix (vector<vector<T>> matrix)
{
    for (int row = 0, y_m = matrix.size(); row < y_m; ++row)
    {
        printVector (matrix[row]);
        cout << endl;
    }
}

void printVertices (vector<int> vertices)
{
    for (int j = 0; j < vertices.size(); j++)
    {
        if (vertices[j] == 1)
        {
            cout << j + 1 << " ";
        }
    }
}

void printCover (vector<int> cover, int size)
{
    // Output the result.
    cout << "Vertex Cover (" << size << "): ";
    printVertices (cover);
    cout << endl;
    cout << "Vertex Cover Size: " << size << endl;
}

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
            printMatrix (graph);
        }

        void findVertexCovers (int k)
        {
            bool found = false;
            cout << "Finding Vertex Covers..." << endl;
            int min = n + 1, counter = 0;
            vector<vector<int> > covers;
            // At the beginning, all vertices belong to a cover.
            vector<int> allcover;

            for (int i = 0; i < graph.size(); i++)
            {
                allcover.push_back (1);
            }

            /*
            Part I. For i = 1, 2, ..., n in turn
              - Initialize the vertex cover Ci = V−{i}.
              - Perform procedure 3.1 on Ci.
              - For r = 1, 2, ..., n−k perform procedure 3.2 repeated r times.
              - The result is a minimal vertex cover Ci.
             */
            for (int i = 0; i < allcover.size(); i++)
            {
                if (found)
                {
                    break;
                }

                counter++;
                cout << counter << ". ";
                // Initialize the vertex cover Ci = V−{i}.
                vector<int> cover = allcover;
                cover[i] = 0;
                // Perform procedure 3.1 on Ci.
                cover = procedure_1 (neighbors, cover);
                int s = cover_size (cover);

                if (s < min)
                {
                    min = s;
                }

                if (s <= k)
                {
                    printCover (cover, s);
                    covers.push_back (cover);
                    found = true;
                    break;
                }

                // For j = 1, 2, ..., n−k perform procedure 3.2 repeated j times.
                for (int j = 0; j < n - k; j++)
                {
                    cover = procedure_2 (neighbors, cover, j);
                }

                s = cover_size (cover);

                if (s < min)
                {
                    min = s;
                }

                printCover (cover, s);

                if (s <= k)
                {
                    found = true;
                    break;
                }
            }

            /*
              Part II. For each pair of minimal vertex covers Ci, Cj found in Part I
              - Initialize the vertex cover Ci, j = Ci∪Cj .
              - Perform procedure 3.1 on Ci, j.
              - For r = 1, 2, ..., n−k perform procedure 3.2 repeated r times.
              - The result is a minimal vertex cover Ci, j.
             */
            for (int p = 0; p < covers.size(); p++)
            {
                if (found)
                {
                    break;
                }

                for (int q = p + 1; q < covers.size(); q++)
                {
                    if (found)
                    {
                        break;
                    }

                    counter++;
                    cout << counter << ". ";
                    // Initialize the vertex cover Ci,j = V−{i}-{j}.
                    vector<int> cover = allcover;

                    for (int r = 0; r < cover.size(); r++)
                    {
                        if (covers[p][r] == 0 && covers[q][r] == 0)
                        {
                            cover[r] = 0;
                        }
                    }

                    // Perform procedure 3.1 on Ci.
                    cover = procedure_1 (neighbors, cover);
                    // Calculate the size of the cover.
                    int s = cover_size (cover);

                    if (s < min)
                    {
                        min = s;
                    }

                    if (s <= k)
                    {
                        printCover (cover, s);
                        found = true;
                        break;
                    }

                    // For j = 1, 2, ..., n−k perform procedure 3.2 repeated j times.
                    for (int j = 0; j < k; j++)
                    {
                        cover = procedure_2 (neighbors, cover, j);
                    }

                    // Calculate the size of the cover.
                    s = cover_size (cover);

                    if (s < min)
                    {
                        min = s;
                    }

                    // Output the result.
                    printCover (cover, s);

                    if (s <= k)
                    {
                        found = true;
                        break;
                    }
                }
            }

            if (found)
            {
                cout << "Found Vertex Cover of size at most " << k << "." << endl;
            }

            else cout << "Could not find Vertex Cover of size at most " << k << "." << endl
                          << "Minimum Vertex Cover size found is " << min << "." << endl;
        }

    private:
        int n;
        vector<vector<int>> graph;
        vector<vector<int>> neighbors;

        void collectNeighbors()
        {
            for (int i = 0; i < graph.size(); i++)
            {
                vector<int> neighbor;

                for (int j = 0; j < graph[i].size(); j++)
                    if (graph[i][j] == 1)
                    {
                        neighbor.push_back (j);
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

        /*
            Get the maximum number of vertices that can be removed from given
            cover.

            Complexity: O(n^4)
        */
        int max_removable (vector<vector<int> > neighbors, vector<int> cover)
        {
            int rho = -1, max = -1;

            // Given a vertex cover C of G and a vertex v in C, we say that
            // v is removable if the set C−{v} is still a vertex cover of G.
            // O(n^4)
            for (int v = 0; v < cover.size(); v++)
            {
                // if v belongs to a cover and is removable from it
                // O(n^3)
                if (cover[v] == 1 && removable (neighbors[v], cover) == true)
                {
                    vector<int> temp_cover = cover;
                    // Remove v from the cover
                    temp_cover[v] = 0;
                    int sum = 0;

                    // Count all neighboring vertices that also belong to that
                    // cover and can be removed from it.
                    // O(n^2)
                    for (int j = 0; j < temp_cover.size(); j++)
                    {
                        if (temp_cover[j] == 1 && removable (neighbors[j], temp_cover) == true)
                        {
                            sum++;
                        }
                    }

                    if (sum > max)
                    {
                        max = sum;
                        rho = v;
                    }
                }
            }

            return rho;
        }


        /*
            Given a simple graph G with n vertices and a vertex cover C of G, if C has no
            removable vertices, output C. Else, for each removable vertex v of C, find the
            number ρ(C−{v}) of removable vertices of the vertex cover C−{v}.
            Let vmax denote a removable vertex such that ρ(C−{vmax}) is a maximum and
            obtain the vertex cover C−{vmax}. Repeat until the vertex cover has no
            removable vertices.

            This procedure effectively minimalizes the vertex cover. The result
            is a minimal vertex cover (as opposed to a minimum vertex cover).

            Complexity: O(n^5)

         */
        vector<int> procedure_1 (vector<vector<int> > neighbors, vector<int> cover)
        {
            vector<int> temp_cover = cover;
            int rho = 0;
            bool vertexCanBeRemovedFromCover = true;

            // O(n^5)
            while (vertexCanBeRemovedFromCover)
            {
                // for each removable vertex v of C, find the
                // number ρ(C−{v}) of removable vertices of the vertex cover C−{v}.
                // O(n^4)
                rho = max_removable (neighbors, temp_cover);
                vertexCanBeRemovedFromCover = rho != -1;

                if (vertexCanBeRemovedFromCover)
                {
                    temp_cover[rho] = 0;
                }
            }

            return temp_cover;
        }

        /*
         Given a simple graph G with n vertices and a minimal vertex cover C of G, if
         there is no vertex v in C such that v has exactly one neighbor w outside C,
         output C. Else, find a vertex v in C such that v has exactly one neighbor w
         outside C. Define Cv,w by removing v from C and adding w to C.
         Perform procedure_1 on Cv,w and output the resulting vertex cover.
         */
        vector<int> procedure_2 (vector<vector<int> > neighbors, vector<int> cover,
                                 int k)
        {
            int reductionAttempts = 0;
            vector<int> temp_cover = cover;
            int v, j;

            // O(n)
            for (v = 0; v < temp_cover.size(); v++)
            {
                if (temp_cover[v] == 1)
                {
                    bool hasOnlyOneNeighborOutsideTheCover = false;
                    int w_index;

                    // Count all the neighbors of vertex v that do not belong to
                    // cover Cv.
                    // O(n)
                    for (j = 0; j < neighbors[v].size(); j++)
                    {
                        if (temp_cover[neighbors[v][j]] == 0)
                        {
                            if (hasOnlyOneNeighborOutsideTheCover)
                            {
                                // has multiple neighbors outside the cover.
                                hasOnlyOneNeighborOutsideTheCover = false;
                                break;
                            }

                            else
                            {
                                w_index = j;
                                hasOnlyOneNeighborOutsideTheCover = true;
                            }
                        }
                    }

                    // If there is exactly one neighbor of v that does not
                    // belong to cover C.
                    // The second part of this conditional is probably redundant
                    if (hasOnlyOneNeighborOutsideTheCover &&
                            cover[neighbors[v][w_index]] == 0)
                    {
                        // Exchange vertices v and w
                        // O(1)
                        temp_cover[neighbors[v][w_index]] = 1;
                        temp_cover[v] = 0;
                        // Minimize the cover.
                        temp_cover = procedure_1 (neighbors, temp_cover);
                        reductionAttempts++;
                    }

                    if (reductionAttempts > k)
                    {
                        break;
                    }
                }
            }

            return temp_cover;
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

int main()
{
    Graph g;
    g.load ("../data/frucht12.txt");
    g.findVertexCovers (7);
    return 0;
}















