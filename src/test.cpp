#include <iostream>

using namespace std;

#define PASS_THROUGH
class Test
{
    public:
        static void title (const string title)
        {
            cout << " *** TESTING: " << title << " *** " << endl;
        }

        static void test (bool result, string msg)
        {
            cout << "  [" <<
                 (result ? "\033[22;32m OK " : "\033[22;31mFAIL") <<
                 "\033[22;0m] " <<
                 msg <<
                 endl;
            #ifndef PASS_THROUGH
            if (!result) {
                abort();
            }
            #endif
        }

};