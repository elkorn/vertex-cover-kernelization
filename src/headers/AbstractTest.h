#include "AbstractTestBase.h"

template<typename T>
class AbstractTest: public AbstractTestBase
{
public:
    AbstractTest(const T &a, const T &b, const std::string &successMessage, const std::string &failureMessage);

    const virtual bool assert();
    const std::string& getMessage() const;

protected:
    const T &_a;
    const T &_b;
    bool result;

private:
    const std::string &_successMessage;
    const std::string &_failureMessage;
};