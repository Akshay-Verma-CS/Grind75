from collections import defaultdict
def two_sum(arr,target):
    n = len(arr)
    d = defaultdict(lambda x: 0)
    for i in range(n):
        d[abs(arr[i] - target)] = i
    for 