#include <string>

class AbstractTestBase {
public:
    const virtual bool assert();
    const virtual std::string& getMessage() const;
};