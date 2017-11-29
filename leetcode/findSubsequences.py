#!/bin/python

def findSubsequences(nums):
    ret = set()
    ret = findSubsequence(0, -101, (), nums, ret)
    print list(ret)

def findSubsequence(p, prev, n, nums, ret):
    if len(n) >= 2:
        ret.add(n)

    for i in range(p, len(nums)):
        if len(n) == 0 or nums[i] >= prev:
            ret = findSubsequence(i+1, nums[i], n + (nums[i],), nums, ret)
    return ret

findSubsequences([1,2,1,1])
