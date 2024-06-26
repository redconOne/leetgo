# [3101. Count Alternating Subarrays][link] (Medium)

[link]: https://leetcode.com/problems/count-alternating-subarrays/

You are given a binary array `nums`.

We call a subarray **alternating** if **no** two **adjacent** elements in the
subarray have the **same** value.

Return the number of alternating subarrays in `nums`.

**Example 1:**

```text
Input:nums = \[0,1,1,1\]
Output:5
Explanation:
The following subarrays are alternating: `[0]`, `[1]`, `[1]`, `[1]`, and `[0,1]`.
```

**Example 2:**

```text
Input:nums = \[1,0,1,0\]
Output:10
Explanation:
Every subarray of the array is alternating. There are 10 possible subarrays
that we can choose.
```

**Constraints:**

- `1 <= nums.length <= 10⁵`
- `nums[i]` is either `0` or `1`.
