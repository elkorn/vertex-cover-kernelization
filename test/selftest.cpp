#include <iostream>

#include "../src/Suite.cpp"

using namespace std;

int main(int argc, char** argv) {
    std::vector<AbstractTestBase*> tests;
    Equals<int> *t2 = new Equals<int>(1,2, "1 equals 2", "1 must equal 2");
    tests.push_back((AbstractTestBase*)t2);

    Suite suite(tests);

    return 0;
}