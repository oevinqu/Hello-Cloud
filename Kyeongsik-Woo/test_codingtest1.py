arr = [11,3,28,43,9,2]
N = 6

for i in range(N-1):
    minIdx = i
    for j in range(i+1, N):
        if arr[j] < arr[minIdx]:
            minIdx = j

    arr[minIdx],arr[i] = arr[i],arr[minIdx]

print(arr)


def rec(num):
    if num<=0:
        return
    print(num, end=" ")
    rec(num-1)

rec(3)

def rec2(num):
    if num <= 0:
        return
    rec(num-1)
    rec(num-1)
    print(num,end=" ")

rec2(3)

def rec3(num):
    if num <=0:
        return
    rec(num-1)
    rec(num-2)
    rec(num-3)
    print(num, end=" ")

rec3(4)

#퀵정렬1
import random

arr = [
    random.randint( 0 , 200) for i in range(100) # 이 부분!
]
       

def qsort(data):
    if len(data)<=1:
        return data
    pivot = data[0]
    left = []
    right = []

    for i in range(1,len(data)):
        if data[i] < pivot:
            left.append(data[i])
        else:
            right.append(data[i])
    left = qsort(left)
    right = qsort(right)
    data = left + [pivot] + right
    return data
arr = qsort(arr)
print(arr)

#퀵정렬2
def qsort2(data):
    if len(data)<=1:
        return data
    pivot = data[0]
    left = [x for x in data[1:] if x < pivot]
    right = [y for y in data[1:] if y > pivot]
    return qsort2(left) + [pivot] + qsort2(right)

lst = [345,32,6,2,62457,1341,3]
re = qsort2(lst)
print(re)

#재귀함수 리밋을 10000까지 늘리기
import sys
sys.setrecursionlimit(10000)

# 합병 정렬
def mergesort(data):
    if len(data) <=1:
        return data
    
    mid = len(data)//2
    left = mergesort(data[:mid])
    right = mergesort(data[mid:])
    leftP = 0
    rightP = 0
    temp = []
    while leftP < len(left) and rightP < len(right):
        if left[leftP] <= right[rightP]:
            temp.append(left[leftP])
            leftP += 1
        else:
            temp.append(right[rightP])
            rightP += 1
    
    while leftP < len(left):
        temp.append(left[leftP])
        leftP += 1
    
    while rightP < len(right):
        temp.append(right[rightP])
        rightP += 1
        
    return temp

arr1 = mergesort(arr)
print(arr1)

def testsort(data):
    if len(data) <= 1:
        return data
    
    mid = len(data) // 2

    left = testsort(data[mid:])
    right = testsort(data[:mid])
    leftP = 0
    rightP = 0
    temp = []

    while leftP < len(left) and rightP <len(right):
        if left[leftP] < right[rightP]:
            temp.append(left[leftP])
            leftP += 1
        else:
            temp.append(right[rightP])
            rightP += 1

    while leftP < len(left):
        temp.append(left[leftP])
        leftP += 1

    while rightP < len(right):
        temp.append(right[rightP])
        rightP += 1
    
    return temp

arr = testsort(arr)
print(arr)
             
#코테 정렬 1
N = int(input())
arr = []
for i in range(N):
    v = int(input())
    arr.append(v)
arr.sort()

for v in arr:
    print(v)

#코테 정렬 2 카운팅 정렬
import sys
input = sys.stdin.readline

N = int(input())

cnt = [0 for i in range(10001)]

for i in range(N):
    v = int(input())
    cnt[v] += 1

for i in range(10001):
    for j in range(cnt[i]):
        print(i)

#코테 중복 x
import sys
input = sys.stdin.readline

N = int(input())

arr = list(map(int,input().split()))
s = set(arr)
arr = list(s)

arr.sort(reverse= True)
print(*arr)

 #코테 
import sys
input = sys.stdin.readline

N = int(input())

arr = []

#달리기 ,팔굽혀,윗몸 순위 알아보기 1<=N<=10

for i in range(N):
    name, a, b, c = input().split()
    a = int(a)
    b = int(b)
    c = int(c)
    arr.append([name,a,b,c])

arr.sort(key= lambda v : (v[1], v[1]+v[2]+v[3]), reverse= True)
for v in arr:
    print(v[0])
print()

arr.sort(key= lambda v : (v[2], v[1]+v[2]+v[3]), reverse= True)
for v in arr:
    print(v[0])
print()

arr.sort(key= lambda v : (v[3], v[1]+v[2]+v[3]), reverse= True)
for v in arr:
    print(v[0])
print()

#코테

N,K = map(int,input().split())

arr = []
for i in range(K):
    a,b = map(int, input().split())
    arr.append(a*b)

arr.sort(reverse=True)

for i in range(N):
    N -= arr[i]
    if N <= 0:
        print(i+1)
        break


import sys

input = sys.stdin.readline

N = int(input())

arr = []

for i in range(N):
    arr.append(int(input()))

arr.sort()
print(arr)



arr = [10,2,5,16,457,4,2,4]

arr= list(dict.fromkeys(arr))

import sys
input = sys.stdin.readline


N= int(input())
arr = list(map(int, input().split()))

arr.sort()
total_time = 0
waiting_time = 0

for j in arr:
    waiting_time += j
    total_time += waiting_time

print(waiting_time)

import sys
input = sys.stdin.readline

N = int(input())

arr = []

for i in range(N):
    arr.append(input().strip())

arr = list(set(arr))

arr.sort(key=lambda x: (len(x),x))

for word in arr:
    print(word)