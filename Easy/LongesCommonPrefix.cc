class Solution {
 public:
  std::string longestCommonPrefix(std::vector<std::string>& strs) {
    int flag = 1;
    std::string str = "";
    for (int i = 0; flag && strs[0][i]; i++) {
      for (int j = 1; j < strs.size(); ++j) {
        if (strs[0][i] != strs[j][i] || !strs[j][i]) {
          flag = 0;
          break;
        }
      }
      if (flag) str += strs[0][i];
    }
    return str;
  }
};
