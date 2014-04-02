#ifndef PRINTER
#define PRINTER

#include <iostream>
#include <vector>
#include <set>

using namespace std;

struct Printer
{
    template <typename T>
    static void printVector (vector<T> vec)
    {
        for (int col = 0, x_m = vec.size(); col < x_m; ++col)
        {
            cout << vec[col] << " ";
        }

        cout << endl;
    }

    template <typename T>
    static void printSet (set<T> s)
    {
        for (typename set<T>::iterator it = s.begin(), end = s.end(); it != end; ++it)
        {
            cout << *it << " ";
        }

        cout << endl;
    }

    template <typename T>
    static void printNeighbors (vector<set<T>> neighbors)
    {
        for (int i = 0, l = neighbors.size(); i < l; ++i)
        {
            cout << i << ": ";
            printSet (neighbors.at (i));
        }
    }


    template<typename T>
    static void printMatrix (vector<vector<T>> matrix)
    {
        for (int row = 0, y_m = matrix.size(); row < y_m; ++row)
        {
            printVector (matrix[row]);
            cout << endl;
        }
    }
};

#endif