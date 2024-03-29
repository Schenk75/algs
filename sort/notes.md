| 算法     | 最好         | 最坏       | 平均              | 空间           | 稳定性 | 是否基于比较 |
| -------- | ------------ | ---------- | ----------------- | -------------- | ------ | ------------ |
| 冒泡排序 | $O(n)$       | $O(n^2)$   | $O(n^2)$          | $O(1)$         | ✓      | ✓            |
| 选择排序 | $O(n^2)$     | $O(n^2)$   | $O(n^2)$          | $O(1)$         | ×      | ✓            |
| 插入排序 | $O(n)$       | $O(n^2)$   | $O(n^2)$          | $O(1)$         | ✓      | ✓            |
| 快速排序 | $O(nlog⁡n)$   | $O(n^2)$   | $O(nlog⁡n)$        | $O(log⁡n)-O(n)$ | ×      | ✓            |
| 归并排序 | $O(nlog⁡n)$   | $O(nlog⁡n)$ | $O(nlog⁡n)$        | $O(n)$         | ✓      | ✓            |
| 希尔排序 | $O(n^{1.3})$ | $O(n^2)$   | $O(nlog⁡n)-O(n^2)$ | $O(1)$         | ×      | ✓            |
| 计数排序 | $O(n+k)$     | $O(n+k)$   | $O(n+k)$          | $O(n+k)$       | ✓      | ×            |
| 基数排序 | $O(nk)$      | $O(nk)$    | $O(nk)$           | $O(n+k)$       | ✓      | ×            |
| 桶排序   | $O(n)$       | $O(n)$     | $O(n)$            | $O(n+m)$       | ✓      | ×            |
| 堆排序   | $O(nlog⁡n)$   | $O(nlog⁡n)$ | $O(nlog⁡n)$        | $O(1)$         | ×      | ✓            |

排序算法的稳定性指：假定在待排序的记录序列中，存在多个具有相同的关键字的记录，若经过排序，这些记录的相对次序保持不变，即在原序列中，r[i]=r[j]，且r[i]在r[j]之前，而在排序后的序列中，r[i]仍在r[j]之前，则称这种排序算法是稳定的；否则称为不稳定的。