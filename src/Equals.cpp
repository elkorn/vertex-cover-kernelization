#include "headers/AbstractTest.h"

template <typename T>
class Equals: public AbstractTest<T>
{
public:
    Equals(const T &a, const T &b, const std::string &successMessage, const std::string &failureMessage):
        AbstractTest<T>(a, b, successMessage, failureMessage) {}
    const bool assert();
};

template <class T>
const bool Equals<T>::assert() {
    this->result = this->_a == this->_b;
    return this->result;
}