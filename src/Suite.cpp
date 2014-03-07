#include "Equals.cpp"
#include "headers/Suite.h"

Suite::Suite(std::vector<AbstractTestBase*> &tests):
    _name(""),
    _tests(tests) {}

Suite::Suite(std::vector<AbstractTestBase*> &tests, const std::string &name):
    _name(name),
    _tests(tests) {}

void Suite::run(std::ostream &out)
{
    bool allOk = true;
    for (int i = 0, l = _tests.size(); i < l && allOk; ++i)
    {
        allOk = _tests.at(i)->assert();
        out << (allOk ? "[ OK ]" : "[FAIL]") << " " << _tests.at(i)->getMessage() << std::endl;
    }
}