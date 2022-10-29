#include <iostream>
#include <vector>

int main()
{
    std::vector<int> v = {1, 2, 3, 4, 5};
    for (size_t i = 0; i < v.size(); ++i)
    {
        std::vector<int> v_new(v.begin(), v.begin() + i);
        v_new.insert(v_new.end(), v.begin() + i + 1, v.end());
        std::cout << "i = " << i << " : ";
        for (auto x : v_new)
            std::cout << x << " ";
        std::cout << std::endl;
    }

    return 0;
}