#include "headers/AbstractTest.h"

template <typename T>
AbstractTest<T>::AbstractTest(const T &a, const T &b, const std::string &successMessage, const std::string &failureMessage):
    _a(a),
    _b(b),
    _successMessage(successMessage),
    _failureMessage(failureMessage)
{}

template <typename T>
const std::string &AbstractTest<T>::getMessage() const
{
    return result ?
           _successMessage :
           _failureMessage;
}
