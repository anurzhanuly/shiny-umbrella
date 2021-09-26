class Solution:
    def kWeakestRows(self, mat: List[List[int]], k: int) -> List[int]:
        for index, arr in enumerate(mat):
            arr.append(index)
            
        mat.sort()
        
        return [arr[-1] for arr in mat[:k]]