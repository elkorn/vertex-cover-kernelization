#include <iostream>
#include <vector>

class Suite {
    public:
        Suite(std::vector<AbstractTestBase*> &tests);
        Suite(std::vector<AbstractTestBase*> &tests, const std::string &name);

    void run(std::ostream& out);

    private:
        const std::string &_name;
        std::vector<AbstractTestBase*> &_tests;

};