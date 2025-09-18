// SPDX-License-Identifier: MIT
pragma solidity ^0.8.10;


// 1. 创建一个名为Voting的合约，包含以下功能：
// 一个mapping来存储候选人的得票数
// 一个vote函数，允许用户投票给某个候选人
// 一个getVotes函数，返回某个候选人的得票数
// 一个resetVotes函数，重置所有候选人的得票数
// 一个checkAllVotes函数，返回所有候选人的得票数
 contract Voting {
    mapping(uint256 => uint256) public votes;
    uint256[] public candidates; // 记录所有投票过的候选人
    
    function vote(uint256 candidate) public {
        if (votes[candidate] == 0) {
            candidates.push(candidate); // 第一次投票时记录候选人
        }
        votes[candidate]++;
    }
    
    function getVotes(uint256 candidate) public view returns (uint256) {
        return votes[candidate];
    }
    function checkAllVotes() public view returns (uint256[] memory) {
        return candidates;
    }
    function resetVotes() public {
        // 遍历所有候选人并重置其得票数
        for (uint256 i = 0; i < candidates.length; i++) {
            votes[candidates[i]] = 0;
        }
        // 清空候选人数组
        delete candidates;
    }
 }

// 2. 反转字符串 (Reverse String)
// 题目描述：反转一个字符串。输入 "abcde"，输出 "edcba"
contract ReverseString {
    function reverseString(string memory s) public pure returns (string memory) {
        bytes memory str = bytes(s);
        bytes memory result = new bytes(str.length);
        uint256 j = 0;
        for (uint256 i = str.length; i > 0; i--) {
            bytes1 temp = str[i];
            result[j] = temp;
            j++;
        }
        return string(result);
    }
}

// 3. 罗马数字转整数
contract RomanToInteger {
    // 从右往左扫描，比前一个小就减，否则就加
    function romanToInt(string memory s) public pure returns (uint256) {
        bytes memory b = bytes(s);
        require(b.length > 0, "Empty string");
        
        uint256 result = 0;
        uint256 prev = 0;
        
        // 从右往左
        for (uint256 i = b.length; i > 0; i--) {
            uint256 curr = getValue(b[i-1]);
            if (curr < prev) {
                result -= curr;
            } else {
                result += curr;
            }
            prev = curr;
        }
        return result;
    }
    
    function getValue(bytes1 c) private pure returns (uint256) {
        if (c == 'I') return 1;
        if (c == 'V') return 5;
        if (c == 'X') return 10;
        if (c == 'L') return 50;
        if (c == 'C') return 100;
        if (c == 'D') return 500;
        if (c == 'M') return 1000;
        revert("Invalid Roman numeral");
    }
}

// 4. 整数转罗马数字
contract IntegerToRoman {
    function intToRoman(uint256 num) public pure returns (string memory) {
        require(num > 0 && num < 4000, "Range: 1-3999");
        
        // 预分配空间，最长罗马数字不超过15字符
        bytes memory result = new bytes(15);
        uint256 pos = 0;
        
        // 处理每个值
        while (num >= 1000) { result[pos++] = 'M'; num -= 1000; }
        if (num >= 900) { result[pos++] = 'C'; result[pos++] = 'M'; num -= 900; }
        if (num >= 500) { result[pos++] = 'D'; num -= 500; }
        if (num >= 400) { result[pos++] = 'C'; result[pos++] = 'D'; num -= 400; }
        while (num >= 100) { result[pos++] = 'C'; num -= 100; }
        if (num >= 90) { result[pos++] = 'X'; result[pos++] = 'C'; num -= 90; }
        if (num >= 50) { result[pos++] = 'L'; num -= 50; }
        if (num >= 40) { result[pos++] = 'X'; result[pos++] = 'L'; num -= 40; }
        while (num >= 10) { result[pos++] = 'X'; num -= 10; }
        if (num >= 9) { result[pos++] = 'I'; result[pos++] = 'X'; num -= 9; }
        if (num >= 5) { result[pos++] = 'V'; num -= 5; }
        if (num >= 4) { result[pos++] = 'I'; result[pos++] = 'V'; num -= 4; }
        while (num >= 1) { result[pos++] = 'I'; num -= 1; }
        
        // 截断到实际长度
        bytes memory trimmed = new bytes(pos);
        for (uint256 i = 0; i < pos; i++) {
            trimmed[i] = result[i];
        }
        return string(trimmed);
    }
}

// 5. 合并两个有序数组
contract MergeSortedArray {
    function merge(uint256[] memory nums1, uint256[] memory nums2) 
        public pure returns (uint256[] memory) {
        uint256[] memory result = new uint256[](nums1.length + nums2.length);
        uint256 i = 0;
        uint256 j = 0;
        uint256 k = 0;
        
        while (i < nums1.length && j < nums2.length) {
            if (nums1[i] <= nums2[j]) {
                result[k++] = nums1[i++];
            } else {
                result[k++] = nums2[j++];
            }
        }
        
        while (i < nums1.length) result[k++] = nums1[i++];
        while (j < nums2.length) result[k++] = nums2[j++];
        
        return result;
    }
}

// 6. 二分查找
contract BinarySearch {
    function search(uint256[] memory nums, uint256 target) 
        public pure returns (int256) {
        uint256 left = 0;
        uint256 right = nums.length;
        
        while (left < right) {
            uint256 mid = left + (right - left) / 2;
            if (nums[mid] == target) {
                return int256(mid);
            } else if (nums[mid] < target) {
                left = mid + 1;
            } else {
                right = mid;
            }
        }
        return -1; // 找不到
    }
}